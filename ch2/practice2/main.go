package main

import "fmt"

//阻塞操作.
func main() {
	ch := make(chan struct{})
	go func() {
		test()
		//ch <- struct{}{}
		//同样的效果
		close(ch)
	}()
	<-ch
	fmt.Println("111")

}

func test() {
	fmt.Println("Hello world")
}

/*
Hello world
111
 */
