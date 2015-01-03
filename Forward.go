package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"time"
)

func forward(localConn net.Conn, config *ssh.ClientConfig) {

	defer localConn.Close()
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
			defer sshClientConnection.Close()
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
			defer sshConn.Close()

			// To be able to close down both transfer threads, we create a channel:
			quit := make(chan bool)

			// Create the transfers to/from both sides (two new threads are created for this):
			go transfer(localConn, sshConn, `Local => Remote`, quit)
			go transfer(sshConn, localConn, `Remote => Local`, quit)

			// Wait and look if any of the two transfer theads are down:
			isRunning := true
			for isRunning {
				select {
				case <-quit:
					log.Println(`At least one transfer was stopped.`)
					isRunning = false
					break
				}
			}

			// Now, close all the channels and therefore, force the other / second thread to go down:
			log.Println(`Close now all connections.`)
			return
		}
	}
}
