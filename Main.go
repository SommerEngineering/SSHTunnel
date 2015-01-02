package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func main() {

	readFlags()
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.PasswordCallback(passwordCallback),
			ssh.KeyboardInteractive(keyboardInteractiveChallenge),
		},
	}

	localListener, err := net.Listen(`tcp`, localAddrString)
	if err != nil {
		log.Printf("net.Listen failed: %v\n", err)
	} else {
		log.Println(`Listen to local address.`)
	}

	for {
		localConn, err := localListener.Accept()
		if err != nil {
			log.Printf("listen.Accept failed: %v\n", err)
		} else {
			log.Println(`Accepted a client.`)
			go forward(localConn, config)
		}
	}
}
