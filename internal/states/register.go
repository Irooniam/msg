package states

import (
	"errors"
	"fmt"
	"log"
	"time"

	pb "github.com/Irooniam/msg/protos"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// craft the proto msg and send out to router via channel
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

func ParseAction(b []byte) (string, error) {
	var actionMsg pb.ActionMsg
	if err := proto.Unmarshal(b, &actionMsg); err != nil {
		return "", errors.New(fmt.Sprintf("Unable to Unmarshal actions %s", err))
	}

	/*
		we dont unmarshall payload because each action
		has its own protos
	*/
	log.Println("action ", actionMsg.Actions)
	return actionMsg.Actions.String(), nil
}

func AddDealer(ID []byte, payload []byte) error {
	var dealer pb.RegisterDealer
	if err := proto.Unmarshal(payload, &dealer); err != nil {
		return errors.New(fmt.Sprintf("Unable to Unmarshal RegisterDeal %s", err))
	}

	DEALERS.Store(string(ID),
		DealerInfo{
			ID:   dealer.Id,
			Host: dealer.Host,
			Port: int(dealer.Port),
		},
	)
	return nil
}
