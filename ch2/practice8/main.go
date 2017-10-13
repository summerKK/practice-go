package main

import (
	"fmt"
)

var ch chan int

func main() {

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

