package states

import (
	pb "github.com/Irooniam/msg/protos"
	"google.golang.org/protobuf/proto"
)

func RegisterDealer() ([]byte, error) {
	msg := &pb.ActionMsg{
		Actions: pb.Actions_ADD_DEALER,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return data, nil
}
