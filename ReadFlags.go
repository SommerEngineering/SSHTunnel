package main

import (
	"flag"
)

func readFlags() {
	flag.StringVar(&serverAddrString, `server`, `127.0.0.1:22`, `The (remote) SSH server, e.g. 'my.host.com', 'my.host.com:22', '127.0.0.1:22', 'localhost:22'.`)
	flag.StringVar(&localAddrString, `local`, `127.0.0.1:50000`, `The local end-point of the tunnel, e.g. '127.0.0.1:50000', 'localhost:50000'.`)
	flag.StringVar(&remoteAddrString, `remote`, `127.0.0.1:27017`, `The remote side end-point (e.g. on the machine with the SSH server), e.g. a MongoDB (port 27017) '127.0.0.1:27017', a web server '127.0.0.1:80'`)
	flag.StringVar(&username, `user`, `username`, `The user's name for the SSH server.`)
	flag.StringVar(&password, `pwd`, ``, `The user's password for the SSH server.`)
	flag.Parse()
}
