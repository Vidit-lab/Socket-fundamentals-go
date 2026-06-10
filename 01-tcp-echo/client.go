package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const serverAddress = "127.0.0.1:8080"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Message:\n")
	message, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Failed to read input: %v\n", err)
		return
	}

	message = strings.TrimRight(message, "\r\n")

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", serverAddress, err)
		return
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(message)); err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		return
	}

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		if err := tcpConn.CloseWrite(); err != nil {
			fmt.Printf("Failed to finish sending: %v\n", err)
			return
		}
	}

	response, err := io.ReadAll(conn)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return
	}

	fmt.Printf("Received:\n%s\n", strings.TrimRight(string(response), "\r\n"))
}
