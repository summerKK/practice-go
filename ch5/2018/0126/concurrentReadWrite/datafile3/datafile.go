package datafile

import (
	"os"
	"sync"
	"errors"
	"io"
	"sync/atomic"
)

//数据文件的接口类型
type DataFile interface {
	//读取数据块
	Read() (rsn int64, d Data, err error)
	//写入一个数据块
	Write(d Data) (wsn int64, err error)
	//获取最后读取的数据块的序列号
	Rsn() int64
	//获取最后写入的数据块的序列号
	Wsn() int64
	//获取数据块的长度
	DataLen() uint32
}

//数据文件的实现类型
type myDataFile struct {
	//文件
	f *os.File
	//文件的读写锁
	fmutex sync.RWMutex
	//写操作的偏移量
	woffset int64
	//读操作的偏移量
	roffset int64
	rcond   *sync.Cond
	//数据块长度
	dataLen uint32
}

type Data []byte

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func (dataFile *myDataFile) Read() (rsn int64, d Data, err error) {
	//读取并更新写偏移量
	var offset int64
	for {
		offset = atomic.LoadInt64(&dataFile.roffset)
		if atomic.CompareAndSwapInt64(&dataFile.roffset, offset, offset+int64(dataFile.dataLen)) {
			break
		}
	}

	//读取一个数据块
	//rsn是读取的数据的编号,他总会是一个整数,因为 offset += int64(dataFile.dataLen)
	rsn = offset / int64(dataFile.dataLen)
	bytes := make([]byte, dataFile.dataLen)
	dataFile.fmutex.RLock()
	defer dataFile.fmutex.RUnlock()
	for {
		_, err := dataFile.f.ReadAt(bytes, offset)
		if err != nil {
			//如果读取到文件边界会尝试进行再次读取,直到成功为止
			//这里使用for循环的目的是为了保证每次返回的rsn序列都是对的
			// 假如读到io.EOF直接返回,那下次的rsn的值就会增加,会遗漏一些数据
			if err == io.EOF {
				//.Wait()函数在调用时一定要确保已经获取了其成员变量锁L ,因为Wait第一件事就是解锁。　
				// 但是需要注意的是，当wait()结束等待返回之前,它会重新对Ｌ进行加锁，也就是说，当wait结束，goruntine仍然会获取lock。
				dataFile.rcond.Wait()
				continue
			}
			return
		}
		d = bytes
		return
	}
}

func (dataFile *myDataFile) Write(d Data) (wsn int64, err error) {
	//读取并更新写偏移量
	var offset int64
	for {
		offset = atomic.LoadInt64(&dataFile.woffset)
		if atomic.CompareAndSwapInt64(&dataFile.woffset, offset, offset+int64(dataFile.dataLen)) {
			break
		}
	}

	//写入一个数据块
	wsn = offset / int64(dataFile.dataLen)
	var bytes []byte
	if len(d) > int(dataFile.dataLen) {
		bytes = d[0:dataFile.dataLen];
	} else {
		bytes = d
	}
	dataFile.fmutex.Lock()
	defer dataFile.fmutex.Unlock()
	_, err = dataFile.f.Write(bytes)
	//每次写入一次发送一个信号
	dataFile.rcond.Signal()
	return
}

func (dataFile *myDataFile) Rsn() int64 {
	offset := atomic.LoadInt64(&dataFile.roffset)
	return offset / int64(dataFile.dataLen)
}

func (dataFile *myDataFile) Wsn() int64 {
	offset := atomic.LoadInt64(&dataFile.woffset)
	return offset / int64(dataFile.dataLen)
}

func (dataFile *myDataFile) DataLen() uint32 {
	return dataFile.dataLen
}
