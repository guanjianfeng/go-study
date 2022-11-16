package main

import (
	"fmt"
	"go_study/dstributed_demo/goods/handle"
	"go_study/dstributed_demo/goods/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	proto.RegisterGoodsServer(s, &handle.GoodsServer{})
	fmt.Println("开启商品服务")
	fmt.Println(s.Serve(listen))
}
