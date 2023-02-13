/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/13 4:58 PM
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"

	"gRPC/auth/proto"
)

type Server struct {
	*proto3.UnimplementedGreeterServer
}

func (s *Server) Simple(_ context.Context, req *proto3.StreamRequestData) (*proto3.StreamResponseData, error) {
	return &proto3.StreamResponseData{
		Message: "Hello, " + req.Data,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("收到了一个新的请求")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "认证信息失败")
		}
		var (
			appId  string
			appKey string
		)
		if v1, ok := md["appid"]; ok {
			appId = v1[0]
		}
		if v2, ok := md["appkey"]; ok {
			appKey = v2[0]
		}
		if appId != "1" || appKey != "app-key" {
			return resp, status.Error(codes.Unimplemented, "认证信息失败")
		}
		res, err := handler(ctx, req)
		fmt.Println("拦截流程完成")
		return res, err
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto3.RegisterGreeterServer(g, &Server{})
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败")
	}
	fmt.Println("准备监听1234端口")
	if err = g.Serve(l); err != nil {
		fmt.Println("启动失败")
	}
}
