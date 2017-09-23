package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	var gs [5]struct {
		id     int
		result int
	}
	for i := 0; i < len(gs); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			gs[i].id = i
			gs[i].result = (i + 1) * 100
		}(i)
	}
	wg.Wait()

	//多个goruntine实现local storage
	fmt.Printf("%#v", gs)
}
