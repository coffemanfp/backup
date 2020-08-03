package main

import (
	"encoding/json"

	"github.com/matryer/filedb"
)

func getPaths(col *filedb.C) (paths []path, err error) {
	var path path

	col.ForEach(func(i int, data []byte) (stop bool) {
		err = json.Unmarshal(data, &path)
		if err != nil {
			return
		}

		paths = append(paths, path)
		return
	})

	return
}

func addPath(col *filedb.C, path path) (err error) {
	if path.Hash == "" {
		path.Hash = "Not yet archived"
	}

	err = col.InsertJSON(path)
	return
}

func removePaths(col *filedb.C, pathsToDelete []string) (pathsDeleted []string, err error) {
	var path path

	col.RemoveEach(func(i int, data []byte) (shouldDelete bool, stop bool) {
		err = json.Unmarshal(data, &path)
		if err != nil {
			stop = true
			return
		}

		for _, pathToDelete := range pathsToDelete {
			if pathToDelete == path.Path {
				shouldDelete = true
				pathsDeleted = append(pathsDeleted, pathToDelete)
				return
			}
		}

		return
	})

	return
}
