package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc-demo/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "浩瀚星河"})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	log.Printf("服务端响应: %s", resp.Message)
}
