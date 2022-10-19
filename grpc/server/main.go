package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	trippb "go_study/grpc/server/proto/gen/go"
	trip "go_study/grpc/server/tripservice"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		return
	}
	go startGRPCGateway()  // 启动grpc-gateway服务
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(listen))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption( // grpc-gateway与json互相转换时候时候的选项
		runtime.MIMEWildcard, &runtime.JSONPb{
		EnumsAsInts: true, // 枚举类型字符串转换成数值
		OrigName: true,  // 字段定义时的原始名称
	}))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()},
	)
	if err !=nil{
		log.Fatalf("cannot start grpc gateway : %v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err !=nil{
		log.Fatalf("cannot listen and server : %v", err)
	}
}
