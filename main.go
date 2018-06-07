package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fetchFileInfo()
}

func fetchFileInfo() ([]os.FileInfo, error) {
	searchDir := "./"

	fileInfo := make([]os.FileInfo, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileInfo = append(fileInfo, f)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileInfo {
		fmt.Printf("Name '%s', Size %d", file.Name(), file.Size())
		fmt.Println("")
	}

	return fileInfo, nil
}
