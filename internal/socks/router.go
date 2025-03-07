package socks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/states"
	"github.com/Irooniam/msg/protos"
	"github.com/pebbe/zmq4"
)

type ZRouter struct {
	id   string
	In   chan [][]byte
	Out  chan [][]byte
	Done chan bool
	sock *zmq4.Socket
}

func ChkRouterConf() error {
	if os.Getenv(conf.MSG_DIR_HOST) == "" {
		return errors.New(fmt.Sprintf("env var for directory host %s is not set", conf.MSG_DIR_HOST))
	}

	if os.Getenv(conf.MSG_DIR_PORT) == "" {
		return errors.New(fmt.Sprintf("env var for directory port %s is not set", conf.MSG_DIR_PORT))
	}

	return nil
}

func (r *ZRouter) Bind(bindstr string) error {
	conn := fmt.Sprintf("tcp://%s", bindstr)
	log.Println("attempting to bind router socket to ", conn)
	err := r.sock.Bind(conn)
	if err != nil {
		return err
	}

	log.Println("successfully bound socket to ", conn)
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
	r.sock.SendBytes(ID, zmq4.SNDMORE)
	r.sock.SendBytes(action, zmq4.SNDMORE)
	r.sock.SendBytes(msg, 0)

}

func (r *ZRouter) RecvMsg() <-chan [][]byte {
	msg, err := r.sock.RecvMessageBytes(zmq4.DONTWAIT)
	if err != nil {
		if err.Error() != "resource temporarily unavailable" {
			log.Printf("error on router recvmsg function '%s' - '%s'", msg, err)
		}
		time.Sleep(time.Millisecond * 10)
		return r.In
	}

	r.In <- [][]byte{msg[0], msg[1], msg[2]}
	return r.In
}

/*
*

	Responsible for parsing all incoming messages from
	Router socket

	We can turn the biz login in the loop into goroutine - leaving for now

*
*/
func (r *ZRouter) ParseIn() {
	for {
		msg := <-r.In

		//have to receive 3 parts: header, action, payload
		if len(msg) != 3 {
			log.Printf("expected msg received to be 3 parts but is: %d", len(msg))
			continue
		}

		action, err := states.ParseAction(msg[1])
		if err != nil {
			log.Printf("tried getting action from msg but got %s", err)
			continue
		}

		switch action {
		case protos.Actions_ADD_DEALER.String():
			log.Println("we are adding dealer yo")
		default:
			log.Printf("actions is %s - and we dont have match", action)

		}

		log.Println("Parse incoming message action ", msg, action)
	}
}

func NewZRouter(ID string) (*ZRouter, error) {
	router, err := zmq4.NewSocket(zmq4.ROUTER)
	if err != nil {
		return &ZRouter{}, err
	}

	err = router.SetIdentity(ID)
	if err != nil {
		return &ZRouter{}, errors.New(fmt.Sprintf("Tried setting identity but got error: %s", err))
	}

	in := make(chan [][]byte)
	out := make(chan [][]byte)
	done := make(chan bool)
	log.Println("new router")
	return &ZRouter{id: ID, In: in, Out: out, Done: done, sock: router}, nil
}
