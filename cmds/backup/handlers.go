package main

import (
	"encoding/json"

	"github.com/matryer/filedb"
)

func getPaths(col filedb.C) (paths []path, err error) {
	var path path

	col.ForEach(func(i int, data []byte) bool {
		err = json.Unmarshal(data, &path)
		if err != nil {
			return false
		}

		paths = append(paths, path)
		return false
	})

	return
}
