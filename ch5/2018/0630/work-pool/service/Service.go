package service

import (
	"practice/ch5/2018/0630/work-pool/work"
	"sync"
)

type Service struct {
	workders *work.WorkerPool
	jobs     chan interface{}
	maxJobs  int
	wg       sync.WaitGroup
}

func NewService(maxWorkers, maxJobs int) *Service {
	return &Service{
		workders: work.NewWorkerPool(maxWorkers),
		jobs:     make(chan interface{}, maxJobs),
	}
}

func (s *Service) Start() {
	s.wg.Add(1)
	s.workders.Start()

	go func() {
		defer s.wg.Done()
		for job := range s.jobs {
			go func(job interface{}) {
				//从工作池取出一个工作者
				worker := s.workders.Get()

				//完成任务后返回工作池
				defer s.workders.Put(worker)

				//提交任务处理(异步)
				worker.AddJob(job)

			}(job)
		}
	}()
}

//停止服务
func (s *Service) Stop() {
	s.workders.Stop()
	close(s.jobs)
	//等待所有线程完成
	s.wg.Wait()
}

//提交任务,任务管道带较大的缓存,延缓阻塞时间
func (s *Service) AddJob(job interface{}) {
	s.jobs <- job
}
