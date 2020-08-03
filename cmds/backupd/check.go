package main

import (
	"fmt"
	"log"

	"github.com/coffemanfp/backup"
	"github.com/matryer/filedb"
)

func check(cool *filedb.C, m *backup.Monitor) (err error) {
	log.Println("Checking...")

	counter, err := m.Now()
	if err != nil {
		err = fmt.Errorf("failed to backup: %s", err)
		return
	}

	if counter > 0 {
		log.Printf("  Archived %d directories\n", counter)

		// update hashes
		updateDB(col, m)

	} else {
		log.Println("  No changes")
	}

	return
}
