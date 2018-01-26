package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

func main() {
	var i64 int64
	atomic.AddInt64(&i64, -3)

	fmt.Println(i64)

	var ui64 uint64
	atomic.AddUint64(&ui64, ^uint64(5-1))

	fmt.Println(ui64,uint64(5-1))

	repeatedlyLock()
}

func repeatedlyLock() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock.(GO)")
	mutex.Lock()
	fmt.Println("The lock is locked. (GO)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (GO%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (GO%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (GO)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (GO)")
	time.Sleep(time.Second)
	mutex.Unlock()
	time.Sleep(time.Second)
	mutex.Unlock()
	time.Sleep(time.Second)
}
