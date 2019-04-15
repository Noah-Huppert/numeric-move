package main

import (
	"testing"
	"fmt"
	
	"github.com/stretchr/testify/assert"
)

// assertListEqual tests that a FileList holds the expected values.
// The expected argument is a string array representation of the FileList. Each item in the
// expected array should be in the format: <expected file name>:<expected PrevDelta>
func assertListEqual(t *testing.T, expected []string, actual *FileList) {
	actualArray := []string{}

	for head := actual.Head; head != nil; head = head.Next {
		actualArray = append(actualArray, fmt.Sprintf("%s:%d", head.Name, head.PrevDelta))
	}

	assert.ElementsMatch(t, expected, actualArray)
}

// TestFileListInsertUpdatesMaxPrefixLength ensures FileList.Insert updates the MaxPrefixLength
func TestFileListInsertUpdatesMaxPrefixLength(t *testing.T) {
	list := &FileList{}

	assert.Equal(t, uint64(0), list.MaxPrefixLength)

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: uint64(1),
		PrefixLength: 1,
	}, false)

	assert.Equal(t, uint64(1), list.MaxPrefixLength)

	list.Insert(&FileNode{
		Name: "_bar.ext",
		Prefix: uint64(10),
		PrefixLength: 2,
	}, false)

	assert.Equal(t, uint64(2), list.MaxPrefixLength)
}

// TestFileListInsertEmpty ensures FileList.Insert inserts a node in an empty list
func TestFileListInsertEmpty(t *testing.T) {
	list := &FileList{}

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: uint64(1),
		PrefixLength: 1,
		PrevDelta: 9,
	}, false)

	assertListEqual(t, []string{"_foo.ext:9"}, list)
}

// TestFileListInsertOrder ensures FileList.Insert inserts a node in the correct place in a list
func TestFileListInsertOrder(t *testing.T) {
	list := &FileList{}

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: uint64(1),
		PrefixLength: 1,
		PrevDelta: 9,
	}, false)

	list.Insert(&FileNode{
		Name: "_bar.ext",
		Prefix: uint64(3),
		PrefixLength: 1,
		PrevDelta: 5,
	}, false)

	list.Insert(&FileNode{
		Name: "_new.ext",
		Prefix: uint64(2),
		PrefixLength: 1,
		PrevDelta: 7,
	}, false)

	assertListEqual(t, []string{"_foo.ext:9", "_new.ext:7", "_bar.ext:5"}, list)
}
