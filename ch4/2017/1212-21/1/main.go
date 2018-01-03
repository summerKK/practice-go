package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squares
	go func() {
		for {
			x, ok := <-naturals
			if ok {
				squares <- x * x
			} else {
				close(squares)
				return
			}
		}
	}()

	for res := range squares {
		fmt.Println(res)
	}

	test3()

	res, _ := test4()
	fmt.Println(res)

	r := test5()
	fmt.Println(r)

}

func test3() {
	ch := make(chan struct{})
	filenames := map[string]string{
		"1": "2",
		"2": "22",
		"3": "222",
		"4": "2222",
		"5": "22222",
	}

	for res, _ := range filenames {
		go func(res string) {
			fmt.Println(res)
			ch <- struct{}{}
		}(res)
	}

	for range filenames {
		<-ch
	}
}

func test4() (thumbfiles []string, err error) {

	type item struct {
		filename string
		err      error
	}

	filenames := map[string]string{
		"1": "2",
		"2": "22",
		"3": "222",
		"4": "2222",
		"5": "22222",
	}

	ch := make(chan item, len(filenames))
	for f := range filenames {
		go func(f string) {
			var it item
			fmt.Println(f)
			_, err := os.Create(f)
			it.filename, it.err = f, err
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.filename)
	}

	return thumbfiles, nil

}

func test5() int64 {
	nums := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 10,}
	sum := make(chan int64)
	var wg sync.WaitGroup

	for _, v := range nums {
		wg.Add(1)
		go func(v int64) {
			defer wg.Done()
			sum <- v
		}(v)
	}

	go func() {
		wg.Wait()
		close(sum)
	}()

	var total int64
	for s := range sum {
		total += s
	}

	return total
}
