/**
 * @Author: Bruce
 * @Description: recv/send the stream
 * @Date: 2023-02-12 21:24
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gRPC/four_modes/bidirectional_side_streaming_gRPC/proto/proto3"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("拨号失败")
	}
	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Println("关闭失败")
		}
	}()
	c := proto3.NewGreeterClient(conn)
	var allStr proto3.Greeter_AllStreamClient
	if allStr, err = c.AllStream(context.Background()); err != nil {
		fmt.Println("连接失败")
	}
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, err := allStr.Recv()
			if err != nil {
				fmt.Println("接收失败")
				break
			}
			fmt.Println(data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := allStr.Send(&proto3.StreamRequestData{
				Data: "我是客户端",
			})
			if err != nil {
				fmt.Println("发送失败")
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
