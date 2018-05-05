package server

import (
	"net"
)

func handleConnection(conn *net.Conn) {

}

func Start() {
	tcpAddr, err := net.ResolveTCPAddr(network, formatAddress(port))
	handleError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	handleError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logError(err.Error())
		} else {
			logMessage("Connection established")
			handleConnection(&conn)
		}
	}
}
