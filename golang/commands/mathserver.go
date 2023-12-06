package commands

import (
	"github.com/abdullahchand/gRPC-Protobuf/golang/proto"
)

// Creation of the MathServer type that will inherit the interface functions for the server
type MathServer struct {
	proto.UnimplementedAdditionServer
}