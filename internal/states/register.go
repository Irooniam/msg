package states

import (
	"log"
)

/*
*
To register Dealer we only need their ID
Dealers connect to Router, they themselves dont listen
*
*/
func RegisterDealer(ID string, out chan [][]byte) error {
	action := []byte(ACTIONS["REGISTER-DEALER"])
	out <- [][]byte{action, []byte(ID)}
	log.Println("register sent")
	return nil
}

func AddDealer(ID []byte, payload []byte) error {
	DEALERS.Store(string(ID),
		DealerInfo{
			ID: string(ID),
		},
	)
	return nil
}
