package main

import (
	"fmt"
	"time"
)

var ErrorRecevier chan map[string]string

func main() {
	go ErrorHandler()

	time.Sleep(time.Second)
}

func ErrorHandler() {


	ErrorRecevier = make(chan map[string]string)

	error := <-ErrorRecevier

	if err := recover(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(error)

}
