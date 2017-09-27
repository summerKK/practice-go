package main

import "fmt"

func main() {
	ch := make(chan struct{})
	close(ch)

	mapp := map[string]chan struct{} {
		"summer":ch,
	}
	mapp["summer"] <- struct{}{}
	fmt.Println(<-mapp["summer"])
}
/*
panic: send on closed channel

goroutine 1 [running]:
panic(0x49aa60, 0xc0420381d0)
	C:/Go/src/runtime/panic.go:500 +0x1af
main.main()
	D:/goprojects/src/practice/ch2/practice1/main.go:12 +0x127

Process finished with exit code 2
 */