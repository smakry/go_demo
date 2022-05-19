package main

import (
	"flag"
	"fmt"
	pb "go_demo/grpc_server_stream/proto"
	"log"
	"math"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "the server port")
)

type server struct {
	pb.UnimplementedStreamServerServer
}

func (s *server) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for i := 0; i < 5; i++ {
		err := srv.Send(&pb.StreamReply{
			StreamValue: fmt.Sprintf("%v:%v", req.Data, strconv.Itoa(i)),
		})

		// 间隔2s下发一次，客户端上下文不可带Deadline
		time.Sleep(time.Second * 2)

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	// 实例化grpc服务端，最大发送消息长度默认值 math.MaxInt32 ，最大接收消息长度 2M
	s := grpc.NewServer(grpc.MaxSendMsgSize(math.MaxInt32), grpc.MaxRecvMsgSize(1024*1024*2))

	// 注册服务
	pb.RegisterStreamServerServer(s, &server{})
	log.Printf("server listening at: %v", lis.Addr())

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server:%v", err)
	}
}
