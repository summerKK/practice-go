package main

import (
	"fmt"
	"time"
)

var x int

func main() {
	v := 100
	go func(v, x int) {
		fmt.Println("go:", v, x)
	}(v, counter())

	v += 100
	fmt.Println("main:", v, counter())

	time.Sleep(1 * time.Second)
}

func counter() int {
	x++
	return x
}
