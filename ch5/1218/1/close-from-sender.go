package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan bool)
	timeout := time.After(time.Millisecond * 400)
	go cSend(ch)
	for {
		select {
		case <-ch:
			fmt.Println("Got message.")
		case <-timeout:
			fmt.Println("time out!")
			return
		default:
			time.Sleep(time.Millisecond * 100)
			fmt.Println("*yawn*")
		}
	}
}
func cSend(ch chan<- bool) {
	time.Sleep(time.Millisecond*120)
	ch <- true
	close(ch)
	fmt.Println("sent and close")
}
