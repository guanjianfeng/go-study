package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"strconv"
)

func main()  {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"59.110.64.200:9876"}))
	if err != nil{
		panic("生成producer失败")
	}
	err = p.Start()
	if err != nil{
		panic("启动producer失败")
	}
	for i:=0;i<10;i++{
		res, err := p.SendSync(context.Background(), &primitive.Message{
			Topic: "test",
			Body: []byte("this is new test" + strconv.Itoa(i)),
		})
		if err != nil {
			fmt.Printf("发送失败 %s", err.Error())
			return
		}
		fmt.Printf("发送成功 %s", res.String())
	}

	err = p.Shutdown()
	if err != nil {
		fmt.Printf("关闭失败 %s", err.Error())
	}
}