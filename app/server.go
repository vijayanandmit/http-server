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

	// Extract the path from the request
	parts := strings.Fields(requestLine)
	if len(parts) < 2 {
		fmt.Println("Invalid request")
		return
	}

	path := parts[0]

	// Check if the path is "/"
	if strings.Fields(path)[1] == "/" {
		// Respond with a 200 OK for the root path
		response := "HTTP/1.1 200 OK\r\n\r\n"
		conn.Write([]byte(response))
	} else {
		// Respond with a 404 Not Found for other paths
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		conn.Write([]byte(response))
	}

}
