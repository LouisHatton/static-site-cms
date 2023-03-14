package files

import "os"

func WalkDirectory(root string) (folders []string, files []string) {
	dir, _ := os.ReadDir(root)

	for _, elem := range dir {
		name := elem.Name()
		if elem.IsDir() {
			folders = append(folders, name)
		} else {
			files = append(files, name)
		}
	}

	return folders, files
}
