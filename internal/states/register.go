package states

import (
	"log"
	"time"

	pb "github.com/Irooniam/msg/protos"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RegisterDealer(ID string, host string, port int32, out chan [][]byte) error {
	log.Println(out)
	apb := &pb.ActionMsg{
		Actions: pb.Actions_ADD_DEALER,
	}

	ab, err := proto.Marshal(apb)
	if err != nil {
		return err
	}

	payload := &pb.RegisterDealer{
		Id:     ID,
		Host:   host,
		Port:   int32(port),
		SentAt: timestamppb.New(time.Now()),
	}

	payloadb, err := proto.Marshal(payload)
	if err != nil {
		return err
	}

	out <- [][]byte{ab, payloadb}
	log.Println("register sent")
	return nil
}
