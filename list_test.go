package main

import (
	"testing"
	"fmt"
	
	"github.com/stretchr/testify/assert"
)

// assertListEqual tests that a FileList holds the expected values.
// The expected argument is a string array representation of the FileList. Each item in the
// expected array should be in the format: <Prefix>:<Name>:<PrevDelta>
func assertListEqual(t *testing.T, expected []string, actual *FileList) {
	// Traverse forwards
	actualForward := []string{}

	for head := actual.Head; head != nil; head = head.Next {
		actualForward = append(actualForward,
			fmt.Sprintf("%d:%s:%d", head.Prefix, head.Name, head.PrevDelta))
	}

	assert.ElementsMatch(t, expected, actualForward)

	// Traverse backwards
	actualBackwards := []string{}
	
	// Get to end of list
	tail := actual.Head

	for ; tail.Next != nil; tail = tail.Next {}

	// Traverse backwards
	for ; tail != nil; tail = tail.Prev {
		actualBackwards = append(actualBackwards,
			fmt.Sprintf("%d:%s:%d", tail.Prefix, tail.Name, tail.PrevDelta))
	}

	// Reverse
	actualBackwardsForwards := []string{}

	for i := len(actualBackwards) - 1; i >= 0; i-- {
		actualBackwardsForwards = append(actualBackwardsForwards, actualBackwards[i])
	}

	assert.ElementsMatch(t, expected, actualBackwardsForwards)
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

	assertListEqual(t, []string{"1:_foo.ext:9"}, list)
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

	assertListEqual(t, []string{"1:_foo.ext:9", "2:_new.ext:7", "3:_bar.ext:5"}, list)
}

