package main

import (
	"encoding/json"
	"log"

	"github.com/coffemanfp/backup"
	"github.com/matryer/filedb"
)

var (
	interval int
	archive  string
	dbPath   string
)

var col *filedb.C

func main() {
	m := &backup.Monitor{
		Destination: archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	var path path

	col.ForEach(func(_ int, data []byte) (stop bool) {
		if err := json.Unmarshal(data, &path); err != nil {
			log.Fatalln(err)
			stop = true
			return
		}
		m.Paths[path.Path] = path.Hash
		return
	})

	if len(m.Paths) < 1 {
		log.Fatalln("no paths - use backup tool to add at least one")
	}
}
