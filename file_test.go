package main

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

const (
	name string = "_name.ext"
	prefix uint64 = 54
	prefixLength uint64 = 4
	fullName string = "0054_name.ext"

	newNoResizePrefix uint64 = 70

	newResizePrefix uint64 = 10001
	newPrefixLength uint64 = 5
)

// TestNewFileNode ensures NewFileNode sets all fields appropriately
func TestNewFileNode(t *testing.T) {
	f, err := NewFileNode(fullName)

	assert.Nil(t, err)
	assert.Equal(t, f.Name, name)
	assert.Equal(t, f.Prefix, prefix)
	assert.Equal(t, f.PrefixLength, prefixLength)
}

// TestNewFileNodeErrs ensures NewFileNode errors when given a non-numerically prefixed file name
func TestNewFileNodeErrs(t *testing.T) {
	_, err := NewFileNode("this is not numerically prefixed")

	assert.Equal(t, err, ErrNotNumPrefixFile)
}


// TestFileNodeFullName ensures FileNode.FullName returns the correct value
func TestFileNodeFullName(t *testing.T) {
	f, err := NewFileNode(fullName)
	assert.Nil(t, err)

	assert.Equal(t, f.FullName(), fullName)
}

// TestFileNodeSetPrefixNoResize ensure FileNode.SetPrefix works when no resize is required
func TestFileNodeSetPrefixNoResize(t *testing.T) {
	f, err := NewFileNode(fullName)
	assert.Nil(t, err)

	f.SetPrefix(newNoResizePrefix)

	assert.Equal(t, f.Prefix, newNoResizePrefix)
	assert.Equal(t, f.PrefixLength, prefixLength)
}

// TestFileNodeSetPrefixResize ensures FileNode.SetPrefix works when a resize is required
func TestFileNodeSetPrefixResize(t *testing.T) {
	f, err := NewFileNode(fullName)
	assert.Nil(t, err)

	f.SetPrefix(newResizePrefix)

	assert.Equal(t, f.Prefix, newResizePrefix)
	assert.Equal(t, f.PrefixLength, newPrefixLength)
}
