package main

import (
	"practice/ch5/2018/0630/publish-and-subscribe/pubsub"
	"time"
	"strings"
	"fmt"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	//全部主题
	all := p.Subscribe()
	//单个主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("Hello,World!")
	p.Publish("Hello,golang!")

	go func() {
		for msg := range all {
			fmt.Println(msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println(msg)
		}
	}()

	time.Sleep(3 * time.Second)

}
