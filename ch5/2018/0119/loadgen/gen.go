package loadgen

import (
	"time"
	"practice/ch5/2018/0119/loadgen/lib"
	"summerWebCrawler/logging"
	"fmt"
	"errors"
	"math"
)

type myGenerator struct {
	//调用器
	caller lib.Caller
	//响应超时时间.单位:纳秒
	timeoutNs time.Duration
	//每秒荷载发送量
	lps uint32
	//负载持续时间单位:纳秒
	durationNs time.Duration
	//调用结果通道
	resultCh chan *lib.CallResult
	//并发量
	concurrency uint32
	//goroutine票池
	tickets lib.GoTickets
	//停止信号的传递通道
	stopSign chan byte
	// 取消发送后续结果的信号。
	cancelSign byte
	//状态
	status lib.GenStatus
}

var (
	logger logging.Logger
)

func init() {
	logger = logging.NewSimpleLogger()
}

func NewGenerator(
	caller lib.Caller,
	timeoutNs time.Duration,
	lps uint32,
	durationNs time.Duration,
	resultCh chan *lib.CallResult) (lib.Generator, error) {
	logger.Infoln("New a load generator...")
	logger.Infoln("Checking the parameters")

	var errMsg string

	if caller == nil {
		errMsg = fmt.Sprintf("Invalid caller!")
	}
	if timeoutNs == 0 {
		errMsg = fmt.Sprintf("Invalid timeoutNs!")
	}
	if lps == 0 {
		errMsg = fmt.Sprintf("Invalid lps(load per second)!")
	}
	if durationNs == 0 {
		errMsg = fmt.Sprintf("Invalid durationNs!")
	}
	if resultCh == nil {
		errMsg = fmt.Sprintf("Invalid result channel!")
	}
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}

	gen := &myGenerator{
		caller:     caller,
		timeoutNs:  timeoutNs,
		lps:        lps,
		durationNs: durationNs,
		resultCh:   resultCh,
		stopSign:   make(chan byte, 1),
		cancelSign: 0,
		status:     lib.STATUS_ORIGINAL,
	}

	logger.Infoln("passed. (timeoutNs=%v,lps=%d,durationNs=%v)", timeoutNs, lps, durationNs)
	err := gen.init()
	if err != nil {
		return nil, err
	}
	return gen, nil
}

func (tickets *myGenerator) init() error {
	logger.Infoln("Initializing the load generator...")
	//并发量 ≈ 单个载荷的响应超时时间 / 载荷的发送间隔时间(1e9 / lps)
	//最后+1代表了某一个时间周期之初向被测试软件发送的那个载荷
	var total64 int64 = int64(tickets.timeoutNs)/int64(1e9/tickets.lps) + 1
	if total64 > math.MaxInt32 {
		total64 = math.MaxInt32
	}
	tickets.concurrency = uint32(total64)
	tick, err := lib.NewGoTickets(tickets.concurrency)
	if err != nil {
		return err
	}
	tickets.tickets = tick
	logger.Infoln("Initialized.(concurrency=%d)", tickets.concurrency)
	return nil
}

func (tickets *myGenerator) Start() {

}

func (tickets *myGenerator) Stop() (uint64, bool) {

}

func (tickets *myGenerator) Status() lib.GenStatus {

}
