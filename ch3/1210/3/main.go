package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		defer func() {
			fmt.Println("Hello world")
		}()
		res := <-c
		fmt.Println()
		fmt.Println(res)
	}()

	c <- 1

	time.Sleep(time.Second*10)

}
