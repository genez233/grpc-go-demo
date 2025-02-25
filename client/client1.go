package main

import (
	"context"
	"github.com/genez233/go-utils/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-demo/proto"
	"log"
	"time"
)

var (
	logger = &glog.GLog{}
)

func main() {
	logger = glog.New(&glog.Config{
		ServerName:       "grpc client 1",
		Version:          "1.0.0",
		ConsoleLog:       true,
		IsUpload:         true,
		RunMode:          "DEBUG",
		LogUrl:           "http://logs.zhiyunai.com.cn/api/default/%s/_json",
		OpenobserveToken: "Basic MTIyNTg0MjkwNUBxcS5jb206QTlNVXdOQm14NWJNejlPTQ==",
	})

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 使用 grpc.DialContext 连接服务器
	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 调用 SayHello 方法
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %s", res.GetMessage())
}
