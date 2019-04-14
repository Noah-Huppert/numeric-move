package file

import (
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"
)

// assertLinkedEqual asserts that nodes appear in the correct order.
// The expected array indicates the orders node appear. This array represents nodes as strings.
// The FileUnPrefixedName is used for the string value for nodes of Type == NodeTypeFile.
// A string representation of SpaceAmount is used for the string value for nodes of Type == NodeTypeSpace.
func assertLinkedEqual(t *testing.T, expected []string, actualHead *NumPrefixFileNode) {
	actual := []string{}
	current := actualHead

	for current != nil {
		if current.Type == NodeTypeFile {
			actual = append(actual, current.FileUnPrefixedName)
		} else {
			actual = append(actual, strconv.FormatUint(current.SpaceAmount, 10))
		}

		current = current.Next
	}

	assert.ElementsMatch(t, expected, actual)
}

// buildTestList createsa doubley linked list in the format: a <-> 2 <-> b <-> 1 <-> c.
// See the assertLinkedEqual documentation for an explinaiton of the string format about.
func buildTestList() *NumPrefixFileNode {
	a := &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "a",
	}
	space2 := &NumPrefixFileNode{
		Type: NodeTypeSpace,
		SpaceAmount: uint64(2),
	}
	b := &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "b",
	}
	space1 := &NumPrefixFileNode{
		Type: NodeTypeSpace,
		SpaceAmount: uint64(1),
	}
	c := &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "c",
	}

	a.Next = space2
	space2.Prev = a

	space2.Next = b
	b.Prev = space2

	b.Next = space1
	space1.Prev = b

	space1.Next = c
	c.Prev = space1

	return a
}

// buildInsertNode builds a test file node with a FileUnPrefixedName of i
func buildInsertNode() *NumPrefixFileNode {
	return &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "i",
	}
}

// TestInsertAfterHead tests inserting a node after the head
func TestInsertAfterHead(t *testing.T) {
	head := buildTestList()

	head.InsertAfter(buildInsertNode())

	assertLinkedEqual(t, []string{"a", "i", "1", "b", "1", "c"}, head)
}

// TestInsertAfterMiddle tests insertng a node after the middle of a list
func TestInsertAfterMiddle(t *testing.T) {
	head := buildTestList()

	head.Next.Next.InsertAfter(buildInsertNode())

	assertLinkedEqual(t, []string{"a", "1", "b", "1", "i", "c"}, head)
}

// TestInsertAfterTail tests inserting a node after the tail of a list
func TestInsertAfterTail(t *testing.T) {
	head := buildTestList()

	head.Next.Next.Next.Next.InsertAfter(buildInsertNode())

	assertLinkedEqual(t, []string{"a", "2", "b", "1", "c", "i"}, head)
}
