package main

import (
	"fmt"
)

type ByteCounter int

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	var name = "Bob"
	fmt.Fprintf(&c, "Hello,%s", name)
	fmt.Println(c)
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}
