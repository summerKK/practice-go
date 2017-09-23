package main

import (
	"time"
	"fmt"
)

func main() {
	//创建无缓冲channel来保证goruntine执行完毕
	exit := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		fmt.Println("goruntine done")
		//通知channel
		close(exit)
		fmt.Println("goruntine donee")
	}()

	fmt.Println("main wait")
	<-exit
	fmt.Println("done")
}
