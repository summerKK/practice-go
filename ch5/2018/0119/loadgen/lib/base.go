package lib

import "time"

//调用结果的结构
type CallResult struct {
	//ID
	Id int64
	//原生请求
	Req RawReq
	//原生响应
	Resp RawResp
	//响应代码
	Code ResultCode
	//结果成因的简述
	Msg string
	//耗时
	Elaspe time.Duration
}

//荷载发生器接口
type Generator interface {
	//启动荷载发生器
	Start()
	//停止何在发生器
	//第一个结果值代表已发荷载总数,且仅在第二个结果值为true时有效
	//第二个结果值代表是否成功将荷载发生器转变为已停止状态
	Stop() (uint64, bool)
	//获取状态
	Status() GenStatus
}

//原生请求的结构
type RawReq struct {
	Id  int64
	Req []byte
}

//原生响应的结构
type RawResp struct {
	Id     int64
	Resp   []byte
	Err    error
	Elaspe time.Duration
}

type ResultCode int

//荷载发生器的状态的类型
type GenStatus int

const (
	STATUS_ORIGINAL GenStatus = 0
	STATUS_STARTED  GenStatus = 1
	STATUS_STOPPED  GenStatus = 2
)
