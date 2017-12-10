package main

import (
	"net/http"
	"runtime/pprof"
	"sync"
	//"fmt"
	"fmt"
)

var (
	task = make(chan int)
	chs  = make(chan int, 20)
	wg   = sync.WaitGroup{}
)

func f(i int) {
	defer func() {
		<-chs
		wg.Done()
	}()
	task <- i
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {

	go func() {
		for i := 0; i < 100000; i++ {
			chs <- 0
			wg.Add(1)
			go f(i)
		}
		wg.Wait()
		close(task)
	}()

	for v := range task {
		fmt.Println(v)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":11811", nil)

}
