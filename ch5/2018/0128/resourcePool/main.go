//模拟资源池操作
package main

import (
	"log"
	"io"
	"sync/atomic"
	"practice/ch5/2018/0128/resourcePool/pool"
	"time"
	"math/rand"
	"sync"
)

//数据库链接
type dbConnection struct {
	//唯一标识符
	id uint32
}

var (
	dbSymbol     uint32
	maxGoruntine int    = 5
	poolSize     uint32 = 10
)

func (db *dbConnection) Close() error {
	log.Println("关闭db连接")
	return nil
}

func createConnection() (io.Closer, error) {
	atomic.AddUint32(&dbSymbol, 1)
	return &dbConnection{id: dbSymbol}, nil
}

//模拟数据库查询
func dbQuery(query int, pool *pool.Pool) {
	//从资源池取出一个连接
	conn, err := pool.Take()
	if err != nil {
		log.Println(err)
		return
	}
	//放回资源池
	defer pool.Return(conn)

	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).id)

}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoruntine)

	pool, err := pool.NewPool(createConnection, poolSize)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < maxGoruntine; i++ {
		go func(query int) {
			dbQuery(query, pool)
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("开始关闭资源池")
	pool.Close()
}
