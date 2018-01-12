package main

import (
	"sync"
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	aLen := len(a)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(start int) {
			item := make([]int, 0, aLen/4)
			for j := start; j < aLen; j += 4 {
				item = append(item, a[j])
			}
			fmt.Println(item)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
