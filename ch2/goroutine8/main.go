package main

import (
	"time"
	"fmt"
)

func main() {
	rate := 5
	ticker := time.NewTicker(time.Second * 10 / time.Duration(rate))
	limiter := make(chan time.Time, 50)
	go func() {
		for t := range ticker.C {
			select {
			case limiter <- t:
			default:
			}
		}
	}()
	for i := 0; i < 6; i++ {
		go test(limiter, i)
	}
	time.Sleep(time.Hour)
}

func test(ch chan time.Time, i int) {
	for t := range ch {
		fmt.Println(t, i)
	}
	fmt.Println(i, "done")
}
