/**
 * @Author: Bruce
 * @Description: 描述
 * @Date: 2023-02-12 15:44
 */

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gRPC/four_modes/server_side_streaming_gRPC/proto/proto3"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Dial error !!!")
	}
	defer func() {
		if err = conn.Close(); err != nil {
			panic("close error!!!")
		}
	}()
	c := proto3.NewGreeterClient(conn)
	r, err := c.GetStream(context.Background(), &proto3.StreamRequestData{Data: "start"})
	if err != nil {
		panic("Get stream from server Error !!!")
	}
	for {
		res, err := r.Recv()
		if err != nil {
			panic("Receive Stream Error !!!")
		}
		fmt.Println(res.Data)
	}
}
