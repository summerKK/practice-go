package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 11; i++ {
		//Add操作尽量放在goruntine外面执行
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("go(%d)run\n", i)
		}(i)
	}

	fmt.Println("watting")
	wg.Wait()
	fmt.Println("done")
}
