package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-go-demo/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现SayHello方法
func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s server) SayAgainHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello Again " + in.GetName()}, nil
}

func main() {

	// 监听端口
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("gRPC server 1 is running on port 8000...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
