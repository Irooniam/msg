package states_test

import (
	"testing"

	"github.com/Irooniam/msg/internal/states"
	"github.com/Irooniam/msg/protos"
	"google.golang.org/protobuf/proto"
)

func TestRegisterDealerHappyPath(t *testing.T) {
	msg, err := states.RegisterDealer()
	if err != nil {
		t.Errorf("was not expecting error but got %s", err)
	}

	pmsg := &protos.ActionMsg{}
	err = proto.Unmarshal(msg, pmsg)
	if err != nil {
		t.Errorf("was not expecting error when unmarshaling but got %s", err)
	}

	if pmsg.GetActions().String() != protos.Actions_ADD_DEALER.String() {
		t.Errorf("was expecting un/marshalled actions to be same but %s != %s", pmsg.Actions.String(), protos.Actions_ADD_DEALER.String())
	}
}
