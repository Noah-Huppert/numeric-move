package main

import (
	"fmt"
	"regexp"
	"strconv"
	"errors"
	"strings"
)

// numPrefixExp matches a numerically prefixed file.
// The first match group is the numeric prefix.
// The second match group is the un-prefixed file name.
var numPrefixExp *regexp.Regexp = regexp.MustCompile("^([0-9]+)(.*)$")

// FileNode is a node in a doubly linked list.
// It holds information about a numerically prefixed file.
type FileNode struct {
	// Name is the file's name without the numeric prefix
	Name string

	// Prefix is the file's numeric prefix
	Prefix uint64

	// PrefixLength is the number of digits used to represent the prefix.
	// If the Prefix can be represented in less the extra digits are used
	// by placing 0's before the prefix.
	PrefixLength uint64

	// PrevDelta is the difference between the previous node's prefix and Prefix.
	// This field will start un-set. Once all files have been added to a FileList
	// the FileList will traverse itself and set this field. This ensures that
	// PrevDelta is accurate.
	PrevDelta uint64

	// Next FileNode in list
	Next *FileNode

	// Prev FileNode in list
	Prev *FileNode
}

// ErrNotNumPrefixFile indicates a file is not numerically prefixed
var ErrNotNumPrefixFile error = errors.New("not a numerically prefixed file")

// NewFileNode creates a new FileNode from a path.
// The given path should just be a file name without any leading directory.
// Returns ErrNoNumPrefixFile if the path points to a non-numericaly prefixed file.
func NewFileNode(p string) (*FileNode, error) {
	// Check that the path is for a numerically prefixed file
	matches := numPrefixExp.FindStringSubmatch(p)

	if len(matches) == 0 {
		return nil, ErrNotNumPrefixFile
	}

	// Parse prefix
	prefix, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse numeric prefix to unsigned integer: %s", err.Error())
	}

	return &FileNode{
		Name: matches[2],
		Prefix: prefix,
		PrefixLength: uint64(len(matches[1])),
		PrevDelta: 0,
	}, nil
}

// String representation
func (f FileNode) String() string {
	return fmt.Sprintf("%d %s (dt: %d)", f.Prefix, f.Name, f.PrevDelta)
}

// FullName returns the file's name with its numeric prefix
func (f FileNode) FullName() string {
	prefixStr := strconv.FormatUint(f.Prefix, 10)
	zeros := strings.Repeat("0", int(f.PrefixLength) - len(prefixStr))

	return zeros + prefixStr + f.Name
}

// SetPrefix sets the Prefix field and updates PrefixLength if it is too small to hold the new prefix
func (f *FileNode) SetPrefix(p uint64) {
	requiredLength := uint64(len(strconv.FormatUint(p, 10)))

	if f.PrefixLength < requiredLength {
		f.PrefixLength = requiredLength
	}

	f.Prefix = p
}

// FileList is an ordered doubly linked list of FileNodes.
// Ordered by the FileNode.Prefix field.
type FileList struct {
	// Directory in which files are located
	Directory string

	// MaxPrefixLength is the maximum number of digits used to represent
	// FileNode.Prefix fields.
	MaxPrefixLength uint64

	// Head of linked list
	Head *FileNode
}

// Insert node in ordered position.
// If squash is set to true: will attempt to remove one from a FileNode.PrevDelta after the
// insert location. If no FileNodes with FileNode.PrevDelta > 0 can be found it will do nothing.
// This argument will only work if ComputeDeltas has been called on the list.
func (l *FileList) Insert(n *FileNode, squash bool) {
	// Update PrefixLength if new node's Prefix won't fit
	if l.MaxPrefixLength < n.PrefixLength {
		l.MaxPrefixLength = n.PrefixLength
	}
	
	// If no nodes in list yet
	if l.Head == nil {
		l.Head = n
		return
	}

	// Find place in list to insert node
	current := l.Head

	for current.Next != nil && current.Next.Prefix < n.Prefix {
		current = current.Next
	}

	// Insert at current
	// from: c <-> x
	// to  : c <-> n <-> x

	// Link n and x
	if current.Next != nil {
	    current.Next.Prev = n
	    n.Next = current.Next
	}

	// Link c and n
	current.Next = n
	n.Prev = current

	// Squash
	if squash {
		// Check if next node and inserted node have same prefix
		if current.Next.Next != nil && current.Next.Next.Prefix == n.Prefix {
			// If this is the case shift the PrevDelta to the newly inserted node
			n.PrevDelta = current.Next.Next.PrevDelta
			current.Next.Next.PrevDelta = 0
		}

		// If not last item in list
		if n.Next != nil { 
			for current = n.Next; current != nil; current = current.Next {
				if current.PrevDelta > 0 {
					current.PrevDelta--
					return
				}
			}
		}
	}
}

// ComputeDeltas traverses the list and sets the FileNode.PrevDelta field based on the
// nodes which are currently in the list.
func (l *FileList) ComputeDeltas() {
	for head := l.Head; head != nil; head = head.Next {
		// If first node
		if head.Prev == nil {
			// PrevDelta is the difference between 0 and Prefix
			head.PrevDelta = head.Prefix
		} else { // If node in middle
			if head.Prefix - head.Prev.Prefix > 0 {
			    head.PrevDelta = head.Prefix - head.Prev.Prefix - 1
			}
		}
	}
}

// ComputePrefixes re-compute's FileNode.Prefix values using FileNode.Prev delta fields.
// Existing FileNode.Prefix fields are ignored. Prefix values are computed based on the
// space between nodes.
func (l *FileList) ComputePrefixes() {
	var nextPrefix uint64 = 0

	for head := l.Head; head != nil; head = head.Next {
		nextPrefix += head.PrevDelta
		head.Prefix = nextPrefix
		nextPrefix++
	}
}


// String returns a string representation of the file list
func (l FileList) String() string {
	out := []string{}
	
	for head := l.Head; head != nil; head = head.Next {
		out = append(out, head.String())
	}

	return strings.Join(out, "\n")
}

// Map creates a map where keys are un-prefixed file names and values are FileNodes
// Note that the FileNodes returned are not pointers, but values.
func (l FileList) Map() map[string]FileNode {
	m := make(map[string]FileNode)

	for head := l.Head; head != nil; head = head.Next {
		m[head.Name] = *head
	}

	return m
}
