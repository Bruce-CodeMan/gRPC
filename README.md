# gRPC

![GoDoc](https://pkg.go.dev/badge/google.golang.org/grpc)

## Why gRPC

- gRPC is a modern open source high performance Remote Procedure Call(RPC) framework that can run in any environment.
- It can efficiently connect services in and across data centers with pluggable support for load balancing,tracing, health checking and authentication.
- It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

## gRPC modes
- Simple RPC
    - The client initiates a request, and the server corresponds to a data
- Server-side streaming RPC
    - The client initiates a request, and the server returns a continuous data flow
    - 案例: 客户端向服务端发送一个股票代码, 服务端就把股票的实时数据源源不断的返回给客户端
- Client-side streaming RPC
    - The client continuously sends data streams to the server, and after sending,the server returns a response
    - 案例: 物联网
- Bidirectional streaming RPC
    - Both client and server can send data streams to each other
    - 案例: 聊天机器人




