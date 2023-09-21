package server

import (
	"context"
	arpc "github.com/vrieske/go-artifacts/rpc"
	grpcgo "github.com/vrieske/go-protobuf/grpc-go"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	arpc.GetGrpcServerInstance().Run(ctx, func(server *grpc.Server) {
		grpcgo.RegisterDispatchServer(server, &Dispatch{})
	})
}
