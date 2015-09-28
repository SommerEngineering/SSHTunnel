package main

import (
	"fmt"
	"github.com/SommerEngineering/SSHTunnel/Tunnel"
	"golang.org/x/crypto/ssh"
	"runtime"
)

func main() {

	// Show the current version:
	fmt.Println(`SSHTunnel v1.2.0`)

	// Allow Go to use all CPUs:
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Read the configuration from the command-line args:
	readFlags()

	// Check if the password was provided:
	for true {
		if password == `` {
			// Promt for the password:
			fmt.Println(`Please provide the password for the connection:`)
			fmt.Scanln(&password)
		} else {
			break
		}
	}

	// Create the SSH configuration:
	Tunnel.SetPassword4Callback(password)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.PasswordCallback(Tunnel.PasswordCallback),
			ssh.KeyboardInteractive(Tunnel.KeyboardInteractiveChallenge),
		},
	}

	// Create the local end-point:
	localListener := Tunnel.CreateLocalEndPoint(localAddrString)

	// Accept client connections (will block forever):
	Tunnel.AcceptClients(localListener, config, serverAddrString, remoteAddrString)
}
