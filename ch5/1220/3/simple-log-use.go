package main

import (
	"log"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Hello World")
	}()
	log.Println("This is a regular message.")
	log.Fatalln("This is a fatal message.")
	log.Println("This is the end of the function.")

}
