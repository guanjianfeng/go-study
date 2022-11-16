package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_study/dstributed_demo/order/proto"
	"google.golang.org/grpc"
)

var OrderSrvClient proto.OrderClient

func init() {
	//初始化订单服务连接
	orderConn, err := grpc.Dial(
		"localhost:8086",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		fmt.Println("[InitSrvConn] 连接 【订单服务失败】")
	}

	OrderSrvClient = proto.NewOrderClient(orderConn)
	fmt.Println("订单服务开启")
}
func main() {
	//var ctx *gin.Context
	fmt.Println("开始请求")
	rsp, err := OrderSrvClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  int32(1),
		Name:    "gjf",
		Mobile:  "18610000001",
		Address: "南山",
		Post:    "请尽快发货",
	})
	fmt.Println("请求结束")
	fmt.Printf("%+v\n", rsp)
	fmt.Printf("%+v\n", err)

	if err != nil {
		zap.S().Errorw("新建订单失败")
		zap.S().Errorw(err.Error())
		return
	}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"id": rsp.Id,
	//})
	//zap.S().Info("id=", rsp.Id)
}
