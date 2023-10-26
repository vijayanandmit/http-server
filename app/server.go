package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	//is this working?
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	buff := make([]byte, 1024)
	_, err = conn.Read(buff)
	if err != nil {
		fmt.Printf("Read error: %s\n", err)
	}

	request := strings.Split(string(buff), "\r\n")
	startLine := request[0]

	if strings.Fields(startLine)[1] == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if strings.HasPrefix(startLine, "/echo/") {
		// Extract the echoed content
		echoedContent := startLine[len("/echo/"):]

		// Respond with a 200 OK with headers and echoed content
		response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: " + fmt.Sprint(len(echoedContent)) + "\r\n\r\n" + echoedContent
		conn.Write([]byte(response))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

}
