package main

import (
	"fmt"
	"github.com/cosmos/ibc-go/v8/exportedrpc"
	"github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	//协议类型，以及ip，port
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
		return
	}

	//定义RPC的server
	server := grpc.NewServer()
	types.RegisterMsgServer(server, &exportedrpc.LongServer{})

	reflection.Register(server)

	//启动服务
	err = server.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
