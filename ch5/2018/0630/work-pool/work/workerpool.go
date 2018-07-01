package work

type WorkerPool struct {
	workers []*Worker
	pool    chan *Worker
}

//构建工作者池
func NewWorkerPool(maxWorkers int) *WorkerPool {
	w := &WorkerPool{
		workers: make([]*Worker, maxWorkers),
		pool:    make(chan *Worker, maxWorkers),
	}

	//初始化工作者
	for i, _ := range w.workers {
		worker := NewWorker(0)
		w.workers[i] = worker
		w.pool <- worker
	}

	return w
}

//启动工作者
func (w *WorkerPool) Start() {
	for _, worker := range w.workers {
		worker.Start()
	}
}

//停止工作者
func (w *WorkerPool) Stop() {
	for _, worker := range w.workers {
		worker.Stop()
	}
}

//获取工作者(阻塞),当所有工作者都在处理任务时,会阻塞等待有可用的工作者可用
func (w *WorkerPool) Get() *Worker {
	return <-w.pool
}

//返回工作者
func (w *WorkerPool) Put(worker *Worker) {
	w.pool <- worker
}
