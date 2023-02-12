/**
 * @Author: Bruce
 * @Description: pull the stream
 * @Date: 2023-02-12 20:33
 */
package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gRPC/four_modes/client_side_streaming_gRPC/proto/proto3"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("拨号失败")
		fmt.Println(err)
	}

	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Println("关闭连接失败")
		}
	}()

	c := proto3.NewGreeterClient(conn)
	putStream, _ := c.PullStream(context.Background())
	i := 0
	for {
		i++
		err := putStream.Send(&proto3.StreamRequestData{
			Data: fmt.Sprintf("%v\n", time.Now().Unix()),
		})
		if err != nil {
			fmt.Println("发送数据失败")
			break
		}
		time.Sleep(time.Second * 1)
		if i > 10 {
			fmt.Println("数据发送结束")
			break
		}
	}
}
