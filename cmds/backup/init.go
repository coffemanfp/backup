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
	flag.StringVar(&dbPath, "db", "./backupdata", "path to database directory")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		log.Fatalln("invalid usage; must specify command")
	}
}

func initDatabase() {
	db, err := filedb.Dial(*dbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	col, err = db.C("paths")
	if err != nil {
		log.Fatalln(err)
	}
}
