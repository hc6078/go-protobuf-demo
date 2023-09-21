package client

import (
	"context"
	"github.com/vrieske/go-artifacts/rpc"
	grpcGo "github.com/vrieske/go-protobuf/grpc-go"
	"sync"
	"time"
)

var dispatchClientOnce sync.Once

var client *DispatchRpcClient

type DispatchRpcClient struct {
	Client grpcGo.DispatchClient
}

func NewDispatchRpcClient() *DispatchRpcClient {
	dispatchClientOnce.Do(func() {
		addr := "127.0.0.1:8080"
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		conn := rpc.NewGrpcClient(ctx, addr)
		client = &DispatchRpcClient{
			Client: grpcGo.NewDispatchClient(conn)}
	})
	return client
}

func (dc *DispatchRpcClient) queryByDispatch(ctx context.Context, domain, ip string) ([]string, string, error) {
	dispatch, err := dc.Client.GetDispatch(ctx, &grpcGo.DispatchReq{
		Domain: domain,
		Ip:     ip,
	})
	if err != nil {
		return []string{}, "", err
	}
	return dispatch.Ips, dispatch.Cname, err
}

func main() {
	ctx := context.Background()
	cli := NewDispatchRpcClient()
	_, _, err := cli.queryByDispatch(ctx, "", "")
	if err != nil {
		return
	}
}
