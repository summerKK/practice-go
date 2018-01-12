package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"sort"
	"os"
)

type su1 struct {
	id      uint32
	summary string
}

func main() {


	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= 6 })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}


	os.Exit(1)


	m := map[string]string{
		"1": "2",
		"2": "2",
		"3": "2",
		"4": "2",
	}

	enc, _ := json.Marshal(m)
	fmt.Println(string(enc))

	type su1 struct {
		id      uint32
		summary string
	}

	err := errors.New("hello world")
	fmt.Println([]error{err})

	s := Newsu()
	s.hello()

}

type Interface interface {
	hello()
}

type su struct {
	s string
}

func Newsu() Interface {
	return &su{}
}

func (s *su) hello() {
	s.s = "summer"
	fmt.Printf("%p\n",s)
	fmt.Println(&s)
	fmt.Println(&s.s)
}
