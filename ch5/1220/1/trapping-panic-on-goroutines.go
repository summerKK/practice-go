package main

import (
	"net"
	"fmt"
	"bufio"
	"errors"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Faild to open port on 1026")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			conn.Close()
		}
	}()
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("failed to read from socket. ")
		conn.Close()
	}
	response(data, conn)
}

func response(bytes []byte, conn net.Conn) {
	conn.Write(bytes)
	panic(errors.New("pretend I'am a real error "))
}
