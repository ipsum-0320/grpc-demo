package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/server_pb"
	"net"
)

type server struct {
	server_pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *server_pb.HelloRequest) (*server_pb.HelloResponse, error) {
	fmt.Printf(in.Age)
	return &server_pb.HelloResponse{Reply: "Hello " + in.Name, Status: "200"}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	server_pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
