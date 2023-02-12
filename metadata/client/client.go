/**
 * @Author: Bruce
 * @Description: client: send message to the server
 * @Date: 2023-02-12 21:54
 */

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"gRPC/metadata/proto/proto3"
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
	md := metadata.New(map[string]string{
		"first_name": "Bruce",
		"last_name":  "Hsu",
	})
	c := proto3.NewGreeterClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.Simple(ctx, &proto3.StreamRequestData{
		Data: "client Data",
	})
	if err != nil {
		fmt.Println("发送失败")
		return
	}
	fmt.Println(r.Data)
}
