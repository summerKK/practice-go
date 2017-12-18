package main

import (
	"fmt"
	"runtime"
	"regexp"
)

func main() {
	test1()

	fmt.Println("hello world")

	s := "https://a.1stdibscdn.com/archivesE/upload/v_435/1513013950788/IMG_0624_master.PNG?width=768"
	regex := regexp.MustCompile(`(?i)(\w|\d|_)+\.(jpg|png|jpeg)`)

	res := regex.FindStringSubmatch(s)
	fmt.Println(res)
}

func test1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch error", err)
		}
	}()

	//panic("error")

	fmt.Println("hello world")

	fmt.Println(runtime.GOOS)
}
