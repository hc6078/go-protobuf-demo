package server

import (
	"context"
	arpc "github.com/hc6078/go-artifacts/rpc"
	grpcgo "github.com/hc6078/go-protobuf/grpc-go"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	arpc.GetGrpcServerInstance().Run(ctx, func(server *grpc.Server) {
		grpcgo.RegisterDispatchServer(server, &Dispatch{})
	})
}
