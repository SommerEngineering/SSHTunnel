package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

func forward(localConn net.Conn, config *ssh.ClientConfig) {

	currentRetriesServer := 0
	currentRetriesRemote := 0
	var sshClientConnection *ssh.Client = nil

	// Loop for retries:
	for {

		// Try to connect to the SSD server:
		if sshClientConn, err := ssh.Dial(`tcp`, serverAddrString, config); err != nil {

			// Failed:
			currentRetriesServer++
			log.Printf("Was not able to connect with the SSH server %s: %s\n", serverAddrString, err.Error())

			// Is a retry alowed?
			if currentRetriesServer < maxRetriesServer {
				log.Println(`Retry...`)
				time.Sleep(1 * time.Second)
			} else {

				// After the return, this thread is closed down. The client can try it again...
				log.Println(`No more retries for connecting the SSH server.`)
				return
			}

		} else {

			// Success:
			log.Println(`Connected to the SSH server ` + serverAddrString)
			sshClientConnection = sshClientConn
			break
		}
	}

	// Loop for retries:
	for {

		// Try to create the remote end-point:
		if sshConn, err := sshClientConnection.Dial(`tcp`, remoteAddrString); err != nil {

			// Failed:
			currentRetriesRemote++
			log.Printf("Was not able to create the remote end-point %s: %s\n", remoteAddrString, err.Error())

			// Is another retry allowed?
			if currentRetriesRemote < maxRetriesRemote {
				log.Println(`Retry...`)
				time.Sleep(1 * time.Second)
			} else {

				// After the return, this thread is closed down. The client can try it again...
				log.Println(`No more retries for connecting the remote end-point.`)
				return
			}
		} else {

			// Fine, the connections are up and ready :-)
			log.Printf("The remote end-point %s is connected.\n", remoteAddrString)

			// Create the transfers to/from both sides (two new threads are created for this):
			go transfer(localConn, sshConn, `Local => Remote`)
			go transfer(sshConn, localConn, `Remote => Local`)
			return
		}
	}
}
