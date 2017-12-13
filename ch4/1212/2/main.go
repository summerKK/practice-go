package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch chan int
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println(<-ch)
	}()

	test1()
	//go test2(&wg)

	//wg.Wait()

}

func test1() {
	time.Sleep(time.Second)
	ch = make(chan int)
	ch <- 1
}

func test2(wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}
