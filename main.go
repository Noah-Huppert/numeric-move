package main

import (
	"fmt"
	"os"
	"strconv"
	"path"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/Noah-Huppert/golog"
)


// numericPrefixExp matches a file name with a numeric prefix
var numericPrefixExp *regexp.Regexp = regexp.MustCompile("^([0-9]+)(.*)$")

// Selector indicates the directory of a file and its numeric prefix
type Selector struct {
	// Directory in which file is located
	Directory string

	// NonNumericFileName is the name of the file without the numeric prefix
	NonNumericFileName string

	// PrefixLength length of numerical prefix
	PrefixLength uint32

	// PrefixNumber number in numerical prefix
	PrefixNumber uint64
}

// NewSelector creates a new select from a string
func NewSelector(s string) (Selector, error) {
	sel := Selector{}
	
	fName := path.Base(s)

	matches := numericPrefixExp.FindStringSubmatch(fName)

	if len(matches) == 0 {
		return sel, fmt.Errorf("not in numeric prefix format")
	}
	
	prefixNumber, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return sel, fmt.Errorf("error parsing numeric prefix string into int64: %s", err.Error())
	}	

	sel.Directory = path.Dir(s)
	sel.NonNumericFileName = matches[2]
	sel.PrefixLength = uint32(len(matches[1]))
	sel.PrefixNumber = prefixNumber

	return sel, nil
}

// FileName returns the name of a file described by a selector.
// The function will return a suggested selector PrefixLength if the selector's PrefixLength is too small to hold the PrefixNumber.
// If the selector's PrefixLength is adequate 0 will be returned.
func (s Selector) FileName(resizePrefix bool) (string, uint32) {
	out := ""
	
	numStr := string(s.PrefixNumber)

	newPrefixLength := 0

	if len(numStr) > s.PrefixLength {
		newPrefixLength = len(numStr)
	}

	for len(out) - len(numStr) < s.PrefixLength {
		out += "0"
	}

	out += numStr

	out += s.NonNumericFileName

	return out, newPrefixLength
}

func printUsage(progName string) {
	fmt.Printf("%s - Numeric move\n", progName)
	fmt.Printf("\n")
	fmt.Printf("Usage: %s FROM TO\n", progName)
	fmt.Printf("\n")
	fmt.Printf("Arguments:\n")
	fmt.Printf("    FROM    Numeric prefix of file to move\n")
	fmt.Printf("    TO      New numeric prefix\n")
}

func main() {
	// {{{1 Setup logger
	logger := golog.NewStdLogger("nmv")
	
	// {{{1 Get command line arguments
	if len(os.Args[1:]) != 2 {
		printUsage(os.Args[0])
		logger.Fatalf("2 arguments required")
	}

	from, err := NewSelector(os.Args[1])
	if err != nil {
		logger.Fatalf("failed to parse FROM argument: %s", err.Error())
	}

	to, err := NewSelector(os.Args[2])
	if err != nil {
		logger.Fatalf("failed to parse TO argument: %s", err.Error())
	}

	// {{{1 Model files in TO directory
	fromFileName, _ := from.FileName()
	
	toDirSelectors := []Selector{}

	err = filepath.Walk(to.Directory, func(wPath string, info os.FileInfo, err error) error {
		if info.IsDir() && wPath != to.Directory {
			return filepath.SkipDir
		}

		fileName := path.Base(wPath)

		if ! numericPrefixExp.Match([]byte(fileName)) {
			return nil
		}

		selector, err := NewSelector(wPath)
		if err != nil {
			return fmt.Errorf("failed to parse file into Selector: %s", err.Error())
		}

		toDirSelectors = append(toDirSelectors, selector)

		return nil
	})
	if err != nil {
		logger.Fatalf("failed to walk directory in TO argument: %s", err.Error())
	}

	sort.SliceStable(toDirSelectors, func(i, j int) bool {
	    return toDirSelectors[i].PrefixNumber < toDirSelectors[j].PrefixNumber
	})

	// {{{1 Determine files which we must rename in TO directory
	// {{{2 If resize
	if from.PrefixLength != to.PrefixLength {
		for i, _ := range toDirSelectors
	}

	logger.Debugf("%#v", toDirSelectors)
}
