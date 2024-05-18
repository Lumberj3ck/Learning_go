package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintFilesWithFilterClosure(".", containsDot)
}

func containsDot(filename string) bool {
	return strings.Contains(filename, "eiler")
}

func PrintFilesWithFilterClosure(path string, filter func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)

		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())

			if filter(filename) {
				fmt.Println(filename)
			}

			if f.IsDir() {
				walk(filename)
			}
		}
	}

	walk(path)
}

func PrintAllFiles(path, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())

		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}
		if f.IsDir() {
			PrintAllFiles(filename, filter)
		}
	}
}
