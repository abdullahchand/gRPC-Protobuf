package commands

import (
	"github.com/abdullahchand/gRPC-Protobuf/golang/proto"
)

type MathServer struct {
	proto.UnimplementedAdditionServer
}