package main

import (
	"time"
	"fmt"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(time.Millisecond * 600)

	go bSend(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			fmt.Println("time out")
			return
		}
	}
}

func bSend(msg chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			close(msg)
			fmt.Println("close msg")
			return
		default:
			msg <- "hello"
			time.Sleep(time.Millisecond * 100)
		}
	}
}
