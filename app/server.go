package main

import (
	"fmt"
	"net/http"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func extractPath(data string) string {
	parts := strings.Split(data, " ")
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

func main() {
	var httpVersion = "HTTP/1.1"
	1
	var statusAccepted int
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		data := make([]byte, 1024)
		n, err := con.Read(data)
		if err != nil {
			fmt.Println("Error reading data from connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Read %d bytes", n)
		path := extractPath(string(data))
		if path == "/" {
			statusAccepted = http.StatusOK
		} else {
			statusAccepted = http.StatusNotFound
		}
		_, err = con.Write([]byte(fmt.Sprintf("%s %d %s\r\n\r\n", httpVersion, statusAccepted, http.StatusText(statusAccepted))))
		if err != nil {
			fmt.Println("Error reading data from connection: ", err.Error())
			os.Exit(1)
		}
		con.Close()
	}
}
