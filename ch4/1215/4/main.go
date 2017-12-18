package main

import (
	"fmt"
	"runtime"
	"regexp"
	"strings"
)

func main() {
	test1()

	fmt.Println("hello world")

	s := "https://a.1stdibscdn.com/archivesE/upload/v_435/1513013950788/IMG_0624_master.PNG?width=768"
	regex := regexp.MustCompile(`(?i)(\w|\d|_)+\.(jpg|png|jpeg)`)

	res := regex.FindStringSubmatch(s)
	fmt.Println(res)

	m := map[string]string{"name": "summer", "sex": "m'an", "age": "1'8"}

	key := []string{}
	val := []string{}
	prepare := []string{}
	for i, v := range m {
		key = append(key, i)
		val = append(val, "'"+strings.Replace(v, "'", "\\'", -1)+"'")
		prepare = append(prepare, "?")
	}

	s = `insert into users (` + strings.Join(key, ",") + `) values (` + strings.Join(prepare, ",") + `)`

	fmt.Println(s)

	fmt.Println(key, val)
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
