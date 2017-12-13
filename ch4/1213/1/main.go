package main

import (
	"practice/ch4/1213/1/sub"
	"time"
)

func main() {

	go func() {
		i := 0
		for {
			sub.Receiver <- i
			i++
			time.Sleep(time.Second)
		}

	}()

	sub.Handle()
}
