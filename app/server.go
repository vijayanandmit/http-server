package main

import (
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

	response := "HTTP/1.1 200 OK\r\n\r\n"

	_, err = con.Write([]byte(response))

	if err != nil {
		fmt.Println("Error sending response:", err)
	}

	buff := make([]byte, 1024)
	_, err = con.Read(buff)
	if err != nil {
		fmt.Printf("Read error: %s\n", err)
	}

	request := strings.Split(string(buff), "\r\n")
	startLine := request[0]
	if strings.Fields(startLine)[1] == "/" {
		con.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		con.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
