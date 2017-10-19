package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	Practice3()
}

func Practice3() {
	runtime.GOMAXPROCS(1) //use no more than one logic cpu
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
		if i == 9 {
			fmt.Println("loop1 over")
		}
	}
	for i := 10; i < 20; i++ {
		//fmt.Println(i)
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
		if i == 19 {
			fmt.Println("loop2 over")
		}
	}
	wg.Wait()
}
