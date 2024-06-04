package main

import (
	"context"
	"fmt"
	pb "github.com/cosmos/ibc-go/v8/RPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"time"
)

func main() {
	//创建rpc链接
	conn, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	//创建客户端
	client := pb.NewGreetsClient(conn)
	//设置超时时间
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	n, _ := anypb.New(&pb.Wrapper{Value: &wrapperspb.StringValue{Value: "X"}})
	m, _ := anypb.New(&pb.Wrapper{Value: &wrapperspb.StringValue{Value: "Do you want to eat?"}})
	//调用方法
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name:    n,
		Message: m,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(reply.Name, reply.Message)
}
