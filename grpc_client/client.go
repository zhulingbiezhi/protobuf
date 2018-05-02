package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	pb "protobuf/grpc"
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
	// new request context
	ctx := context.Background()
	// add key-value pairs of metadata to context
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("appid", "101010", "appkey", "key"),
	)
	r, err := c.SayHello(ctx, reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	log.Println(r.Message)
}
