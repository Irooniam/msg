package states

import (
	"log"

	pb "github.com/Irooniam/msg/protos"
	"google.golang.org/protobuf/proto"
)

func RegisterDealer(out chan [][]byte) error {
	msg := &pb.ActionMsg{
		Actions: pb.Actions_ADD_DEALER,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	log.Println(msg.String())
	out <- [][]byte{data, []byte("payload")}
	log.Println("register sent")
	return nil
}
