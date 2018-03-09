package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	conn,_ := net.Dial("tcp","baidu.com:80")
	fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
	status,_ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

	for i:=0;i<100;i++{
		fmt.Printf("%d ",i)
	}
}
