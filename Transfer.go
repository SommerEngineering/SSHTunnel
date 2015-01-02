package main

import (
	"io"
	"log"
)

// The transfer function.
func transfer(fromReader io.Reader, toWriter io.Writer, name string) {

	log.Printf("%s transfer started.", name)

	// This call blocks until the client or service will close the connection.
	// Therefore, this call maybe takes hours or even longer. Concern, may this
	// program will be used to connect multiple servers to make e.g. a database
	// available...
	if _, err := io.Copy(toWriter, fromReader); err != nil {

		// In this case, we do not fail the whole program: Regarding how the client
		// or the service was e.g. shut down, the error may only means 'client has been closed'.
		log.Printf("%s transfer failed: %s\n", name, err.Error())
	} else {
		log.Printf("%s transfer closed.\n", name)
	}
}
