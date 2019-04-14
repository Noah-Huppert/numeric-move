package file

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"sort"
	"os"
)

// numPrefixFileExp is a regular expression which matches against a numerically prefixed file
var numPrefixFileExp *regexp.Regexp = regexp.MustCompile("^([0-9]+)(.*)$")

// NumPrefixFile holds information about a numerically prefixed file
type NumPrefixFile struct {
	// Directory of file
	Directory string

	// UnPrefixedName is the file's name without the numerical prefix
	UnPrefixedName string

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
		Directory:          path.Dir(absPath),
		UnPrefixedName: matches[2],
		NumPrefix:          num,
		PrefixLength:       uint64(len(matches[1])),
	}, nil
}

// FileName returns the prefixed name of a file. Assumes PrefixLength can fit NumPrefix.
func (f NumPrefixFile) FileName() string {
	s := ""
	numStr := strconv.FormatUint(f.NumPrefix, 10)

	for uint64(len(s)+len(numStr)) < f.PrefixLength {
		s += "0"
	}

	return s + numStr + f.UnPrefixedName
}

// Path is the file's full path. Assumes PrefixLength can fit NumPrefix.
func (f NumPrefixFile) Path() string {
	return filepath.Join(f.Directory, f.FileName())
}

// LoadDirectory returns a list sorted by NumPrefix of all the numerically prefixed files in a directory
func LoadDirectory(dir string) ([]*NumPrefixFile, error) {
	files := []*NumPrefixFile{}

	err := filepath.Walk(dir, func(wPath string, info os.FileInfo, err error) error {
		if info.IsDir() && wPath != dir {
			return filepath.SkipDir
		}

		fileName := path.Base(wPath)

		if !numPrefixFileExp.Match([]byte(fileName)) {
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
		return nil, fmt.Errorf("failed to walk directory: %s", err.Error())
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].NumPrefix < files[j].NumPrefix
	})

	return files, nil
}

// BuildArray builds an array of NumPrefixFiles from a linked list of NumPrefixFileNodes
func BuildArray(dir string, length uint64, head *NumPrefixFileNode) ([]*NumPrefixFile, error) {
	if head == nil {
		return nil, fmt.Errorf("empty linked list")
	}

	a := []*NumPrefixFile{}

	nextNum := uint64(0)
	current := head

	for current != nil {
		if current.Type == NodeTypeSpace {
			nextNum += current.SpaceAmount
		} else {
			a = append(a, &NumPrefixFile{
				Directory: dir,
				UnPrefixedName: current.FileUnPrefixedName,
				NumPrefix: nextNum,
				PrefixLength: length,
			})
			nextNum++
		}

		current = current.Next
	}

	return a, nil
}
