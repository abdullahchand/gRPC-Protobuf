package commands

import (
	"log"
	"github.com/abdullahchand/gRPC-Protobuf/golang/proto"
)


// Assigning MathServer Type the AddStreamer function and overriding the function from the generated files
func (math_server MathServer) AddStreamer(numbers *proto.Request, streamer proto.Addition_AddStreamerServer)  error {
	log.Printf("Stream request recieved...")

	for i := 1; i <= 3; i++ {
		res:= &proto.Response{Result: numbers.A + numbers.B}
		
		if err := streamer.Send(res); err != nil {
			return err
		}
	}
	
	return nil
}