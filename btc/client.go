package main

import (
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Println("Unable to resolve the given address\n\tERROR: " + err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Printf("Unable to make a connection to the Server.\n\tERROR: %v", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// Inform the server to start the download
	conn.Write([]byte("Start download"))

	log.Println("Receiving files...")

	// Read the received data from the server
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from server: %v", err)
			break
		}

		log.Printf("Received %d bytes from server", n)
	}

	log.Println("All files received. Press Ctrl+C to exit.")
	select {}
}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
