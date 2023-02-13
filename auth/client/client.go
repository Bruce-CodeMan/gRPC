/**
 * @Author: Bruce
 * @Description:
 * @Date: 2023/2/13 5:41 PM
 */
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gRPC/auth/proto"
)

type Custom struct{}

func (c Custom) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "2",
		"appKey": "app-key",
	}, nil
}

func (c Custom) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(Custom{}))

	conn, err := grpc.Dial("127.0.0.1:1234", opts...)
	if err != nil {
		fmt.Println("拨号失败")
		return
	}
	defer conn.Close()
	c := proto3.NewGreeterClient(conn)
	r, err := c.Simple(context.Background(), &proto3.StreamRequestData{
		Data: "Bruce",
	})
	if err != nil {
		fmt.Println("调用远程函数失败")
	}
	fmt.Println(r.Message)
}
