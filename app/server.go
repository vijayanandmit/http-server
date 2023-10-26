package main

import (
	"bufio"
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	con, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer con.Close()

	//      buffer := make([]byte, 1024)
	//		n, err := con.Read(buffer)

	response := "HTTP/1.1 200 OK\r\n\r\n"

	_, err = con.Write([]byte(response))

	if err != nil {
		fmt.Println("Error sending response:", err)
	}

	reader := bufio.NewReader(con)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	parts := strings.Fields(requestLine)
	if len(parts) < 2 {
		fmt.Println("Invalid request")
		return
	}

	path := parts[0]

	// Check if the path is "/"
	if path == "/\r\n" {
		// Respond with a 200 OK for the root path
		response := "HTTP/1.1 200 OK\r\n\r\n"
		con.Write([]byte(response))
	} else {
		// Respond with a 404 Not Found for other paths
		response := "HTTP/1.1 404 Not Found\r\n\r\n"
		con.Write([]byte(response))
	}

}
