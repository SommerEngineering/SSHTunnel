package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func forward(localConn net.Conn, config *ssh.ClientConfig) {

	sshClientConn, err := ssh.Dial("tcp", serverAddrString, config)
	if err != nil {
		log.Printf("ssh.Dial failed: %s\n", err)
		return
	}

	if sshConn, err := sshClientConn.Dial("tcp", remoteAddrString); err != nil {
		log.Println(`Was not able to create the tunnel: ` + err.Error())
	} else {
		go transfer(localConn, sshConn)
		go transfer(sshConn, localConn)
	}
}
