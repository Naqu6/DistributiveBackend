package main

import (
	"io/ioutil"
	"strings"
)

type DirectoryStructure struct {
	name     string
	path     string
	url      string
	children []DirectoryStructure
}

func directoryContents(directoryName string) (results []string) {
	files, err := ioutil.ReadDir(directoryName)

	if err != nil {
		return results
	}

	for _, file := range files {
		if file.IsDir() {
			filePath := directoryName + "/" + file.Name()

			results = append(results, filePath)
			results = append(results, directoryContents(filePath)...)
		}
	}

	return results
}

func directoryContentsHierarchy(directoryName string) (results []DirectoryStructure) {
	files, err := ioutil.ReadDir(directoryName)

	if err != nil {
		return results
	}

	for _, file := range files {
		if file.IsDir() {
			filePath := directoryName + "/" + file.Name() + "/"
			results = append(results, DirectoryStructure{file.Name(), filePath, strings.Replace(filePath, "pages/", "", 0), directoryContentsHierarchy(filePath)})
		}
	}

	return results
}
