package main

import (
	"errors"
	"math/rand"
	"fmt"
)

var ErrTimeOut = errors.New("The request time out")
var ErrRejected = errors.New("The request was rejected")

var random = rand.New(rand.NewSource(35))

func main() {
	resp, err := SendRequest("Hello")
	if err == ErrTimeOut {
		fmt.Println("Timeout. Retrying.")
		resp, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
func SendRequest(s string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrTimeOut
	default:
		return "", ErrRejected
	}
}
