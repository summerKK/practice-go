package sub

import "fmt"

var (
	Receiver chan int
)

func Handle() {
	Receiver = make(chan int)
	for {
		msg := <-Receiver
		fmt.Println(msg)
	}
}
