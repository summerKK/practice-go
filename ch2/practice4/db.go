package main

import (
	"github.com/boltdb/bolt"
	"time"
	"log"
	"fmt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	walk := make(chan struct{})

	go func() {

		time.Sleep(3*time.Second)
		fmt.Println("sleep")
		close(walk)
	}()
	fmt.Println("watting")
	fmt.Println(<-walk)
}
