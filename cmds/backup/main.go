package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/matryer/filedb"
)

var dbPath string
var col *filedb.C

func main() {
	args := flag.Args()

	switch strings.ToLower(args[0]) {
	case "list":
		paths, err := getPaths(col)
		if err != nil {
			log.Fatalln(err)
		}

		for _, path := range paths {
			fmt.Printf("= %s\n", path.String())
		}
	case "add":
		if len(args[1:]) == 0 {
			log.Fatalln("must specify path to add")
		}

		var path path
		var err error

		for _, p := range args[1:] {
			path.Path = p
			path.Hash = "Not yet archived"

			if err = addPath(col, path); err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("+ %s\n", path.String())
		}

	case "remove":
		pathsToDelete := args[1:]

		pathsDeleted, err := removePaths(col, pathsToDelete)
		if err != nil {
			fmt.Println(err)
		}

		for _, pathDeleted := range pathsDeleted {
			fmt.Printf("- %s\n", pathDeleted)
		}
	}
}
