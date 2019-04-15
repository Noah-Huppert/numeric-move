package main

import (
	"testing"
	"fmt"
	"strings"
	"strconv"
	
	"github.com/stretchr/testify/assert"
)

// assertStrArrayEqual asserts that a string array is equal
func assertStrArrayEqual(t *testing.T, expected []string, actual []string) {
	// Check lengths
	assert.Equal(t, len(expected), len(actual), "array lengths do not match")

	// Check elements
	notMatchingIndices := []int{}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			notMatchingIndices = append(notMatchingIndices, i)
		}
	}

	if len(notMatchingIndices) > 0 {
		w := "index"
		if len(notMatchingIndices) > 1 {
			w = "indices"
		}

		iStr := []string{}

		for _, v := range notMatchingIndices {
			iStr = append(iStr, strconv.FormatInt(int64(v), 10))
		}

		t.Errorf("values at %s [%s] not equal \n expected: [%s] \n actual  : [%s]\n\n",
			w,
			strings.Join(iStr, ", "),
			strings.Join(expected, ", "),
			strings.Join(actual, ", "))
	}
}

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

	assertStrArrayEqual(t, expected, actualForward)

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

	assertStrArrayEqual(t, expected, actualBackwardsForwards)
}

// TestFLInsertUpdatesMaxPrefixLength ensures FileList.Insert updates the MaxPrefixLength
func TestFLInsertUpdatesMaxPrefixLength(t *testing.T) {
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

// TestFLInsertEmpty ensures FileList.Insert inserts a node in an empty list
func TestFLInsertEmpty(t *testing.T) {
	list := &FileList{}

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: uint64(1),
		PrefixLength: 1,
		PrevDelta: 9,
	}, false)

	assertListEqual(t, []string{"1:_foo.ext:9"}, list)
}

// TestFLInsertOrder ensures FileList.Insert inserts a node in the correct place in a list
func TestFLInsertOrder(t *testing.T) {
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

// TestFLInsertSquashNotEq ensures FileList.Insert squash == true modifies surrounding nodes PrevDelta when Prefixes are not equal
func TestFLInsertSquashNotEq(t *testing.T) {
	list := &FileList{}

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: 0,
		PrefixLength: 1,
		PrevDelta: 0,
	}, true)

	list.Insert(&FileNode{
		Name: "_bar.ext",
		Prefix: 10,
		PrefixLength: 2,
		PrevDelta: 9,
	}, true)

	list.Insert(&FileNode{
		Name: "_new.ext",
		Prefix: 5,
		PrefixLength: 1,
		PrevDelta: 0,
	}, true)

	assertListEqual(t, []string{"0:_foo.ext:0", "5:_new.ext:4", "10:_bar.ext:4"}, list)
}

// TestFLInsertSquashEq ensures FileList.Insert squash == true modifies surrounding nodes PrevDelta when Prefixes are equal
func TestFLInsertSquashEq(t *testing.T) {
	list := &FileList{}

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: 0,
		PrefixLength: 1,
		PrevDelta: 0,
	}, true)

	list.Insert(&FileNode{
		Name: "_bar.ext",
		Prefix: 10,
		PrefixLength: 2,
		PrevDelta: 9,
	}, true)

	list.Insert(&FileNode{
		Name: "_new.ext",
		Prefix: 10,
		PrefixLength: 2,
		PrevDelta: 0,
	}, true)

	assertListEqual(t, []string{"0:_foo.ext:0", "10:_new.ext:9", "10:_bar.ext:0"}, list)
}

// TestFLInsertSquashLast ensures FileList.Insert squash == true sets the inserted node's PrevDelta correctly if last node
func TestFLInsertSquashLast(t *testing.T) {
	list := &FileList{}

	list.Insert(&FileNode{
		Name: "_foo.ext",
		Prefix: 0,
		PrefixLength: 1,
		PrevDelta: 0,
	}, true)

	list.Insert(&FileNode{
		Name: "_bar.ext",
		Prefix: 5,
		PrefixLength: 1,
		PrevDelta: 4,
	}, true)

	list.Insert(&FileNode{
		Name: "_new.ext",
		Prefix: 10,
		PrefixLength: 2,
		PrevDelta: 0,
	}, true)

	assertListEqual(t, []string{"0:_foo.ext:0", "5:_bar.ext:4", "10:_new.ext:4"}, list)
}
