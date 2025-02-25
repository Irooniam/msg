package socks

import (
	"errors"
	"fmt"
	"log"

	"github.com/pebbe/zmq4"
)

type ZRouter struct {
	ID   string
	In   chan []byte
	Out  chan []byte
	Err  chan error
	Done chan bool
	Sock *zmq4.Socket
}

func (r *ZRouter) Bind(ip string, port int) error {
	conn := fmt.Sprintf("%s:%d", ip, port)
	err := r.Sock.Bind(conn)
	if err != nil {
		return err
	}

	return nil
}

func (r *ZRouter) Run() {
loop:
	for {
		select {
		case msg := <-r.In:
			fmt.Print(msg)

		case err := <-r.Err:
			fmt.Printf("Got error %s", err)
		case <-r.Done:
			log.Println("received done signal")
			break loop
		}
	}
}

func NewRouter(ID string) (*ZRouter, error) {
	router, err := zmq4.NewSocket(zmq4.ROUTER)
	if err != nil {
		return &ZRouter{}, err
	}

	err = router.SetIdentity(ID)
	if err != nil {
		return &ZRouter{}, errors.New(fmt.Sprintf("Tried setting identity but got error: %s", err))
	}

	in := make(chan []byte)
	out := make(chan []byte)
	er := make(chan error)
	done := make(chan bool)
	return &ZRouter{ID: ID, In: in, Out: out, Err: er, Done: done, Sock: router}, nil
}
