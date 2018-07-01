package work

import (
	"sync"
	"fmt"
)

type Worker struct {
	job  chan interface{}
	quit chan bool
	wg   sync.WaitGroup
}

func NewWorker(maxJobs int) *Worker {
	return &Worker{
		job:  make(chan interface{}, maxJobs),
		quit: make(chan bool),
	}
}

//开始程序
func (w *Worker) Start() {
	w.wg.Add(1)

	go func() {
		defer w.wg.Done()

		for {
			select {
			case job := <-w.job:
				//处理任务
				fmt.Println("job:", job)
			case <-w.quit:
				//程序退出
				return
			}
		}
	}()
}

//关闭程序
func (w *Worker) Stop() {
	w.quit <- true
	//等待所有线程工作全部完成
	w.wg.Wait()
}

//接受任务
func (w *Worker) AddJob(job interface{}) {
	w.job <- job
}
