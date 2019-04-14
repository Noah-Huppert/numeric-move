package file

import (
	"fmt"
)

// NumPrefixFileNode is a node in a linked list which represents all numerically prefixed files in a directory
// Representing these files as a linked list allows for easier shifting of prefixes.
// Some nodes represent actual files. Some represent spaces in the prefix numbers.
type NumPrefixFileNode struct {
	// Type indicates the information the node stores
	Type NumPrefixFileNodeType

	// FileUnPrefixedName is the name of a file without its numeric prefix
	FileUnPrefixedName string

	// SpaceAmount is the amount of space between numeric prefixes
	SpaceAmount uint64

	// Next is a pointer to the next NumPrefixFileNode, nil if end of list
	Next *NumPrefixFileNode

	// Prev is the pointer to the previous NumPrefixFileNode, nil if beginning of list
	Prev *NumPrefixFileNode
}

// NumPrefixFileNodeType is the type of a NumPrefixFileNode. Depending on this type value different fields in the
// NumPrefixFileNode will be set.
type NumPrefixFileNodeType uint32

const (
	// NodeTypeFile indicates that the node stores information about a file. Fields prefixed with File will be set.
	NodeTypeFile NumPrefixFileNodeType = 0

	// NodeTypeSpace indicates the the node stores information about a space between files. Fields prefixed with Space will be set.
	NodeTypeSpace NumPrefixFileNodeType = 1
)

// InsertAfter adds a NumPrefixFileNode after the current
func (n *NumPrefixFileNode) InsertAfter(i *NumPrefixFileNode) {
	// If we are adding on to the end of the list
	if n.Next == nil {
		// from: x <-> n
		// to  : x <-> n <-> i
		n.Next = i
		i.Prev = n
	} else { // If inserting in the middle
		// from: n <-> x
		// to  : n <-> i <-> x
		n.Next.Prev = i
		i.Next = n.Next
		i.Prev = n
		n.Next = i
	}
}

// String returns a string representation of a node
func (n NumPrefixFileNode) String() string {
	if n.Type == NodeTypeFile {
		return fmt.Sprintf("file: %s", n.FileUnPrefixedName)
	} else {
		return fmt.Sprintf("space: %d", n.SpaceAmount)
	}
}

// BuildList builds a linked list of NumPrefixFileNodes from an array of NumPrefixFiles
func BuildList(a []*NumPrefixFile) (*NumPrefixFileNode, error) {
	// Check not empty
	if len(a) == 0 {
		return nil, fmt.Errorf("array empty")
	}

	// Add first node
	var head *NumPrefixFileNode

	fileNode := &NumPrefixFileNode{
	    Type: NodeTypeFile,
	    FileUnPrefixedName: a[0].UnPrefixedName,
	}
	
	if a[0].NumPrefix == 0 {
		head = fileNode
	} else {
		head = &NumPrefixFileNode{
			Type: NodeTypeSpace,
			SpaceAmount: a[0].NumPrefix - 1,
			Next: fileNode,
		}
		fileNode.Prev = head
	}

	currentF := a[0]
	current := fileNode

	for _, f := range a[1:] {
		if f.NumPrefix - 1 != currentF.NumPrefix {
			current.InsertAfter(&NumPrefixFileNode{
				Type: NodeTypeSpace,
				SpaceAmount: (f.NumPrefix - currentF.NumPrefix) - 1,
			})

			current = current.Next
		}

		current.InsertAfter(&NumPrefixFileNode{
			Type: NodeTypeFile,
			FileUnPrefixedName: f.UnPrefixedName,
		})

		current = current.Next
		currentF = f
	}

	return head, nil
}
