package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func parseArgs() string {
	if len(os.Args) <= 1 {
		panic("Wrong usage: search your_search")
	}

	var search string

	for argIndex, arg := range os.Args {
		if argIndex != 0 {
			search += arg + " "
		}
	}

	return strings.Trim(search, " ")
}

func isSupportedFile(path string) bool {
	validSuffixes := []string{".txt", ".pdf", ".doc", ".docx", ".png", ".jpeg", ".webp", ".go", ".c", ".cpp", ".h", ".hpp", ".deb"}
	isSupported := false

	for _, suffix := range validSuffixes {
		isSupported = strings.HasSuffix(path, suffix)
	}

	return isSupported
}

func search(searchContent string) int {
	fmt.Printf("Searching for %v\n", searchContent)
	root := "/home/namu"
	var resultCount int

	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Erreur lors de l'accès au fichier %q : %v\n", path, err)
			return err
		}

		if strings.Contains(path, searchContent) && isSupportedFile(path) {
			resultCount += 1
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return resultCount
}

func main() {
	searchContent := parseArgs()
	resultCount := search(searchContent)
	if resultCount == 0 {
		fmt.Println("No result found")
	}
}
