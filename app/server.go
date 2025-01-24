package main

import (
	"fmt"
	"net"
	"os"
)

type (
	KafkaServer struct {
		listener net.Listener
	}
)

func (ks *KafkaServer) Listen() {
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		defer conn.Close()
		req := ks.parseRequest(conn)
		response := ks.process(&req)
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		ks.writeResponse(conn, response)
	}
}

func (ks *KafkaServer) process(req *KafkaRequest) KafkaResponse {
	if !req.Supported() {
		return NewKafkaResponse(req, 35)
	}
	return NewKafkaResponse(req, 0)
}

func (ks *KafkaServer) writeResponse(conn net.Conn, response KafkaResponse) {
	conn.Write(response.Bytes())
}

func (ks *KafkaServer) parseRequest(conn net.Conn) KafkaRequest {
	request := make([]byte, 12)
	conn.Read(request)
	return NewKafkaRequest(request)
}

func main() {
	ks := KafkaServer{}
	ks.Listen()
}
