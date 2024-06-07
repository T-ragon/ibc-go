package main

import (
	"context"
	"fmt"
	pb "github.com/T-ragon/ibc-go/v9/RPC"
	V1 "github.com/T-ragon/ibc-go/v9/proto/ibc/lightclients/zkp/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) MustEmbedUnimplementedGreetsServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	ap := V1.AggregatePacket{}
	err := ap.XXX_Unmarshal(in.Message.Value)
	if err != nil {
		return nil, err
	}
	log.Println(ap.Signer)

	n := &pb.Wrapper{
		Value: &wrapperspb.StringValue{
			Value: "long"}}
	m := &pb.Wrapper{Value: &wrapperspb.StringValue{
		Value: "I am eating",
	}}
	//将wrapper消息打包成Any
	anyMsg, _ := anypb.New(n)
	me, _ := anypb.New(m)
	return &pb.HelloReply{
		Name:    anyMsg,
		Message: me,
	}, nil
}

func main() {
	//协议类型，以及ip，port
	lis, err := net.Listen("tcp", ":8012")
	if err != nil {
		fmt.Println(err)
		return
	}

	//定义RPC的server
	server := grpc.NewServer()
	//注册服务，相当于注册SayHello接口
	pb.RegisterGreetsServer(server, &Server{})

	//进行映射绑定
	reflection.Register(server)

	//启动服务
	err = server.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
