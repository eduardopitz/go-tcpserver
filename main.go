package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	println("...")

	netLiten, netListenError := net.Listen("tcp", ":8080")
	if (netListenError != nil) {
		log.Panic(netListenError)
	}
	defer netLiten.Close()

	for {
		connection, connectionError := netLiten.Accept()
		if (connectionError != nil) {
			log.Panic(connectionError)
		}

		io.WriteString(connection, "\nOlá! Eu sou o seu servidor TCP.")
		fmt.Fprintln(connection, "Como vai você?")

		connection.Close()
	}
}