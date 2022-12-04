package _client

import (
	"context"
	"github.com/junyang7/go-common/src/_client/pb"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func Rpc(addr string, header map[string]string, body map[string]string) []byte {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrGrpcDial).
		Message(err).
		Do()
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := pb.NewServiceClient(conn).Call(ctx, &pb.Request{
		Header: header,
		Body:   body,
	})
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrPbNewServiceClientCall).
		Message(err).
		Do()
	return r.GetResponse()
}
