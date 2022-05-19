package main

import (
	"context"
	"flag"
	pb "go_demo/grpc_server_stream/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", "server_stream", "name to stream server")
)

func main() {
	flag.Parse()
	// 连接grpc服务器
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	// 初始化StreamServer服务客户端
	c := pb.NewStreamServerClient(conn)

	// 调用ListValue接口
	r, err := c.ListValue(context.Background(), &pb.SimpleRequest{Data: *name})
	if err != nil {
		log.Fatalf("could not stream:%v", err)
	}

	for {
		// 接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := r.Recv()

		// 接收服务端消息流结束
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("listen get stream err: %v", err)
		}

		log.Println(res.StreamValue)
	}
}
