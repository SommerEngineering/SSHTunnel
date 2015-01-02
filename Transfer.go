package main

import (
	"io"
	"log"
)

func transfer(fromReader io.Reader, toWriter io.Writer) {

	if _, err := io.Copy(toWriter, fromReader); err != nil {
		log.Printf("io.Copy failed: %v\n", err)
	} else {
		log.Println(`Transfer closed.`)
	}
}
