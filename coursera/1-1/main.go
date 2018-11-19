package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	prefix := []string{}
	printPath(out, path, prefix, printFiles)
	return nil
}

func printPath(out io.Writer, path string, prefix []string, printFiles bool) error {
	files := readDirectory(path, printFiles)

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for index, file := range files {
		isLastElement := index+1 == len(files)
		printPrefix(out, prefix, isLastElement)
		if file.IsDir() {
			prefix = addPrefix(prefix, isLastElement)
			printFolder(out, file.Name())
			path := filepath.Join(path, file.Name())
			errPath := printPath(out, path, prefix, printFiles)
			if errPath != nil {
				panic(errPath.Error())
			}
			prefix = remotePrefix(prefix)
		} else if printFiles {
			printFile(out, file)
		}
	}
	return nil
}

func printPrefix(out io.Writer, prefix []string, isLastElement bool) {
	filePrefix := ""
	if isLastElement {
		filePrefix = "└───"
	} else {
		filePrefix = "├───"
	}

	fmt.Fprintf(out, "%v%v", strings.Join(prefix, ""), filePrefix)
}

func printFolder(out io.Writer, folderName string) {
	fmt.Fprintf(out, "%v\n", folderName)
}

func printFile(out io.Writer, file os.FileInfo) {
	sizeName := "empty"
	if file.Size() > 0 {
		sizeName = strconv.FormatInt(file.Size(), 10) + "b"
	}
	fmt.Fprintf(out, "%v (%v)\n", file.Name(), sizeName)
}

func addPrefix(prefix []string, isLastElement bool) []string {
	if isLastElement {
		return append(prefix, "\t")
	}
	return append(prefix, "│\t")
}

func remotePrefix(prefix []string) []string {
	if len(prefix) > 0 {
		return prefix[:len(prefix)-1]
	}
	return prefix
}

func readDirectory(path string, printFiles bool) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}

	if !printFiles {
		folders := files[:0]
		for _, file := range files {
			if file.IsDir() {
				folders = append(folders, file)
			}
		}
		return folders
	}
	return files
}
