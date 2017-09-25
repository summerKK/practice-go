package main

import (
	"runtime"
	"fmt"
)

//gosched
func main() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()

		for i := 0; i < 4; i++ {
			if (i == 1) {
				//让出调度器执行b
				runtime.Gosched()
			}
			fmt.Printf("a:%d\n", i)
		}
	}()

	<-exit
}

/*
a:0
b
a:1
a:2
a:3
 */
