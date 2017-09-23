package main

import (
	"sync"
	"fmt"
	"time"
)

//可以在多处添加wait阻塞
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Wait()
		//time.Sleep(time.Second)
		fmt.Println("go1:done")
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println("go2:done")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main")
}

/*
go2:done
go1:done
main
 */
