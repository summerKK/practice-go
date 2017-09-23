package main

import (
	"math"
	"sync"
	"runtime"
	"fmt"
	"time"
)

func counter() {
	x := 0
	for i := 0; i < math.MaxInt32; i++ {
		x += i
	}
	fmt.Println(x)
}

func test1(n int) {
	var now = time.Now()
	for i := 0; i < n; i++ {
		counter()
	}
	fmt.Printf("%f\n", time.Since(now).Seconds())
}

func test2(n int) {
	var now = time.Now()
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			counter()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%f\n", time.Since(now).Seconds())
}

func main() {
	//返回核心数
	n := runtime.GOMAXPROCS(0)
	fmt.Println(n)
	test1(n)
	test2(n)
}

/*
test1
4
2305843005992468481
2305843005992468481
2305843005992468481
2305843005992468481
3.043131
 */

/*
test2
4
2305843005992468481
2305843005992468481
2305843005992468481
2305843005992468481
1.595129
 */
