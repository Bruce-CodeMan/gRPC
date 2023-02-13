/**
 * @Author: Bruce
 * @Description: client interceptor
 * @Date: 2023/2/13 4:27 PM
 */
package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gRPC/interceptor/proto"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, res, cc, opts...)
		fmt.Println("时间间隔: ", time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("127.0.0.1:1234", opts...)
	if err != nil {
		fmt.Println("拨号失败")
	}
	defer func() {
		if err = conn.Close(); err != nil {
			fmt.Println("关闭失败")
		}
	}()
	c := proto3.NewGreeterClient(conn)
	r, err := c.Simple(context.Background(), &proto3.StreamRequestData{
		Data: "Bruce",
	})
	if err != nil {
		fmt.Println("调用失败")
	}
	fmt.Println(r.Message)
}
