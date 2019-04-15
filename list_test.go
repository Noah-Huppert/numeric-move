package main

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

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
