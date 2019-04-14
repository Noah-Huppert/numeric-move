package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertLinkedEqual asserts that the values in the expected array appear in that order in the actualHead list.
// The expected array will hold FileUnPrefixedName values
func assertLinkedEqual(t *testing.T, expected []string, actualHead *NumPrefixFileNode) {
	actual := []string{}
	current := actualHead

	for current != nil {
		actual = append(actual, current.FileUnPrefixedName)
		current = current.Next
	}

	assert.ElementsMatch(t, expected, actual)
}

// BuildTestList createsa doubley linked list in the format: a <-> b <-> c.
// Where the labels above are the FileUnPrefixedName field values.
func buildTestList() *NumPrefixFileNode {
	a := &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "a",
	}
	b := &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "b",
	}
	c := &NumPrefixFileNode{
		Type:               NodeTypeFile,
		FileUnPrefixedName: "c",
	}

	a.Next = b
	b.Prev = a
	b.Next = c
	c.Prev = b

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

	assertLinkedEqual(t, []string{"a", "i", "b", "c"}, head)
}

// TestInsertAfterMiddle tests insertng a node after the middle of a list
func TestInsertAfterMiddle(t *testing.T) {
	head := buildTestList()

	head.Next.InsertAfter(buildInsertNode())

	assertLinkedEqual(t, []string{"a", "b", "i", "c"}, head)
}

// TestInsertAfterTail tests inserting a node after the tail of a list
func TestInsertAfterTail(t *testing.T) {
	head := buildTestList()

	head.Next.Next.InsertAfter(buildInsertNode())

	assertLinkedEqual(t, []string{"a", "b", "c", "i"}, head)
}

// TestInsertBeforeHead tests inserting a node before the head of a list
func TestInsertBeforeHead(t *testing.T) {
	head := buildTestList()

	head.InsertBefore(buildInsertNode())

	assert.NotNil(t, head.Prev)

	assertLinkedEqual(t, []string{"i", "a", "b", "c"}, head.Prev)
}

// TestInsertBeforeMiddle tests inserting a node before the middle of a list
func TestInsertBeforeMiddle(t *testing.T) {
	head := buildTestList()

	head.Next.InsertBefore(buildInsertNode())

	assertLinkedEqual(t, []string{"a", "i", "b", "c"}, head)
}
