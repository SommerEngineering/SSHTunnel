package Tunnel

const (
	maxRetriesLocal  = 16 // How many retries are allowed to create the local end-point?
	maxRetriesRemote = 16 // How many retries are allowed to create the remote end-point?
	maxRetriesServer = 16 // How many retries are allowed to create the SSH server's connection?
)
