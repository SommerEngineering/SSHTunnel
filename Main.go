package main

import (
	"golang.org/x/crypto/ssh"
	"runtime"
)

func main() {

	// Allow Go to use all CPUs:
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Read the configuration from the command-line args:
	readFlags()

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
