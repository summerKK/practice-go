package pool

import (
	"sync"
	"io"
	"errors"
	"log"
)

type Pool struct {
	//互斥锁,保证并发安全
	m sync.Mutex
	//资源池是否关闭
	closed bool
	//资源池
	poolChan chan io.Closer
	//资源生成器
	factory func() (io.Closer, error)
}

var (
	poolClosed = errors.New("资源池已被关闭")
)

func NewPool(fn func() (io.Closer, error), size uint32) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("资源池大小必须大于0")
	}
	return &Pool{
		poolChan: make(chan io.Closer, size),
		factory:  fn,
	}, nil
}

func (pool *Pool) Take() (io.Closer, error) {
	select {
	case r, ok := <-pool.poolChan:
		if !ok {
			return nil, poolClosed
		}
		log.Println("池内资源")
		return r, nil
	default:
		//默认生成一个新的
		log.Println("生成新资源")
		return pool.factory()
	}
}

func (pool *Pool) Return(r io.Closer) error {

	pool.m.Lock()
	defer pool.m.Unlock()

	//判断资源池是否被关闭
	if pool.closed {
		//资源池已关闭,关闭这个资源
		r.Close()
		return poolClosed
	}

	select {
	case pool.poolChan <- r:
		log.Println("资源放入池中")
	default:
		r.Close()
		log.Println("资源池已满,释放资源")
	}

	return nil
}

func (pool *Pool) Close() {
	//保证并发安全
	pool.m.Lock()
	defer pool.m.Unlock()

	//判断资源池是否已经关闭
	if pool.closed {
		return
	}

	pool.closed = true

	//关闭资源池
	close(pool.poolChan)

	//释放池内资源
	for r := range pool.poolChan {
		r.Close()
	}
}
