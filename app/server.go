package main

import (
	"bufio"
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
)

func main() {
	// Create a TCP listener
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the HTTP request from the connection
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	request := strings.Split(string(requestLine), "\r\n")
	startLine := request[0]
	if strings.Fields(startLine)[1] == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

}
