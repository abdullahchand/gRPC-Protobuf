package commands

import (
	"context"
	"github.com/abdullahchand/gRPC-Protobuf/golang/proto"
)


func (math_server MathServer) Add(ctx context.Context, numbers *proto.Request) (*proto.Response, error) {
	res := numbers.A + numbers.B
	return &proto.Response{Result: res}, nil
}