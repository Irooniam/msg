package socks

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Irooniam/msg/conf"
	"github.com/pebbe/zmq4"
)

type ZRouter struct {
	id   string
	In   chan []byte
	Out  chan []byte
	Err  chan error
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
			r.SendMsg([]byte("dealer"), out)
			log.Printf("Out channel - send mess from router socket: %s", string(out))
		case <-r.RecvMsg():
		}
	}
}

func (r *ZRouter) SendMsg(ID []byte, msg []byte) {
	r.sock.SendBytes([]byte(ID), zmq4.SNDMORE)
	r.sock.SendBytes(msg, 0)

}

func (r *ZRouter) RecvMsg() <-chan []byte {
	msg, err := r.sock.RecvMessage(0)
	log.Println("From: ", msg[0], " -- ", msg[2])

	if err != nil {
		log.Println("error on router recvmsg function", msg, err)
		return r.In
	}

	go func() {
		r.In <- []byte(msg[2])
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

	log.Println("new router")
	in := make(chan []byte)
	out := make(chan []byte)
	er := make(chan error)
	done := make(chan bool)
	return &ZRouter{id: ID, In: in, Out: out, Err: er, Done: done, sock: router}, nil
}
