package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := connection.SetDeadline(time.Now().Add(10 * time.Second))
	if (err != nil) {
		log.Println("Timeout")
	}

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer connection.Close()
	fmt.Println("...")
}