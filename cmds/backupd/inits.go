package main

import (
	"flag"
	"log"

	"github.com/matryer/filedb"
)

func init() {
	initFlags()
	initDatabase()
}

func initFlags() {
	flag.IntVar(&interval, "interval", 10, "interval between checks (seconds)")
	flag.StringVar(&archive, "archive", "archive", "path to archive location")
	flag.StringVar(&dbPath, "db", "./db", "path to filedb database")

	flag.Parse()
}

func initDatabase() {
	db, err := filedb.Dial(dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	col, err = db.C("paths")
	if err != nil {
		log.Fatalln(err)
	}
}
