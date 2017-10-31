//go语言实现线程池
package threadPool

import "fmt"

type GorountinePool struct {
	Queue chan func() error
	//开的goruntine个数
	Number int
	//任务总个数
	Total int
	//任务结果
	result chan error
	//任务结束执行的回调函数
	finishCallback func()
}

//初始化
func (self *GorountinePool) Init(number, total int) {
	self.Queue = make(chan func() error, total)
	self.Number = number
	self.Total = total
	self.result = make(chan error, total)
}

//任务开始
func (self *GorountinePool) Start() {
	//开启number个goruntine
	for i := 0; i < self.Number; i++ {
		go func() {
			for {
				//从队列取出任务
				task, ok := <-self.Queue
				if !ok {
					break
				}
				//执行任务
				err := task()
				//把任务结果返回给result chan
				self.result <- err
			}
		}()
	}

	//获取每个任务的结果
	for j := 0; j < self.Total; j++ {
		res, ok := <-self.result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	//所有任务执行完毕后的回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}
}

//关闭任务
func (self *GorountinePool) Stop() {
	close(self.Queue)
	close(self.result)
}

//添加任务
func (self *GorountinePool) AddTask(task func() error) {
	self.Queue <- task
}

//设置结束回调函数
func (self *GorountinePool) SetFinishCallback(callback func()) {
	self.finishCallback = callback
}
