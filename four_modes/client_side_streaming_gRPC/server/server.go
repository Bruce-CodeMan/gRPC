/**
 * @Author: Bruce
 * @Description: server pull the stream from client
 * @Date: 2023-02-12 20:57
 */

package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"gRPC/four_modes/client_side_streaming_gRPC/proto/proto3"
)

type Server struct {
	*proto3.UnimplementedGreeterServer
}

func (s *Server) PutStream(cli proto3.Greeter_PutStreamServer) error {
	for {
		if res, err := cli.Recv(); err != nil {
			fmt.Println("接收失败")
			break
		} else {
			fmt.Println(res.Data)
		}
	}
	return nil
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
