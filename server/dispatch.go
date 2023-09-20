package server

import (
	"context"
	grpcgo "github.com/hc6078/go-protobuf/grpc-go"
)

type Dispatch struct {
	grpcgo.UnimplementedDispatchServer
}

func (d *Dispatch) GetDispatch(context.Context, *grpcgo.DispatchReq) (*grpcgo.DispatchRes, error) {
	res := new(grpcgo.DispatchRes)
	return res, nil
}
