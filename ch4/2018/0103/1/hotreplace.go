package main

import (
	"time"
	"fmt"
)

var (
	chan1       chan int
	chan1Length int           = 18
	interval    time.Duration = time.Millisecond * 1500
)

func main() {

	chan1 = make(chan int, chan1Length)

	go func() {
		for i := 0; i < chan1Length; i++ {
			if i > 0 && i%3 == 0 {
				fmt.Println("reset chan1")
				chan1 = make(chan int, chan1Length)
			}
			fmt.Printf("send element %d...\n", i)
			chan1 <- i
			time.Sleep(interval)
		}
		fmt.Println("close chan1...")
		close(chan1)
	}()

	recevie()
}

func getChan1() chan int {
	return chan1
}

func recevie() {
	fmt.Println("receive element from chan1...")
	timer := time.After(30 * time.Second)
Loop:
	for {
		select {
		case e, ok := <-getChan1():
			if !ok {
				fmt.Println("...closed chan1.")
				break Loop
			}
			fmt.Printf("recevie a element:%d\n", e)
			time.Sleep(time.Millisecond*1600)
		case <-timer:
			fmt.Println("time out")
			break Loop
		}
	}
	fmt.Println("--end")
}
