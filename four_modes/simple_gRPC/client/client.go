/**
 * @Author: Bruce
 * @Description: 描述
 * @Date: 2023-02-12 15:25
 */

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gRPC/four_modes/simple_gRPC/proto/proto3"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("dial err!!!")
	}
	defer func() {
		if err = conn.Close(); err != nil {
			panic("conn err!!!")
		}
	}()
	c := proto3.NewGreetClient(conn)
	r, err := c.SayHello(context.Background(), &proto3.HelloRequest{Name: "Bruce"})
	if err != nil {
		panic("Remote Procedure Call err!!!")
	}
	fmt.Println(r.Message)
}
