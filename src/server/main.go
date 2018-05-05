package server

import (
	"fmt"
	"io"
	"net"

	"../protocol"
)

func getRequestFrom(conn net.Conn) protocol.Request {
	buffer := make([]byte, bufferInitialLength, bufferCapacity)

	bytesRead, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		handleError(err)
	}

	decodedRequest := protocol.Request{}
	protocol.DecodeRequest(buffer[:bytesRead], &decodedRequest)

	return decodedRequest
}

func establishConnection(sender string, conn *net.Conn, activeConnections *map[string]*net.Conn) {
	logMessage(fmt.Sprintf("Connection for %v successfully established.", sender))
	(*activeConnections)[sender] = conn
}

func closeConnection(sender string, conn *net.Conn, activeConnections *map[string]*net.Conn) {
	activeConn, ok := (*activeConnections)[sender]
	if ok {
		(*activeConn).Close()
	}
}

func sendMessage(request *protocol.Request, conn *net.Conn, activeConnections *map[string]*net.Conn) {
	activeConn, ok := (*activeConnections)[request.Receiver]
	if ok {
		encodedRequest, err := protocol.EncodeRequest(request)
		if err != nil {
			logError(err.Error())
		} else {
			(*activeConn).Write(encodedRequest)
			logMessage(fmt.Sprintf("Message for %v was sent out.", request.Receiver))
		}
	} else {
		logMessage(fmt.Sprintf("Receiver %v was not found.", request.Receiver))
	}
}

func processRequest(request *protocol.Request, conn *net.Conn, activeConnections *map[string]*net.Conn) {
	if request.Code == protocol.EstablishConnectionCode {
		establishConnection(request.Sender, conn, activeConnections)
	} else if request.Code == protocol.SendMessageCode {
		sendMessage(request, conn, activeConnections)
	} else {
		closeConnection(request.Sender, conn, activeConnections)
	}
}

func Start() {
	formattedAddr := formatAddress(port)

	tcpAddr, err := net.ResolveTCPAddr(network, formattedAddr)
	handleError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	handleError(err)
	logMessage("Listening at localhost" + formattedAddr)

	activeConnections := map[string]*net.Conn{}

	for {
		conn, err := listener.Accept()
		if err != nil {
			logError(err.Error())
		} else {
			request := getRequestFrom(conn)
			processRequest(&request, &conn, &activeConnections)
		}
	}
}
