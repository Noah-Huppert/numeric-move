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

	// {{{1 Convert to space based representation
	// {{{2 First node
	var spaceList *SpaceNode = nil

	if filesList.Head.Name == target.Name { // If target node
		// Replace with space
		spaceList = &SpaceNode{
			IsSpace: true,
			Space: 1,
		}
	} else if filesList.Head.Prefix == 0 { // If no space at beginning of list
		spaceList = &SpaceNode{
			IsSpace: false,
			File: filesList.Head.Name,
			MetaFilePrefix: filesList.Head.Prefix,
		}
	} else { // Space at beginning of list
		spaceList = &SpaceNode{
			IsSpace: true,
			Space: filesList.Head.Prefix,
		}
	}

	// {{{2 Add rest of list
	currentS := spaceList
	
	for currentF := filesList.Head.Next; currentF != nil; currentF = currentF.Next {
		// If space required
		if currentF.Prefix - currentF.Prev.Prefix > 0 {
			space := &SpaceNode{
				IsSpace: true,
				Space: currentF.Prefix - currentF.Prev.Prefix - 1,
			}

			currentS.Next = space
			space.Prev = currentS
			currentS = space
		}

		// If target, replace with space
		if currentF.Name == target.Name {
			space := &SpaceNode{
				IsSpace: true,
				Space: 1,
			}
			currentS.Next = space
			space.Prev = currentS
			currentS = space
		} else {
		    // Insert file
		    file := &SpaceNode{
			    IsSpace: false,
			    File: currentF.Name,
			    MetaFilePrefix: currentF.Prefix,
		    }

		    currentS.Next = file
		    file.Prev = currentS
		    currentS = file
		}
	}

	// TODO: Cleanup SpaceNode to be neater
	// TODO: Fix the target movement logic

	// {{{1 Insert target into space list with new prefix
	for currentS = spaceList; currentS != nil; currentS = currentS.Next {
		if currentS.IsSpace {
			continue
		}

		if currentS.MetaFilePrefix > newPrefix {
			node := &SpaceNode{
				IsSpace: false,
				File: target.Name,
			}

			if currentS.Next != nil {
				// from: c <-> x
				// to  : c <-> n <-> x
				currentS.Next.Prev = node
				node.Next = currentS.Next
			}

			node.Prev = currentS
			currentS.Next = node

			break
		}
	}

	// {{{1 Rebuild file names from space list
	filesList = FilesFromSpaces(filesList.Directory, filesList.PrefixLength, spaceList)

	for head := filesList.Head; head != nil; head = head.Next {
		logger.Debugf("%s", head.FullName(filesList.PrefixLength))
	}
}
