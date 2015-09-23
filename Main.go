package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"runtime"
)

func main() {

	// Show the current version:
	fmt.Println(`SSHTunnel v1.1.0`)

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
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.PasswordCallback(passwordCallback),
			ssh.KeyboardInteractive(keyboardInteractiveChallenge),
		},
	}

	// Create the local end-point:
	localListener := createLocalEndPoint()

	// Accept client connections (will block forever):
	acceptClients(localListener, config)
}
