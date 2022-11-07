package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"59.110.64.200:9876"}),
		consumer.WithGroupName("group1"), //防止重复消费
	)
	err := c.Subscribe(
		"test",
		consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i:= range msgs{
				fmt.Printf("获取到值：%v\n",msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		},
	)
	if err != nil{
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	time.Sleep(time.Hour) // 防止主线程退出
	_ = c.Shutdown()
}
