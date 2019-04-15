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
