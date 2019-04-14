package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Noah-Huppert/numeric-move/file"

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

	target, err := file.NewNumPrefixFile(os.Args[1])
	if err != nil {
		logger.Fatalf("failed to parse target argument: %s", err.Error())
	}

	newPrefix, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		logger.Fatalf("failed to parse new prefix argument: %s", err.Error())
	}

	//newPrefixLength := len(string(newPrefix))

	logger.Debugf("newPrefix: %d", newPrefix)

	// {{{1 Model target directory files
	// {{{2 Create NumPrefixFile for each file
	files, err := file.LoadDirectory(target.Directory)
	if err != nil {
		logger.Fatalf("failed to model target directory: %s", err.Error())
	}

	// {{{2 Build linked list
	/*filesHead := &file.NumPrefixFileNode{
		Type:               file.NodeTypeFile,
		FileUnPrefixedName: files[0].UnPrefixedName,
		Next:               nil,
		Prev:               nil,
	}
*/

	//currentNode := filesHead
	lastPrefix := files[0].NumPrefix

	for _, file := range files[1:] {
		// Check if there is a space between this file and the last file
		if file.NumPrefix-lastPrefix > 0 {
			// Add spacer node

		}
	}
}