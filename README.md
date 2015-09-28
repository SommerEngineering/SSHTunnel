SSHTunnel
=========
SSHTunnel is a tiny small program to tunnel something through a SSH without any external dependencies. Just download the executable which matches your OS and architecture (32 vs. 64 bits) and run it.

### Syntax
*This example uses the Microsoft Windows executable, but the syntax is the same for e.g. Linux, Unix, Mac, etc.*
`SSHTunnel.exe -local 127.0.0.1:53001 -remote 127.0.0.1:27017 -server your-server.org:22 -user john -pwd johndow`

- Connects to your external server `your-server.org` to port `22`. At this port, the SSH service should run
- At the SSH server's side, connects to `127.0.0.1` to port `27017` (a MongoDB database)
- At your local side, provides a listener at `127.0.0.1` at the port `53001`
- The username for the SSH service is `john`
- The user's password would be `johndow` ;-) You can avoid the `-pwd` argument. Thus, the SSHTunnel will ask for the password.
- Now, you are able to use your local MongoDB software and can connect to port `53001` at `localhost`.

### Features
- The whole code is open source and can be used for any purpose (also commercial)
- If you want, you can compile the code by your own by using the [Go](http://www.golang.org)
- The program just needs very low resources e.g. around 1.3 MB memory for Microsoft Windows 8.1
- SSHTunnel is scalable and, if necessary, can utilise all your CPUs
- If a connection cannot setup, the program re-tries it
- At the moment, SSHTunnel uses only the password authentication methods. Therefore, it is currently not possible to use e.g. a certificate, etc. Nevertheless, the implementation of this feature is possible.
- The configuration must be provided by using the command-line arguments. It is currently not possible to use e.g. a configuration file.
- You can avoid the password argument if you prefer to provide the password on demand.
- [Ocean Remote Connections](https://github.com/SommerEngineering/OceanRemoteConnections) is a simple GUI for SSH Tunnel, PuTTY and WinSCP.

### Download
Go and get the latest release from the [release page](https://github.com/SommerEngineering/SSHTunnel/releases).

*Based on damick's example code from http://stackoverflow.com/questions/21417223/simple-ssh-port-forward-in-golang*
