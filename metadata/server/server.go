/**
 * @Author: Bruce
 * @Description: server: receive the message + metadata from the client
 * @Date: 2023-02-12 21:49
 */
package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"gRPC/metadata/proto/proto3"
)

type Server struct {
	*proto3.UnimplementedGreeterServer
}

func (s *Server) Simple(ctx context.Context, req *proto3.StreamRequestData) (*proto3.StreamResponseData, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("Get metadata ERROR!!!")
	}
	for k, v := range md {
		fmt.Println(k, v)
	}
	return &proto3.StreamResponseData{
		Data: "Receive Data:" + req.Data,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto3.RegisterGreeterServer(g, &Server{})
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败")
	}
	fmt.Println("开始监听1234端口")
	if err = g.Serve(l); err != nil {
		fmt.Println("启动失败")
	}
}
