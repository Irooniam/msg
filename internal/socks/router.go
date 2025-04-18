package socks

import (
	"errors"
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4/draft"
)

type ZRouter struct {
	id    string
	In    chan [][]byte
	Out   chan [][]byte
	Done  chan bool
	PDone chan bool //channel for ParseIn Done
	sock  *zmq.Socket
}

func (r *ZRouter) Bind(bindstr string) error {
	log.Println("attempting to bind router socket to ", bindstr)
	err := r.sock.Bind(bindstr)
	if err != nil {
		return err
	}

	log.Println("successfully bound socket to ", bindstr)
	return nil
}

func (r *ZRouter) Run() {
	log.Println("Starting loop for sending / receiving messages on router socket")
	for {
		select {
		case out := <-r.Out:
			log.Printf("sending message out to %s - %s", out[0], out[1])
			r.sendMsg(out[0], out[1], out[2])
		case <-r.RecvMsg():
			log.Println("RecvMsg channel select")
		case <-r.Done:
			log.Println("looks like we are done...")
			return
		case <-time.After(time.Millisecond * 10):
			continue
		}
	}
}

func (r *ZRouter) sendMsg(ID []byte, action []byte, msg []byte) {
	r.sock.SendBytes(ID, zmq.SNDMORE)
	r.sock.SendBytes(action, zmq.SNDMORE)
	r.sock.SendBytes(msg, 0)

}

func (r *ZRouter) RecvMsg() <-chan [][]byte {
	msg, err := r.sock.RecvMessageBytes(zmq.DONTWAIT)
	if err != nil {
		if err.Error() != "resource temporarily unavailable" {
			log.Printf("error on router recvmsg function '%s' - '%s'", msg, err)
		}
		time.Sleep(time.Millisecond * 10)
		return r.In
	}

	/*
		dis/connects events are received as message.  Message is msg[0] = Dealer ID,
		msg[1] = empty slice
	*/

	/*
		if len(msg) == 2 && len(msg[1]) == 0 {
			r.In <- [][]byte{msg[0], []byte(states.ACTIONS["DEALER-EVENT"]), []byte{0x0}}
			return r.In
		}
	*/

	/*
		this is a dis/connect dealer event so only has 2 frames
		append 3rd null
	*/
	if len(msg) == 2 && len(msg[1]) == 0 {
		msg = append(msg, []byte{0x0})
	}

	r.In <- [][]byte{msg[0], msg[1], msg[2]}
	return r.In
}

func NewZRouter(ID string) (*ZRouter, error) {
	router, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		return &ZRouter{}, err
	}

	err = router.SetIdentity(ID)
	if err != nil {
		return &ZRouter{}, errors.New(fmt.Sprintf("Tried setting identity but got error: %s", err))
	}

	//make sure we are notified of dis/connects
	router.SetConnectTimeout(0)
	router.SetRouterNotify(zmq.NotifyConnect | zmq.NotifyDisconnect)

	in := make(chan [][]byte)
	out := make(chan [][]byte)
	done := make(chan bool)
	pdone := make(chan bool)
	log.Println("new router")
	return &ZRouter{id: ID, In: in, Out: out, Done: done, PDone: pdone, sock: router}, nil
}
