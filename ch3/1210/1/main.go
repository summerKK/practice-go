package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	c chan int = make(chan int)
)

func main() {

	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(10)
	fmt.Printf("%v\n",&wg)

	go work(wg)

	go func() {
		for {
			res := <-c
			wg.Done()
			fmt.Println(res)
		}
	}()

	time.Sleep(time.Second * 2)

	wg.Wait()

}

func work(wg sync.WaitGroup) {
	fmt.Printf("%v\n",wg)
	for i := 0; i < 10; i++ {
		c <- i
		wg.Add(1)
		time.Sleep(time.Second)
	}
}
