package main

import (
	"encoding/json"
	"log"

	"github.com/coffemanfp/backup"
	"github.com/matryer/filedb"
)

func updateDB(col *filedb.C, m *backup.Monitor) {
	var path path

	col.SelectEach(func(_ int, data []byte) (include bool, newData []byte, stop bool) {
		include = true
		newData = data

		if err := json.Unmarshal(data, &path); err != nil {
			log.Println("faile to unmarshal data (skipping):", err)
			return
		}

		path.Hash, _ = m.Paths[path.Path]

		newData, err := json.Marshal(&path)
		if err != nil {
			log.Println("failed to marshal data (skipping):", err)
		}

		return
	})
}
