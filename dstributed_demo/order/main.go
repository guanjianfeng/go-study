package main

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"go_study/dstributed_demo/order/handle"
	"go_study/dstributed_demo/order/initialize"
	"go_study/dstributed_demo/order/proto"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	//"syscall"
)

func main() {
	//IP := flag.String("ip", "0.0.0.0", "ip地址")
	//Port := flag.Int("port", 50051, "端口号")

	//初始化
	//initialize.InitLogger()
	//initialize.InitConfig()
	initialize.InitDB()
	initialize.InitSrvConn()
	initialize.InitMq()
	//zap.S().Info(global.ServerConfig)

	listen, err := net.Listen("tcp", "localhost:8086")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	proto.RegisterOrderServer(s, &handle.OrderServer{})
	fmt.Println("开启订单服务")
	fmt.Println(s.Serve(listen))

	//监听订单超时topic
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"59.110.64.200:9876"}),
		consumer.WithGroupName("mxshop-order"),
	)

	if err := c.Subscribe("order_timeout", consumer.MessageSelector{}, handle.OrderTimeout); err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主goroutine退出

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_ = c.Shutdown()

}
