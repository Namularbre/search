package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func parseArgs() string {
	if len(os.Args) < 1 {
		panic("Wrong usage: search your_search")
	}

	return os.Args[1]
}

func isSupportedFile(path string) bool {
	validSuffixes := []string{".txt", ".pdf", ".doc", ".docx", ".png", ".jpeg", ".webp", ".go"}
	isSupported := false

	for _, suffix := range validSuffixes {
		isSupported = strings.HasSuffix(path, suffix)
	}

	return isSupported
}

func search(search string) []string {
	root := "/home/namu"
	var results []string

	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Erreur lors de l'accès au fichier %q : %v\n", path, err)
			return err
		}

		if strings.Contains(path, search) && isSupportedFile(path) {
			fmt.Println(path)
			results = append(results, path+"\n")
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return results
}

func main() {
	searchContent := parseArgs()
	results := search(searchContent)
	if len(results) > 0 {
		fmt.Println(results)
	} else {
		fmt.Println("No results found")
	}
}
