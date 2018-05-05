package server

import (
	"io"
	"net"

	"../protocol"
)

func handleConnection(conn net.Conn) protocol.Protocol {
	buffer := make([]byte, 0, bufferCapacity)

	_, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		handleError(err)
	}

	decodedProtocol := protocol.Protocol{}
	protocol.Decode(buffer, &decodedProtocol)

	return decodedProtocol
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
			protocol := handleConnection(conn)
		}
	}
}
