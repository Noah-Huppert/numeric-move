package main

import (
	"fmt"
	"regexp"
	"strconv"
	"errors"
	"strings"
)

// numPrefixExp matches a numerically prefixed file.
// The first match group is the numeric prefix.
// The second match group is the un-prefixed file name.
var numPrefixExp *regexp.Regexp = regexp.MustCompile("^([0-9]+)(.*)$")

// FileNode is a node in a doubly linked list.
// It holds information about a numerically prefixed file.
type FileNode struct {
	// Name is the file's name without the numeric prefix
	Name string

	// Prefix is the file's numeric prefix
	Prefix uint64

	// PrefixLength is the number of digits used to represent the prefix.
	// If the Prefix can be represented in less the extra digits are used
	// by placing 0's before the prefix.
	PrefixLength uint64

	// PrevDelta is the difference between the previous node's prefix and Prefix.
	// This field will start un-set. Once all files have been added to a FileList
	// the FileList will traverse itself and set this field. This ensures that
	// PrevDelta is accurate.
	PrevDelta uint64

	// Next FileNode in list
	Next *FileNode

	// Prev FileNode in list
	Prev *FileNode
}

// ErrNotNumPrefixFile indicates a file is not numerically prefixed
var ErrNotNumPrefixFile error = errors.New("not a numerically prefixed file")

// NewFileNode creates a new FileNode from a path.
// The given path should just be a file name without any leading directory.
// Returns ErrNoNumPrefixFile if the path points to a non-numericaly prefixed file.
func NewFileNode(p string) (*FileNode, error) {
	// Check that the path is for a numerically prefixed file
	matches := numPrefixExp.FindStringSubmatch(p)

	if len(matches) == 0 {
		return nil, ErrNotNumPrefixFile
	}

	// Parse prefix
	prefix, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse numeric prefix to unsigned integer: %s", err.Error())
	}

	return &FileNode{
		Name: matches[2],
		Prefix: prefix,
		PrefixLength: uint64(len(matches[1])),
		PrevDelta: 0,
	}, nil
}

// String representation
func (f FileNode) String() string {
	return fmt.Sprintf("%d %s (dt: %d)", f.Prefix, f.Name, f.PrevDelta)
}

// FullName returns the file's name with its numeric prefix
func (f FileNode) FullName() string {
	prefixStr := strconv.FormatUint(f.Prefix, 10)
	zeros := strings.Repeat("0", int(f.PrefixLength) - len(prefixStr))

	return zeros + prefixStr + f.Name
}

// SetPrefix sets the Prefix field and updates PrefixLength if it is too small to hold the new prefix
func (f *FileNode) SetPrefix(p uint64) {
	requiredLength := uint64(len(strconv.FormatUint(p, 10)))

	if f.PrefixLength < requiredLength {
		f.PrefixLength = requiredLength
	}

	f.Prefix = p
}