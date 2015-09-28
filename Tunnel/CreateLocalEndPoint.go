package Tunnel

import (
	"log"
	"net"
	"time"
)

func CreateLocalEndPoint(localAddrString string) (localListener net.Listener) {

	// Loop for the necessary retries
	for {

		// Try to create the local end-point
		if localListenerObj, err := net.Listen(`tcp`, localAddrString); err != nil {

			// It was not able to create the end-point:
			currentRetriesLocal++
			log.Printf("Was not able to create the local end-point %s: %s\n", localAddrString, err.Error())

			// Is another retry possible?
			if currentRetriesLocal < maxRetriesLocal {
				log.Println(`Retry...`)
				time.Sleep(1 * time.Second)
			} else {
				log.Fatalln(`No more retries for the local end-point: ` + localAddrString) // => Exit
			}
		} else {
			// Success!
			log.Println(`Listen to local address ` + localAddrString)
			localListener = localListenerObj
			return
		}
	}
}
