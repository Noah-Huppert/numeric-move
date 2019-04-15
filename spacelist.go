package main

// SpaceNode stores information about a file or a space between files.
type SpaceNode struct {
	// IsSpace is true if the node holds information about space between files.
	// In this case the Space field is valid.
	// Otherwise the node holds informaion about a file and the File field is valid.
	IsSpace bool

	// Space the amount of space between file's numeric prefixes
	Space uint64

	// File name
	File string

	// MetaFilePrefix is the prefix of the file when it was put in the list
	MetaFilePrefix uint64

	// Next node
	Next *SpaceNode

	// Prev node
	Prev *SpaceNode
}
