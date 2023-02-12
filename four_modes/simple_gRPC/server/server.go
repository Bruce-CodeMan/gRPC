/**
 * @Author: Bruce
 * @Description: receive the requests
 * @Date: 2023-02-12 15:18
 */
package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"gRPC/four_modes/simple_gRPC/proto/proto3"
)

type Server struct {
	*proto3.UnimplementedGreetServer
}

func (s *Server) SayHello(ctx context.Context, req *proto3.HelloRequest) (*proto3.HelloReply, error) {
	return &proto3.HelloReply{
		Message: "Hello, " + req.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto3.RegisterGreetServer(g, &Server{})
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("listen err!!!")
	}
	fmt.Println("starting listening")
	if err = g.Serve(l); err != nil {
		panic("serve err!!!")
	}
}
