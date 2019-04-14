package file

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

// buildFile builds a test NumPrefixFile
func buildFile() *NumPrefixFile {
	return &NumPrefixFile{
		Directory: "/a/directory",
		UnPrefixedName: "_name.foo",
		NumPrefix: uint64(50),
		PrefixLength: uint64(4),
	}
}

// TestNew tests the NewNumPrefixFile function
func TestNew(t *testing.T) {
	actual, err := NewNumPrefixFile("/a/directory/0050_name.foo")

	assert.Nil(t, err)

	assert.Equal(t, "/a/directory", actual.Directory)
	assert.Equal(t, "_name.foo", actual.UnPrefixedName)
	assert.Equal(t, uint64(50), actual.NumPrefix)
	assert.Equal(t, uint64(4), actual.PrefixLength)
}

// TestFileName tests the FileName method
func TestFileName(t *testing.T) {
	f := buildFile()

	assert.Equal(t, "0050_name.foo", f.FileName())
}

// TestPath rests the Path method
func TestPath(t *testing.T) {
	f := buildFile()

	assert.Equal(t, "/a/directory/0050_name.foo", f.Path())
}
