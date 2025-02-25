package socks

import (
	"errors"
	"fmt"

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

func (d *ZRouter) Bind(port int) error {
	conn := fmt.Sprintf("0.0.0.0:%d", port)
	err := d.Sock.Bind(conn)
	if err != nil {
		return err
	}

	return nil
}

func NewRouter(ID string) (*ZRouter, error) {
	dealer, err := zmq4.NewSocket(zmq4.DEALER)
	if err != nil {
		return &ZRouter{}, err
	}

	err = dealer.SetIdentity(ID)
	if err != nil {
		return &ZRouter{}, errors.New(fmt.Sprintf("Tried setting identity but got error: %s", err))
	}

	in := make(chan []byte)
	out := make(chan []byte)
	er := make(chan error)
	done := make(chan bool)
	return &ZRouter{ID: ID, In: in, Out: out, Err: er, Done: done, Sock: dealer}, nil
}
