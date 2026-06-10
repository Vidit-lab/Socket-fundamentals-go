package main

import (
	"fmt"
	"io"
	"net"
)

const serverAddress = "127.0.0.1:8080"

func main() {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Failed to start server on %s: %v\n", serverAddress, err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", serverAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	message, err := io.ReadAll(conn)
	if err != nil {
		fmt.Printf("Read error from %s: %v\n", conn.RemoteAddr(), err)
		return
	}

	if len(message) == 0 {
		return
	}

	if _, err := conn.Write(message); err != nil {
		fmt.Printf("Write error to %s: %v\n", conn.RemoteAddr(), err)
		return
	}

	fmt.Printf("Echoed %q to %s\n", string(message), conn.RemoteAddr())
}
