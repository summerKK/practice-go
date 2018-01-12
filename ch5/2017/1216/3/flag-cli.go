package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "world", "A name to say Hello to .")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "use spanish language")
	flag.BoolVar(&spanish, "s", false, "use spanish language")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s !\n", *name)
	} else {
		fmt.Printf("Hello %s !\n", *name)
	}
}
