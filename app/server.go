package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		messageSize := []byte{0, 0, 0, 0} // anything
		headers := []byte{0, 0, 0, 7}     // correlation id
		conn.Write(messageSize)
		conn.Write(headers)
	}

}
