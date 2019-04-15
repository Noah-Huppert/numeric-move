package main

import (
	"fmt"
	"path"
	"os"
	"strconv"
	"path/filepath"

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

	logger.Debugf("%#v, %#v", target, newPrefix)

	// {{{1 Read all numerically prefixed files
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

		filesList.Insert(node)

		return nil
	})

	if err != nil {
		logger.Fatalf("failed to list numerically prefixed files: %s", err.Error())
	}

	if filesList.Head == nil {
		logger.Fatalf("found no numerically prefixed files")
	}


	for head := filesList.Head; head != nil; head = head.Next {
		logger.Debugf("%s", head.FullName(filesList.PrefixLength))
	}
}
