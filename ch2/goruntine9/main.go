package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 3, 4, 5, 6, 7, 8, 9}

	c := make(chan int)

	go sum(s[:len(s)/2], c, 1)
	go sum(s[len(s)/2:], c, 2)

	x, y := <-c, <-c

	fmt.Println(x, y)

	calculate()
	timeOut()

	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	ticker1 := time.NewTicker(time.Second * 1)

	go func() {
		for t := range ticker1.C {
			fmt.Println("ticker at", t)
		}
	}()

	//time.Sleep(time.Second * 10)

	c1 := make(chan int, 10)
	c1 <- 1
	c1 <- 2

	close(c1)

	for v := range c1 {
		fmt.Println(v)
	}

	c2 := make(chan int, 2)
	c2 <- 2
	close(c2)
	<-c2

	i, ok := <-c2
	fmt.Printf("%d,%t\n", i, ok)

	c3 := make(chan int, 10)

	go func() {
		fmt.Println("watting")
		c3 <- 4
		c3 <- 5
	}()

	<-c3
	fmt.Println(<-c3)
}

func sum(s []int, c chan int, smbol int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	if smbol == 2 {
		time.Sleep(time.Second)
	}

	c <- sum
}

func calculate() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func timeOut() {
	c := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c <- "Hello World!"
	}()

	select {
	case msg := <-c:
		fmt.Println(msg)
	case <-time.After(time.Second):
		fmt.Println("time out")
	}
}
