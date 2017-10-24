package main

import (
	"fmt"
	"net/url"
	"strings"
)

var ch chan int

func main() {

	s := "https://img.aitaotu.cc:8089/Pics/2015/1104/18/01.jpg"

	parseUrl,_ := url.Parse(s)

	fmt.Println(parseUrl.Path)

	tmp := strings.TrimLeft(parseUrl.Path,"/")

	fmt.Println(tmp)

	data := [...]int{1,2,3,4,5,6}

	for _,d := range data {
		v := d
		fmt.Printf("%v",&v)
		fmt.Printf("  %v",&d)
		fmt.Println()
	}

	x := 1
	defer func(a int) { fmt.Println("a=", a) }(x)
	defer func() { fmt.Println("x=", x) }()
	defer fmt.Println("c=",x)
	x++

	arr := new([10]int)

	arr1 := []int{}

	map1 := map[string]int{}

	fmt.Printf("\n%T\n",arr)

	test(arr)

	fmt.Println(arr)

	test1(arr1)

	fmt.Println(arr1)

	test2(map1)

	fmt.Println(map1)
}

func test(arr *[10]int)  {
	*arr = [10]int{1,2,3,4}
}

func test1(arr []int)  {
	arr = []int{9,8,7,6,5}
}

func test2(map1 map[string]int) {
	map1["summer"] = 1
}

