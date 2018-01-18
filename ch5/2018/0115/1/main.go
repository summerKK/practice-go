package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}

	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello,%s.\n", who)
		}(name)
		runtime.Gosched()
	}

	ch := make(chan int, 5)
	sign := make(chan byte, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		fmt.Println("the channel is closed.")
		sign <- 0
	}()

	go func() {
		for {
			e, ok := <-ch
			fmt.Printf("%d(%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(time.Second * 2)
		}
		fmt.Println("done")
		sign <- 1
	}()

	<-sign
	<-sign

	origs := make(chan int, 100)
	go func() {
		for i := 0; ; i++ {
			origs <- i
			if i == 100 {
				break
			}
		}
		fmt.Println("the channel origs has been closed")
		close(origs)
	}()

	for i := range batch(origs) {
		fmt.Println(i)
	}

	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected")
	default:
		fmt.Println("default")
	}
}

func getNumber(i int) int {
	var numbers = []int{1, 2, 3, 4, 5}
	return numbers[i]
}

func getChan(i int) chan int {
	var ch1 chan int
	var ch2 chan int
	var chs = []chan int{ch1, ch2}
	return chs[i]
}

func batch(origs <-chan int) <-chan int {
	dests := make(chan int, 100)
	go func() {
		for p := range origs {
			dests <- p
		}
		fmt.Println("all the information has been handled")
		fmt.Println("the channel dests has been closed")
		close(dests)
	}()
	return dests
}
