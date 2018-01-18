package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var chan1 = make(chan int, 100)
	var sign = make(chan byte, 2)

	go func() {
		for i := 0; i < 100; i++ {
			chan1 <- i
		}
		sign <- 0
		close(chan1)
	}()

	go func() {
	L:
		for {
			select {
			case e, ok := <-chan1:
				if !ok {
					fmt.Println("End.")
					break L
				}
				fmt.Println(e)
			}
		}
		sign <- 1
	}()

	<-sign
	<-sign

	var wg sync.WaitGroup
	wg.Add(2)
	var chan2 = make(chan int, 100)

	go func() {
		time.Sleep(time.Second)
		for i := 100; i < 200; i++ {
			chan2 <- i
		}
		close(chan2)
		wg.Done()
	}()

	go func() {
	N:
		for {
			select {
			case e, ok := <-chan2:
				if !ok {
					break N
				}
				fmt.Println(e)
			case <-func() chan bool {
				timeout := make(chan bool, 1)
				go func() {
					time.Sleep(time.Millisecond)
					timeout <- false
				}()
				return timeout
			}():
				fmt.Println("timeout")
				break N
			}
		}
		wg.Done()
	}()

	wg.Wait()

	unbufChan := make(chan int)
	go func() {
		time.Sleep(time.Second)
		num := <-unbufChan
		fmt.Println(num)
	}()

	num := 1
	unbufChan <- num
	fmt.Println("done")

	t := time.NewTimer(time.Second * 2)
	now := time.Now()
	fmt.Printf("Now time:%v.\n", now)
	expire := <-t.C
	fmt.Printf("expiration time:%v.\n", expire)

	wg.Add(1)

	go func() {
		var timer *time.Timer
	M:
		for {
			select {
			case <-func() <-chan time.Time {
				if timer == nil {
					timer = time.NewTimer(time.Millisecond)
				} else {
					timer.Reset(time.Millisecond)
				}
				return timer.C
			}():
				fmt.Println("time out")
				break M
			}
		}
		wg.Done()
	}()

	wg.Wait()

	var t1 *time.Timer
	f := func() {
		fmt.Printf("expiration time:%v.\n", time.Now())
		fmt.Printf("c's len:%d\n", len(t1.C))
	}
	t1 = time.AfterFunc(1*time.Second, f)

	time.Sleep(time.Second * 2)
}
