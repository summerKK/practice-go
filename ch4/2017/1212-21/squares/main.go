package main

import "fmt"

func main() {

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("----------------")

	var x int = 10
	fmt.Println(&x)
	add(x)
	fmt.Println(x)

}

func squares() func() int {
	var x int
	fmt.Println(&x)
	return func() int {
		x++
		fmt.Println(&x)
		return x * x
	}
}

func add(x int) {
	fmt.Println(&x)
	x++
}
