package main

import (
	"log"
	"os"
	"path/filepath"
)

func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func filterPaths(paths []string) []string {
	var result []string

	for _, path := range paths {
		if ok, _ := IsDirectory(path); ok {
			result = append(result, path)
		} else {
			log.Printf("Invalid path: '%v'\n", path)
		}
	}

	return result
}

func Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			name := info.Name()
			// skip folders that begin with a dot
			hidden := filepath.HasPrefix(name, ".") && name != "." && name != ".."
			if hidden {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}
