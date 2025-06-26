package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	EOL             = "\n"
	BRANCHING_TRUNK = "├───"
	LAST_BRANCH     = "└───"
	TRUNC_TAB       = "│\t"
	LAST_TAB        = "\t"
	EMPTY_FILE      = "empty"
	ROOT_PREFIX     = ""
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTreeR(out, path, printFiles, "")
}

func dirTreeR(out io.Writer, path string, printFiles bool, prefix string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Failed to read directory: %w", err)
	}

	if !printFiles {
		entries = slices.DeleteFunc(entries, func(elem os.DirEntry) bool {
			return !elem.IsDir()
		})
	}

	slices.SortFunc(entries, func(i, j os.DirEntry) int {
		return strings.Compare(i.Name(), j.Name())
	})

	for i, entry := range entries {
		isLast := (i == len(entries)-1)
		if isLast {
			fmt.Fprintf(out, prefix+LAST_BRANCH)
		} else {
			fmt.Fprintf(out, prefix+BRANCHING_TRUNK)
		}

		fmt.Fprint(out, entry.Name())

		if entry.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += LAST_TAB
			} else {
				newPrefix += TRUNC_TAB
			}

			fullPath := filepath.Join(path, entry.Name())
			fmt.Fprint(out, EOL)
			if err := dirTreeR(out, fullPath, printFiles, newPrefix); err != nil {
				return err
			}
		} else if printFiles {
			info, err := entry.Info()
			if err != nil {
				return fmt.Errorf("Unable to get FileInfo: %w", err)
			}

			if info.Size() == 0 {
				fmt.Fprint(out, " ("+EMPTY_FILE+")")
			} else {
				fmt.Fprintf(out, " (%db)", info.Size())
			}
			fmt.Fprint(out, EOL)
		}
	}
	return nil
}

func main() {
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage: go run main.go . [-f]")
	}

	out := os.Stdout
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
