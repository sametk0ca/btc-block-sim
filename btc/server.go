package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	HOST         = "localhost"
	PORT         = "8080"
	TYPE         = "tcp"
	passwordFile = "blocksize.txt"
)

func main() {
	log.Println("Starting TCP Server!")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	handleErrors(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		handleErrors(err)
		go handleRequests(conn)
	}
}

func handleRequests(conn net.Conn) {
	defer conn.Close()

	log.Println("Download starting...")

	// Open and read the file
	file, err := os.Open(passwordFile)
	handleErrors(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			byteSize, err := strconv.Atoi(parts[2])
			handleErrors(err)

			// Create a byte array of the specified size
			data := make([]byte, byteSize)
			// Send the byte array to the client
			_, err = conn.Write(data)
			handleErrors(err)
		}
	}

	log.Println("All files transferred. Press Ctrl+C to exit.")
}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
