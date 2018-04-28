package main

import (
	pb "protobuf/grpc"
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
)

var Address = "127.0.0.1:9999"

type service struct {
}

func (h service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	log.Println("welcome hello client !")
	resp.Message = "Hello " + in.Name + "."
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()
	svc := &service{}
	// 注册HelloService
	pb.RegisterHelloServer(s, svc)

	log.Println("Listen on " + Address)

	s.Serve(listen)
}