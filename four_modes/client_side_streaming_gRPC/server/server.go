/**
 * @Author: Bruce
 * @Description: receive the stream
 * @Date: 2023-02-12 20:27
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

func (s *Server) PullStream(client proto3.Greeter_PullStreamServer) error {
	for {
		if res, err := client.Recv(); err != nil {
			fmt.Println("接收失败")
			fmt.Println(err)
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
		fmt.Println(err)
	}
	fmt.Println("监听1234端口")
	if err = g.Serve(l); err != nil {
		fmt.Println("启动服务失败")
		fmt.Println(err)
	}
}
