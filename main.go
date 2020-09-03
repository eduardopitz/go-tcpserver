package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	println("...")

	netListen, netListenError := net.Listen("tcp", ":8080")
	if (netListenError != nil) {
		log.Panic(netListenError)
	}
	defer netListen.Close()

	for {
		connection, connectionError := netListen.Accept()
		if (connectionError != nil) {
			log.Println(connectionError)
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer connection.Close()
}