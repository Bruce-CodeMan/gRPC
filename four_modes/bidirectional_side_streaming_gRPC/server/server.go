/**
 * @Author: Bruce
 * @Description: send/recv the stream
 * @Date: 2023-02-12 21:17
 */
package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"gRPC/four_modes/bidirectional_side_streaming_gRPC/proto/proto3"
)

type Server struct {
	*proto3.UnimplementedGreeterServer
}

func (s *Server) AllStream(client proto3.Greeter_AllStreamServer) error {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, err := client.Recv()
			if err != nil {
				fmt.Println("接收失败")
				break
			} else {
				fmt.Println(data.Data)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if err := client.Send(&proto3.StreamResponseData{
				Data: "我是服务器",
			}); err != nil {
				fmt.Println("发送消息失败")

			}
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	g := grpc.NewServer()
	proto3.RegisterGreeterServer(g, &Server{})
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败")
	}
	fmt.Println("开始启动服务，准备监听1234端口")
	if err = g.Serve(l); err != nil {
		fmt.Println("启动失败")
	}
}
