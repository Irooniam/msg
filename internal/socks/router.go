package socks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Irooniam/msg/conf"
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
			r.sendMsg(out[0], out[1])
		case msg := <-r.RecvMsg():
			log.Printf("received message on router socket. From %s - payload %s", msg[0], msg[1])
		case <-r.Done:
			log.Println("looks like we are done...")
			return
		}
	}
}

func (r *ZRouter) sendMsg(ID []byte, msg []byte) {
	r.sock.SendBytes(ID, zmq4.SNDMORE)
	r.sock.SendBytes(msg, 0)

}

func (r *ZRouter) RecvMsg() chan [][]byte {
	msg, err := r.sock.RecvMessageBytes(zmq4.DONTWAIT)

	if err != nil {
		log.Printf("error on router recvmsg function '%s' - '%s'", msg, err)
		time.Sleep(time.Millisecond * 10)
		return r.In
	}

	go func() {
		r.In <- [][]byte{msg[0], msg[1]}
	}()
	return r.In
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
