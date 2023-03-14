package files

import (
	"fmt"
	"os"
	"strings"
)

type DirectoryMap struct {
	Root   string  `json:"root"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Path  string   `json:"path"`
	Files []string `json:"files"`
}

// Takes a path to a directory and returns a list of folders and files in there.
// The folders are absolute paths based on the root provided.
func WalkDirectory(root string) (folders []string, files []string, err error) {
	dir, err := os.ReadDir(root)
	if err != nil {
		return folders, files, fmt.Errorf("failed to read directory %s: "+err.Error(), root)
	}

	for _, elem := range dir {
		name := elem.Name()
		if elem.IsDir() {
			folders = append(folders, root+name)
		} else {
			files = append(files, name)
		}
	}

	return folders, files, nil
}

func GenerateDirectoryMap(root string) (DirectoryMap, error) {

	// For consistency we will remove any trailing slashes
	if root[len(root)-1:] == "/" {
		root = root[:len(root)-1]
	}

	// Parse the root directory first and generate the map
	requiredFolders, files, err := WalkDirectory(root + "/")
	if err != nil {
		return DirectoryMap{}, fmt.Errorf("failed to parse root directory: " + err.Error())
	}

	dirMap := DirectoryMap{
		Root:   root,
		Routes: []Route{{Path: "/", Files: files}},
	}

	for {
		if len(requiredFolders) == 0 {
			break
		}

		currentFolder := requiredFolders[0]
		requiredFolders = requiredFolders[1:]

		folders, files, err := WalkDirectory(currentFolder + "/")
		if err != nil {
			continue
		}

		requiredFolders = append(requiredFolders, folders...)

		path := strings.Split(currentFolder, root)[1]
		dirMap.Routes = append(dirMap.Routes, Route{Path: path, Files: files})
	}

	return dirMap, nil
}
