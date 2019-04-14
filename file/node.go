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

// InsertAfter adds a NumPrefixFileNode after the current.
// If inserting near a spacer node it will decrease the size of this node by 1 if possible
func (n *NumPrefixFileNode) InsertAfter(i *NumPrefixFileNode) {
	// If we are adding on to the end of the list
	if n.Next == nil {
		// Collapse space if possible
		if n.Prev != nil && n.Prev.Type == NodeTypeSpace {
			s := n.Prev

			if s.SpaceAmount > 1 {
				s.SpaceAmount--
			}
		}

		// from: x <-> n
		// to  : x <-> n <-> i
		n.Next = i
		i.Prev = n

	} else { // If inserting in the middle
		// Collapse space if possible
		// Favore space before the node
		collapsed := false

		if n.Prev != nil && n.Prev.Type == NodeTypeSpace {
			s := n.Prev

			if s.SpaceAmount > 1 {
				s.SpaceAmount--
				collapsed = true
			}
		}

		if ! collapsed && n.Next.Type == NodeTypeSpace {
			s := n.Next

			if s.SpaceAmount > 1 {
				s.SpaceAmount--
			}
		}
		
		// from: n <-> x
		// to  : n <-> i <-> x
		n.Next.Prev = i
		i.Next = n.Next
		i.Prev = n
		n.Next = i
	}
}

// BuildFromArray creates a linked list from an array of NumPrefixFile-s
/*
func BuildFromArray(a []*NumPrefixFile) (*NumPrefixFileNode, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("array cannot be empty")
	}

	var head *NumPrefixFileNode = nil

	for _, f := range a {
		if head == nil {
		}
	}
}
*/
