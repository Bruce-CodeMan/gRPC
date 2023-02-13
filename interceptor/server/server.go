/**
 * @Author: Bruce
 * @Description: server interceptor
 * @Date: 2023/2/13 4:18 PM
 */

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	"gRPC/interceptor/proto"
)

type Server struct {
	*proto3.UnimplementedGreeterServer
}

func (s *Server) Simple(_ context.Context, req *proto3.StreamRequestData) (*proto3.StreamResponseData, error) {
	return &proto3.StreamResponseData{
		Message: req.Data,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
		fmt.Println("接收到了一个新的请求")
		return handler(ctx, req)
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
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
