/*
pprof是个神马玩意儿？

pprof - manual page for pprof (part of gperftools)

是gperftools工具的一部分

gperftools又是啥？

These tools are for use by developers so that they can create more robust applications. Especially of use to those developing multi-threaded applications in C++ with templates. Includes TCMalloc, heap-checker, heap-profiler and cpu-profiler.

一个性能分析的工具，可以查看堆栈、cpu信息等等。



在golang中如何使用呢？下面就来看看。
 */
package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	flag.Parse()

	//这里实现了远程获取pprof数据的接口
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}

	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		time.Sleep(time.Millisecond * 100)
		counter++
	}
	wg.Done()
}

/*
通过浏览器访问

http://localhost:6060/debug/pprof/
*/
