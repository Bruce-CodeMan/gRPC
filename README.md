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

## metadata
```text
gRPC 让我们可以像本地调用一样实现远程调用，对于每一次的RPC调用，都可能会有一些有用的数据，而这些数据就可以通过metadata来传递
metadata是以key-value的形式存储数据的，其中key是string类型，而value是[]string，就是一个字符串数组类型
metadata使得client 和server能够为对方提供关于本次调用的一些信息，就像一次http请求的RequestHeader和ResponseHeader一样
http中header的生命周期是一次http请求，那么metadata的生命周期就是一次rpc调用
```


