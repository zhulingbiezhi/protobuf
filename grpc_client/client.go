package main

import (
	pb "protobuf/grpc"
	"google.golang.org/grpc"
	"context"
	"log"
	"google.golang.org/grpc/grpclog"
)

var Address = "127.0.0.1:9999"

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	log.Println(r.Message)
}