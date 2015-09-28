package Tunnel

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func AcceptClients(connection net.Listener, config *ssh.ClientConfig, serverAddrString, remoteAddrString string) {

	// Endless loop
	for {

		// Accept (another) client connection:
		if localConn, err := connection.Accept(); err != nil {

			// Fail
			log.Printf("Accepting a client failed: %s\n", err.Error())
		} else {

			// Success
			log.Println(`Client accepted.`)

			// Start the forwarding:
			go forward(localConn, config, serverAddrString, remoteAddrString)
		}
	}
}
