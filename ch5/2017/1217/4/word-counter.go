package main

import (
	"sync"
	"os"
	"fmt"
	"bufio"
	"strings"
)

type words struct {
	found map[string]int
	sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	w := newWords()

	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(file)
	}

	wg.Wait()

	for word, count := range w.found {
		fmt.Printf("%s %d\n", word, count)
	}

}

func tallyWords(s string, words *words) error {
	file, err := os.Open(s)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		words.add(word, 1)
	}
	return scanner.Err()
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}
