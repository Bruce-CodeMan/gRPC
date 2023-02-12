/**
 * @Author: Bruce
 * @Description: client put the stream to server
 * @Date: 2023-02-12 21:01
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
	}
	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Println("关闭失败")
		}
	}()
	c := proto3.NewGreeterClient(conn)
	putStream, err := c.PutStream(context.Background())
	if err != nil {
		fmt.Println("推送信息失败")
	}
	i := 0
	for {
		if err = putStream.Send(&proto3.StreamRequestData{
			Data: fmt.Sprintf("%v\n", time.Now().Unix()),
		}); err != nil {
			fmt.Println("推送信息失败")
		}
		i++
		time.Sleep(time.Second * 1)
		fmt.Printf("第%d次发送数据\n", i)
		if i > 10 {
			fmt.Println("推送信息完成")
			break
		}
	}
}
