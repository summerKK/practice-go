package main

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"time"
)

type signalHandle func(s os.Signal, arg interface{})

type signalSet struct {
	m map[os.Signal]signalHandle
}

func signalSetNew() (*signalSet) {
	ss := new(signalSet)
	ss.m = make(map[os.Signal]signalHandle)
	return ss
}

func (set *signalSet) register(s os.Signal, handler signalHandle) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}

}

func (set *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}

	panic("won't reach here")
}

func sysSignalHandleDemo() {
	ss := signalSetNew()
	handler := func(s os.Signal, arg interface{}) {
		fmt.Printf("handle signal:%v\n", s)
	}
	ss.register(syscall.SIGINT, handler)
	ss.register(syscall.SIGQUIT, handler)
	ss.register(syscall.SIGABRT, handler)

	for {
		c := make(chan os.Signal)
		var sigs []os.Signal
		for sig := range ss.m {
			sigs = append(sigs, sig)
		}
		signal.Notify(c)
		sig := <-c

		err := ss.handle(sig, nil)
		if (err != nil) {
			fmt.Printf("unknown signal received: %v\n", sig)
			os.Exit(1)
		}
	}

}

func main() {
	go sysSignalHandleDemo()
	time.Sleep(time.Hour)
}
