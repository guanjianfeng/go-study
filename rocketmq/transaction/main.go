package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderList struct {
}

func (OrderList) ExecuteLocalTransaction(*primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地逻辑")
	time.Sleep(10*time.Second)
	fmt.Println("执行本地逻辑完成")
	return primitive.CommitMessageState
}

func (OrderList) CheckLocalTransaction(*primitive.MessageExt) primitive.LocalTransactionState {
	return primitive.CommitMessageState
}

func main() {
	p, err := rocketmq.NewTransactionProducer(
		&OrderList{},
		producer.WithNameServer([]string{"59.110.64.200:9876"}))
	if err != nil {
		panic("生成producer失败")
	}
	err = p.Start()
	if err != nil {
		panic("启动producer失败")
	}
	res, err := p.SendMessageInTransaction(context.Background(), &primitive.Message{
		Topic: "transaction",
		Body:  []byte("this is transaction test" ),
	})
	if err != nil {
		fmt.Printf("发送失败 %s", err.Error())
		return
	}
	fmt.Printf("发送成功 %s", res.String())
	time.Sleep(time.Hour)
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("关闭失败 %s", err.Error())
	}
}
