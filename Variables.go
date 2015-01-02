package main

var (
	username            = `` // The SSH user's name
	password            = `` // The user's password
	serverAddrString    = `` // The SSH server address
	localAddrString     = `` // The local end-point
	remoteAddrString    = `` // The remote end-point (on the SSH server's side)
	currentRetriesLocal = 0  // Check how many retries are occur for creating the local end-point
)
