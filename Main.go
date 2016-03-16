package main

import (
	"fmt"
	"github.com/SommerEngineering/SSHTunnel/Tunnel"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"runtime"
)

func main() {

	// Show the current version:
	log.Println(`SSHTunnel v1.3.0`)

	// Allow Go to use all CPUs:
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Read the configuration from the command-line args:
	readFlags()

	// Check if the password was provided:
	for true {
		if password == `` {
			// Promt for the password:
			fmt.Println(`Please provide the password for the connection:`)
			if pass, errPass := gopass.GetPasswd(); errPass != nil {
				log.Println(`There was an error reading the password securely: ` + errPass.Error())
				os.Exit(1)
				return
			} else {
				password = string(pass)
			}
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
