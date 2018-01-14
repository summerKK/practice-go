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
	"math/rand"
	"sync"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = ":8085"
	DELIMITER      = '\t'
)

var (
	logSn int
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(time.Second)
	go clientGo(1)
	wg.Wait()
}

func serverGo() {
	defer wg.Done()
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printLog("listen error:%s\n", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			printLog("accept error:%s\n", err)
		}
		printLog("established a connection with a client application.(remote address:%s)\n", conn.RemoteAddr())
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer func() {
		wg.Done()
		conn.Close()
	}()
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printLog("the connecion is closed by another side.(server)\n")
			} else {
				printLog("read error:%s (server)\n", err)
			}
			break
		}
		printLog("received request:%s (server)\n", strReq)

		i32Req, err := convert2Int32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			if err != nil {
				printLog("write error (written %d bytes): %s (server)\n", n, err)
			}
			printLog("sent response (written %d bytes):%s (server)\n", n, err)
			continue
		}
		f64Resp := cbrt(i32Req)
		respMsg := fmt.Sprintf("the cube root of %d is %f.", i32Req, f64Resp)
		n, err := write(conn, respMsg)
		if err != nil {
			printLog("write error:%s (server)\n", err)
		}
		printLog("sent response (written %d bytes):%s (server)\n", n, respMsg)
	}
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buf bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buf.WriteByte(readByte)
	}
	return buf.String(), nil
}

func write(conn net.Conn, s string) (int, error) {
	var buf bytes.Buffer
	buf.WriteString(s)
	buf.WriteByte(DELIMITER)
	return conn.Write(buf.Bytes())
}

func printLog(format string, args ...interface{}) {
	fmt.Printf("%d:%s", logSn, fmt.Sprintf(format, args...))
	logSn++
}

func convert2Int32(s string) (int32, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		printLog("parse error:%s\n", err)
		return 0, err
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		printLog("convert error:the integer %d is too large/small.\n", num)
		return 0, errors.New(fmt.Sprintf("convert error:the integer %s is too large/small.\n", num))
	}
	return int32(num), nil
}

func cbrt(i int32) float64 {
	return math.Cbrt(float64(i))
}

func clientGo(id int) {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printLog("dial error:%s (client[%d])\n", err, id)
		return
	}
	defer conn.Close()
	printLog("connected to server (remote address:%s,local address:%s) (client[%d])\n", conn.RemoteAddr(), conn.LocalAddr(), id)
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < requestNumber; i++ {
		i32Req := r.Int31()
		n, err := write(conn, fmt.Sprintf("%d", i32Req))
		if err != nil {
			printLog("write error:%s (client[%d]) \n", err, id)
			continue
		}
		printLog("sent request (written %d bytes):%d (client[%d]) \n", n, i32Req, id)
	}

	for j := 0; j < requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printLog("the connection is closed by another side. (client[%d])\n", id)
			} else {
				printLog("read error:%s (client[%d])\n", err, id)
			}
			break
		}
		printLog("received response:%s (client[%d])\n", strResp, id)
	}
}
