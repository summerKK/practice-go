package main

import (
	"practice/ch4/2017/1211/subdir"

	"fmt"
	"time"
)

func main() {
	go subdir.Start()

	go func() {
		fmt.Println(<-subdir.Ch)
	}()

	time.Sleep(time.Second)

}
