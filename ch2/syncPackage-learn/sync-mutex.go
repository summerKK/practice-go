package main

import (
	"sync"
	"fmt"
)

type SafeInt struct {
	sync.Mutex
	Num int
}

/*
------------------------------------------------------------

互斥锁

　　互斥锁用来保证在任一时刻，只能有一个例程访问某对象。Mutex 的初始值为解锁状态。Mutex 通常作为其它结构体的匿名字段使用，使该结构体具有 Lock 和 Unlock 方法。

　　Mutex 可以安全的在多个例程中并行使用。

------------------------------
*/

func main() {
	count := SafeInt{}
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			count.Lock()
			count.Num += i
			fmt.Print(count.Num, " ")
			count.Unlock()
			done <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	fmt.Println("done")
}

/*不固定
1 5 7 13 16 25 30 38 45 45 done
 */

 /*
------------------------------------------------------------

读写互斥锁

　　RWMutex 比 Mutex 多了一个“读锁定”和“读解锁”，可以让多个例程同时读取某对象。RWMutex 的初始值为解锁状态。RWMutex 通常作为其它结构体的匿名字段使用。

　　Mutex 可以安全的在多个例程中并行使用。

------------------------------

// Lock 将 rw 设置为写锁定状态，禁止其他例程读取或写入。
func (rw *RWMutex) Lock()

// Unlock 解除 rw 的写锁定状态，如果 rw 未被写锁定，则该操作会引发 panic。
func (rw *RWMutex) Unlock()

// RLock 将 rw 设置为读锁定状态，禁止其他例程写入，但可以读取。
func (rw *RWMutex) RLock()

// Runlock 解除 rw 的读锁定状态，如果 rw 未被读锁顶，则该操作会引发 panic。
func (rw *RWMutex) RUnlock()

// RLocker 返回一个互斥锁，将 rw.RLock 和 rw.RUnlock 封装成了一个 Locker 接口。
func (rw *RWMutex) RLocker() Locker

------------------------------------------------------------
*/
