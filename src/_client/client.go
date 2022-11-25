package _client

import (
	"context"
	"github.com/junyang7/go-common/src/_client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func Rpc(addr string, header map[string]string, body map[string]string) []byte {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if nil != err {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := pb.NewServiceClient(conn)
	r, err := c.Call(ctx, &pb.Request{
		Header: header,
		Body:   body,
	})
	if nil != err {
		panic(err)
	}
	return r.GetResponse()
}
