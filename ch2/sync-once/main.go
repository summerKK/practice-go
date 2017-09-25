package main

import (
	"sync"
	"fmt"
	"time"
)

//sync.Once用法
func main() {
	//引用
	o := &sync.Once{}
	go oHandle(o)
	go oHandle(o)
	time.Sleep(2 * time.Second)
}

func oHandle(o *sync.Once) {
	fmt.Println("Start do")
	o.Do(func() {
		//执行一次
		fmt.Println("Doing something")
	})
	fmt.Println("Done")
}

/*
Start do
Doing something
Done
Start do
Done
 */
