package main

import (
	"fmt"
	"regexp"
	"strconv"
	"errors"
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
	}, nil
}

// FullName returns the file's name with its numeric prefix
func (f FileNode) FullName(length uint64) string {
	zeros := ""
	prefixStr := strconv.FormatUint(f.Prefix, 10)

	for uint64(len(zeros) + len(prefixStr)) < length {
		zeros += "0"
	}

	return zeros + prefixStr + f.Name
}

// FileList is an ordered doubly linked list of FileNodes.
// Ordered by the FileNode.Prefix field.
type FileList struct {
	// Directory in which files are located
	Directory string

	// PrefixLength is the number of digits used to represent FileNode.Prefix fields.
	// If a number takes less than PrefixLength digits it will be prefixed with 0's.
	PrefixLength uint64

	// Head of linked list
	Head *FileNode
}

// Insert node in ordered position.
func (l *FileList) Insert(n *FileNode) {
	// Update PrefixLength if new node's Prefix won't fit
	requiredLength := uint64(len(strconv.FormatUint(n.Prefix, 10)))

	if l.PrefixLength < requiredLength {
		l.PrefixLength = requiredLength
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
}

// FilesFromSpaces creates a new files list from a spaces list
func FilesFromSpaces(dir string, length uint64, head *SpaceNode) *FileList {
	filesList := &FileList{
		Directory: dir,
		PrefixLength: length,
	}

	nextPrefix := uint64(0)
	
	for head = head; head != nil; head = head.Next {
		// If space node
		if head.IsSpace {
			nextPrefix += head.Space
		} else {
			filesList.Insert(&FileNode{
				Name: head.File,
				Prefix: nextPrefix,
			})

			nextPrefix++
		}
	}

	return filesList
}
