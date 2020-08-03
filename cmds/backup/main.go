package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/matryer/filedb"
)

var dbPath string
var col filedb.C

func main() {
	switch strings.ToLower(os.Args[0]) {
	case "list":
		paths, err := getPaths(col)
		if err != nil {
			log.Fatalln(err)
		}

		for _, path := range paths {
			fmt.Sprintln("= %s\n", path.String())
		}
	case "add":
	case "remove":
	}
}
