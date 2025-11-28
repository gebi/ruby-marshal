package main

import (
	"log"
	"os"

	. "github.com/gebi/ruby-marshal"
)

func main() {
	var input *os.File
	var err error

	if len(os.Args) > 1 {
		input, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	err = DumpSchema(input, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
