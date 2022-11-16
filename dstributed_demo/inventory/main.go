package main

import (
	"fmt"
	"go_study/dstributed_demo/inventory/handle"
	"go_study/dstributed_demo/inventory/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	proto.RegisterInventoryServer(s, &handle.InventoryServer{})
	fmt.Println("开启库存服务")
	fmt.Println(s.Serve(listen))

	//监听库存归还topic
	//c, _ := rocketmq.NewPushConsumer(
	//	consumer.WithNameServer([]string{"59.110.64.200:9876"}),
	//	consumer.WithGroupName("mxshop-inventory"),
	//)
	//
	//if err := c.Subscribe("order_reback", consumer.MessageSelector{},handle.AutoReback); err != nil {
	//	fmt.Println("读取消息失败")
	//}
	//_ = c.Start()
	////不能让主goroutine退出
	//
	////接收终止信号
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//_ = c.Shutdown()
	////if err = register_client.DeRegister(serviceId); err != nil {
	////	zap.S().Info("注销失败:", err.Error())
	////}else{
	////	zap.S().Info("注销成功:")
	////}
}
