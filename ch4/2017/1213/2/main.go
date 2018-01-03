package main

import "fmt"

type byteCounter int

func main() {
	var c byteCounter
	n, _ := c.Write([]byte("Hello World"))
	fmt.Println(n)
	fmt.Println(c)
	c.Write([]byte("summer"))
	fmt.Println(c)
	c = 0
	fmt.Fprint(&c,"hello,%s","summer")
	fmt.Println(c)
}

func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}
