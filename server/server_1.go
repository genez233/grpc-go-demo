package main

import (
	"context"
	"github.com/genez233/go-utils/glog"
	"google.golang.org/grpc"
	pb "grpc-go-demo/proto"
	"log"
	"net"
)

var (
	logger = &glog.GLog{}
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {

	logger = glog.New(&glog.Config{
		ServerName:       "grpc server 1",
		Version:          "1.0.0",
		ConsoleLog:       true,
		IsUpload:         true,
		RunMode:          "DEBUG",
		LogUrl:           "http://logs.zhiyunai.com.cn/api/default/%s/_json",
		OpenobserveToken: "Basic MTIyNTg0MjkwNUBxcS5jb206QTlNVXdOQm14NWJNejlPTQ==",
	})

	// 监听端口
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		logger.Error("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	logger.Info("gRPC server 1 is running on port 8000...")
	if err := s.Serve(lis); err != nil {
		logger.Error("failed to serve: %v", err)
	}
}
