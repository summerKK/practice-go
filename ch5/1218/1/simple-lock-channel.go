package main

import (
	"time"
	"fmt"
)

func main() {
	lock := make(chan bool, 1)
	for i := 0; i < 6; i++ {
		go work(i, lock)
	}
	time.Sleep(time.Second * 10)
}
func work(i int, lock chan bool) {
	fmt.Printf("%d want the lock\n", i)
	lock <- true
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("%d has the lock\n", i)
	<-lock
}
