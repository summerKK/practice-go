package main

import (
	"os"
	"syscall"
	"fmt"
	"os/signal"
	"sync"
	"time"
)

func main() {
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1)

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s...[sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1:%s\n", sig)
		}
		fmt.Printf("End. [sigRecv1]\n")
		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2:%s\n", sig)
		}
		fmt.Printf("End. [sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("wait for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Printf("stop notification...\n ")
	signal.Stop(sigRecv1)
	close(sigRecv1)

	wg.Wait()

}
