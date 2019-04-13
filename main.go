package main

import (
	"os"
	"fmt"
	"path"
	"regexp"
	"path/filepath"
	"strconv"
	"sort"

	"github.com/Noah-Huppert/golog"
)

// numPrefixFileExp is a regular expression which matches against a numerically prefixed file
var numPrefixFileExp *regexp.Regexp = regexp.MustCompile("^([0-9]+)(.*)$")

// NumPrefixFile holds information about a numerically prefixed file
type NumPrefixFile struct {
	// Directory of file
	Directory string

	// UnPrefixedFileName is the file's name without the numerical prefix
	UnPrefixedFileName string

	// NumPrefix is the file's numerical prefix
	NumPrefix uint64

	// PrefixLength is the number of digits used to display a numerical prefix
	PrefixLength uint64
}	

// NewNumPrefixFile creates a new NumPrefixFile from a path
func NewNumPrefixFile(s string) (*NumPrefixFile, error) {
	// Check if s is a path to a numerically prefixed file	
	fName := path.Base(s)
	matches := numPrefixFileExp.FindStringSubmatch(fName)
	
	if len(matches) == 0 {
		return nil, fmt.Errorf("path is not a numerically prefixed file")
	}

	// Resolve to absolute path
	absPath, err := filepath.Abs(s)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path: %s", err.Error())
	}

	// Parse numerical prefix
	num, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse numeric prefix: %s", err.Error())
	}

	return &NumPrefixFile{
		Directory: path.Dir(absPath),
		UnPrefixedFileName: matches[2],
		NumPrefix: num,
		PrefixLength: uint64(len(matches[1])),
	}, nil
}

// Resize returns a new PrefixLength if the current one will not fit NumPrefix.
// Otherwise 0 is returned.
func (f NumPrefixFile) Resize() uint64 {
	l := uint64(len(string(f.NumPrefix)))

	if l > f.PrefixLength {
		return l
	}

	return 0
}

// FileName returns the prefixed name of a file. Assumes PrefixLength can fit NumPrefix.
func (f NumPrefixFile) FileName() string {
	s := ""
	numStr := string(f.NumPrefix)

	for uint64(len(s) + len(numStr)) < f.PrefixLength {
		s += "0"
	}

	return s + numStr
}

// Path is the file's full path. Assumes PrefixLength can fit NumPrefix.
func (f NumPrefixFile) Path() string {
	return filepath.Join(f.Directory, f.FileName())
}

func main() {
	// {{{1 Setup logger
	logger := golog.NewStdLogger("nmv")

	// {{{1 Get arguments
	if len(os.Args[1:]) != 2 {
		fmt.Printf("%s - Numeric move\n\n", os.Args[0])
		fmt.Printf("Usage: %s TARGET NEW_PREFIX\n\n", os.Args[0])
		fmt.Println("Arguments:")
		fmt.Println("    TARGET        Path of file to move")
		fmt.Println("    NEW_PREFIX    New numerical prefix")
	}

	target, err := NewNumPrefixFile(os.Args[1])
	if err != nil {
		logger.Fatalf("failed to parse target argument: %s", err.Error())
	}

	newPrefix, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		logger.Fatalf("failed to parse new prefix argument: %s", err.Error())
	}
	logger.Debugf("newPrefix: %d", newPrefix)

	// {{{1 Model target directory files
	files := []*NumPrefixFile{}

	err = filepath.Walk(target.Directory, func(wPath string, info os.FileInfo, err error) error {
		if info.IsDir() && wPath != target.Directory {
			return filepath.SkipDir
		}

		fileName := path.Base(wPath)

		if ! numPrefixFileExp.Match([]byte(fileName)) {
			return nil
		}

		f, err := NewNumPrefixFile(wPath)
		if err != nil {
			return fmt.Errorf("failed to parse file into num prefix file: %s", err.Error())
		}

		files = append(files, f)

		return nil
	})

	if err != nil {
		logger.Fatalf("failed to walk target directory: %s", err.Error())
	}

	sort.SliceStable(files, func(i, j int) bool {
	    return files[i].NumPrefix < files[j].NumPrefix
	})

	for _, f := range files {
		logger.Debugf("%#v", f)
	}

	// {{{1 Augment files for new target
	// TODO: Something with a double loop
}
