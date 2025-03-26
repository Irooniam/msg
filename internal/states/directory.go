package states

import (
	"log"
	"sync"
)

/*
*
We register WS - not just dealer zmq socket
*
*/

type DirRegistry struct {
	Dealers *sync.Map
}

// make sure to change ID to string in map lookup
func (d *DirRegistry) RemoveDealer(m *sync.Map, ID []byte, out chan [][]byte) {
	d.Dealers.Delete(string(ID))
}

func RegisterWS(ID string, out chan [][]byte) error {
	action := []byte(ACTIONS["REGISTER-DEALER"])
	out <- [][]byte{action, []byte(ID)}
	log.Println("register sent")
	return nil
}

// make sure to change ID to string in map lookup
func RemoveDealer(m *sync.Map, ID []byte, out chan [][]byte) {
	m.Delete(string(ID))
}

func AddService(m *sync.Map, info ServiceInfo, out chan [][]byte) error {
	m.Store(info.ID, info)
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
