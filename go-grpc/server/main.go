package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/proto"
	"log"
	"net"
)

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	reply := fmt.Sprintf("你好，%s！", req.Name)
	return &proto.HelloReply{Message: reply}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &greeterServer{})
	log.Println("gRPC 服务端已启动，端口: 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动失败: %v", err)
	}
}
