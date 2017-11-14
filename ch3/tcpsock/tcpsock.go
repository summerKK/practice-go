package main

import (
	"net"
	"fmt"
	"time"
	"io"
	"bytes"
	"strconv"
	"math"
	"errors"
	"sync"
	"math/rand"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8080"
	DELIMITER      = '\t'
)

var logSn = 1
var wg sync.WaitGroup

func printLog(format string, args ...interface{}) {
	fmt.Printf("%d:%s\n", logSn, fmt.Sprintf(format, args...))
	logSn++
}

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printLog("Listen error:%s", err)
		return
	}
	defer listener.Close()
	printLog("Got listener for the server.(local address:%s)", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			printLog("Accept error:", err)
		}
		printLog("Establish a connection with a clicent application.(remote address:%s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printLog("The connection is closed by another side. (server)")
			} else {
				printLog("Read error:%s (server)", err)
			}
			break
		}
		printLog("Received request:%s (server)", strReq)
		i32Req, err := convertToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			if err != nil {
				printLog("Write Error (written %d bytes): %s (Server)\n", err)
			}
			printLog("Sent response (written %d bytes): %s (Server)\n", n, err)
			continue
		}
		f64Resp := cbrt(i32Req)
		respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
		n, err := write(conn, respMsg)
		if err != nil {
			printLog("write error:%s.(server)", err)
		}
		printLog("Sent response (written %d bytes):%s (server)", n, respMsg)
	}
}

func clientGo(id int) {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printLog("Dial Error:%s (client[%d])", err, id)
		return
	}
	defer conn.Close()
	printLog("connected to server.(remote address:%s,local address:%s) (client[%d])", conn.RemoteAddr(), conn.LocalAddr(), id)
	time.Sleep(200 * time.Millisecond)
	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i < requestNumber; i++ {
		i32Req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", i32Req))
		if err != nil {
			printLog("write error:%s (client[%d])", err, id)
			continue
		}
		printLog("sent request (written %d bytes):%d (client[%d])", n, i32Req, id)
	}

	for j := 0; j < requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printLog("The connection is closed by another side. (Client[%d])\n", id)
			} else {
				printLog("Read Error: %s (Client[%d])\n", err, id)
			}
			break
		}
		printLog("Received response: %s (Client[%d])\n", strResp, id)
	}

}

func cbrt(param int32) float64 {
	return math.Cbrt(float64(param))
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func convertToInt32(str string) (int32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		printLog("Parse error:%s", err)
		return 0, err
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		printLog("convert error the inteager %s is to large/small.", num)
		return 0, errors.New(fmt.Sprintf("%s is not 32-bit inteager", num))
	}
	return int32(num), nil

}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}
