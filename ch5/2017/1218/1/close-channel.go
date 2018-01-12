package main

import (
	"time"
	"fmt"
)

func main() {
	msg := make(chan string)
	until := time.After(time.Second * 5)

	go send(msg)

	for {
		select {
		case s := <-msg:
			fmt.Println(s)
		case <-until:
			time.Sleep(time.Microsecond * 500)
			close(msg) //panic: send on closed channel
			fmt.Println("done")
			return
		}
	}
}
func send(msg chan<- string) {
	for {
		msg <- "hello world"
		time.Sleep(time.Microsecond * 500)

	}
}
