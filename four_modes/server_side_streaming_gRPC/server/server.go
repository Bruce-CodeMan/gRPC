/**
 * @Author: Bruce
 * @Description: server_side streaming service
 * @Date: 2023-02-12 15:35
 */

package main

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"gRPC/four_modes/server_side_streaming_gRPC/proto/proto3"
)

type Server struct {
	*proto3.UnimplementedGreeterServer
}

func (s *Server) GetStream(_ *proto3.StreamRequestData, res proto3.Greeter_GetStreamServer) error {
	i := 0
	for {
		err := res.Send(&proto3.StreamResponseData{
			Data: fmt.Sprintf("%v\n", time.Now().Unix()),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
		i++
		if i > 10 {
			break
		}
	}
	return nil
}

func main() {
	g := grpc.NewServer()
	proto3.RegisterGreeterServer(g, &Server{})
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("listen err!!!")
	}
	fmt.Println("start listening!!!")
	if err = g.Serve(l); err != nil {
		panic("serve err!!!")
	}
}
