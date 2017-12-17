package main

import (
	"time"
	"os"
	"fmt"
	"log"
)

func main() {
	os.Stdout.WriteString("hello world\n")
	done := time.After(time.Second * 15)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("time out")
			os.Exit(0)
		}
	}
}
func readStdin(echo chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, err := os.Stdin.Read(data)
		if err != nil{
			log.Fatal(err)
		}
		if l > 0 {
			echo <- data
		}
	}
}
