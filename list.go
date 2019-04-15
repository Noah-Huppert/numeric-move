package main

import (	
	"strings"
)

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
		// Squash surrounding spaces if not last item
		if n.Next != nil {
		    // If the next node has PrevDelta, adjust its PrevDelta based on the inserted
		    // node's prefix
		    if n.Next.PrevDelta > 0 {
			    if n.Next.Prefix != n.Prefix {
				    n.Next.PrevDelta = n.Next.Prefix - n.Prefix - 1
			    } else {
				    n.Next.PrevDelta = 0
			    }

			    if n.Prev != nil && n.Prev.Prefix != n.Prefix {
				    n.PrevDelta = n.Prefix - n.Prev.Prefix - 1
			    } else {
				    n.Prev.PrevDelta = 0
			    }

			    for current = n.Next; current != nil; current = current.Next {
				    if current.PrevDelta > 0 {
					    current.PrevDelta--
					    return
				    }
			    }
		    }
		} else { // If last item, ensure PrevDelta is set correctly
			if n.Prev.Prefix != n.Prefix {
				n.PrevDelta = n.Prefix - n.Prev.Prefix - 1
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
