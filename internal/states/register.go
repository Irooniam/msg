package states

import (
	"log"
	"sync"
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

// make sure to change ID to string in map lookup
func RemoveDealer(m *sync.Map, ID []byte, out chan [][]byte) {
	m.Delete(string(ID))
}

func AddDealer(m *sync.Map, ID []byte, out chan [][]byte) error {
	m.Store(string(ID),
		DealerInfo{
			ID: string(ID),
		},
	)
	return nil
}

/*
*
All we know is that there was an event.
Have to check map to see if dealer exists = disconnect
Dealer doesnt exist = connect
*
*/
func DealerEvent(m *sync.Map, ID []byte, out chan [][]byte) {
	_, ok := m.Load(string(ID))

	//dont have this dealer in DEALER map
	if !ok {
		log.Println("dont have dealer so we add ", string(ID))
		AddDealer(m, ID, out)
		return
	}

	//dealers exists so remove them from DEALER map
	log.Println("already have dealer so remove ", string(ID))
	RemoveDealer(m, ID, out)
}
