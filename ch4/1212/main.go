package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		fmt.Println("sleep1")
		time.Sleep(5 * time.Second)
		fmt.Println("sleep2")
	}()

	ch <- 1
	ch <- 2
	ch <- 3 // 程序执行5秒的时候会在这行报错(所有goruntine执行完毕)

	time.Sleep(time.Second * 10)
}

/*
sleep1
sleep2
fatal error: all goroutines are asleep - deadlock!

fatal error: all goroutines are asleep - deadlock!

出错信息的意思是：
在main goroutine线，期望从管道中获得一个数据，而这个数据必须是其他goroutine线放入管道的
但是其他goroutine线都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。
所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。
这显然是没有结果的，所以这个程序就说“算了吧，不坚持了，我自己自杀掉，报一个错给代码作者，我被deadlock了”

这里是系统自动在除了主协程之外的协程都关闭后，做的检查，继而报出的错误， 证明思路如下， 在5秒内， 我们看不到异常， 5秒后，系统报错。
*/
