package main

import (
	"context"
	"fmt"
	"github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	//创建rpc连接
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

	client := types.NewMsgClient(conn)
	//设置超时时间
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//尝试反序列化 client state
	//var clientState exported.ClientState
	//clientFilleName := "./clientstate.json"
	//contents, err := os.ReadFile(clientFilleName)

	req := &types.MsgCreateClient{
		ClientState:    nil,
		ConsensusState: nil,
		Signer:         "Along",
	}

	reply, err := client.CreateClient(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(reply.String())
}
