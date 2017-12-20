package main

import (
	"errors"
	"github.com/Masterminds/cookoo/safely"
	"time"
)

func main() {
	safely.Go(message)
	println("Outside goruninte")
	time.Sleep(time.Second * 1)
}

func message() {
	println("Inside goruntine")
	panic(errors.New("Oops!"))
}
