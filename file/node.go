package file

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
