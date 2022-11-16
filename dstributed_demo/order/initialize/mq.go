package initialize

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
	"go_study/dstributed_demo/order/global"
)

func InitMq() {
	//发送延时消息
	var err error
	global.MqDelay, err = rocketmq.NewProducer(producer.WithNameServer([]string{"59.110.64.200:9876"}),
		producer.WithGroupName("mxshop-order"))

	if err != nil {
		fmt.Println("延迟消息链接失败")
	}
	if err = global.MqDelay.Start(); err != nil {
		zap.S().Errorf("启动producer失败 %s", err.Error())
	}

	global.MqInTimeOut, err = rocketmq.NewProducer(producer.WithNameServer([]string{"59.110.64.200:9876"}))
	if err != nil {
		panic("生成producer失败")
	}

	if err = global.MqInTimeOut.Start(); err != nil {
		zap.S().Errorf("启动MqInTimeOut producer失败 %s", err.Error())
	}

}