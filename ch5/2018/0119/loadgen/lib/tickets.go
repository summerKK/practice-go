package lib

import (
	"fmt"
	"errors"
)

//goruntine票池的接口
type GoTickets interface {
	//获得一张票
	Take()
	//归还一张票
	Return()
	//票池是否已被激活
	Active() bool
	//票的总数
	Total() uint32
	//剩余的票数
	Remainder() uint32
}

//goruntine票池的实现
type myGoTickets struct {
	//票的总数
	total uint32
	//票的容器
	ticketCh chan byte
	//票池是否已被激活
	active bool
}

func NewGoTickets(total uint32) (GoTickets, error) {
	tickets := &myGoTickets{}
	if !tickets.init(total) {
		errMsg := fmt.Sprintf("The goroutine ticket pool can not be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return tickets, nil
}

func (tickets *myGoTickets) init(total uint32) bool {
	if tickets.active {
		return false
	}
	if total == 0 {
		return false
	}
	ch := make(chan byte, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- 1
	}
	tickets.ticketCh = ch
	tickets.total = total
	tickets.active = true
	return true
}

func (tickets *myGoTickets) Take() {
	<-tickets.ticketCh
}

func (tickets *myGoTickets) Return() {
	tickets.ticketCh <- 1
}

func (tickets *myGoTickets) Active() bool {
	return tickets.active
}

func (tickets *myGoTickets) Total() uint32 {
	return tickets.total
}

func (tickets *myGoTickets) Remainder() uint32 {
	return uint32(len(tickets.ticketCh))
}
