package main

import (
	"strings"
	"fmt"
	"path"
	"os"
	"strconv"
	"path/filepath"
	"bufio"

	"github.com/Noah-Huppert/golog"
)

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

	targetDir := path.Dir(os.Args[1])
	target, err := NewFileNode(path.Base(os.Args[1]))

	if err != nil {
		logger.Fatalf("failed to parse target argument: %s", err.Error())
	}

	newPrefix, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		logger.Fatalf("failed to parse new prefix argument: %s", err.Error())
	}

	// {{{1 Read all numerically prefixed files
	// {{{2 Read
	filesList := &FileList{
		Directory: targetDir,
	}

	err = filepath.Walk(targetDir, func(wPath string, info os.FileInfo, err error) error {
		if info.IsDir() && wPath != targetDir {
			return filepath.SkipDir
		}

		node, err := NewFileNode(path.Base(wPath))

		if err != nil && err != ErrNotNumPrefixFile {
			return fmt.Errorf("error creating file node: %s", err.Error())
		} else if err == ErrNotNumPrefixFile {
			return nil
		}

		if node.Name == target.Name {
			return nil
		}

		filesList.Insert(node, false)

		return nil
	})

	if err != nil {
		logger.Fatalf("failed to list numerically prefixed files: %s", err.Error())
	}

	if filesList.Head == nil {
		logger.Fatalf("found no numerically prefixed files")
	}

	// {{{2 Place spaces in-betwen files
	filesList.ComputeDeltas()

	// {{{2 Make map of files before target insert
	beforeFiles := filesList.Map()

	// {{{1 Add target in new location
	// {{{{2 Insert
	newTarget := &FileNode{
		Name: target.Name,
		Prefix: target.Prefix,
		PrefixLength: target.PrefixLength,
	}
	newTarget.SetPrefix(newPrefix)
	filesList.Insert(newTarget, true)	

	// {{{2 Re-compute prefixes with target inserted
	filesList.ComputePrefixes()

	// {{{1 Determine which files need to be moved
	// toMove is a list of files to move, values are tuples in format (from, to)
	toMove := [][]string{}

	for after := filesList.Head; after != nil; after = after.Next {
		// If file is target
		if _, inBefore := beforeFiles[after.Name]; !inBefore {
			newTarget.PrefixLength = filesList.MaxPrefixLength 

			toMove = append(toMove, []string{
				target.FullName(),
				newTarget.FullName(),
			})
				
		} else {
			// If prefix has changed
			before, _ := beforeFiles[after.Name]
			after.PrefixLength = filesList.MaxPrefixLength

			if before.Prefix != after.Prefix || before.PrefixLength != after.PrefixLength {
				toMove = append(toMove, []string{
					before.FullName(),
					after.FullName(),
				})
			}
		}
	}

	// {{{1 Confirm moves with user
	logger.Infof("Directory: %s", filesList.Directory)
	logger.Infof("")
	logger.Infof("Files to be moved:")
	logger.Infof("")

	for _, m := range toMove {
		logger.Infof("%s -> %s", m[0], m[1])
	}

	logger.Infof("")
	logger.Infof("Proceed? [N/y]")
	reader := bufio.NewReader(os.Stdin)
	read, err := reader.ReadString('\n')
	if err != nil {
		logger.Fatalf("error reading user input: %s", err.Error())
	}

	if strings.ToLower(read) != "y" {
		logger.Fatalf("user did not confirm, exiting...")
		os.Exit(1)
	}
}
