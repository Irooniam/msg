package socks

import (
	"errors"
	"fmt"
	"log"

	"github.com/pebbe/zmq4"
)

type ZDealer struct {
	ID   string
	In   chan []byte
	Out  chan []byte
	Sock *zmq4.Socket
}

func (d *ZDealer) Connect(port int) error {
	conn := fmt.Sprintf("0.0.0.0:%d", port)
	err := d.Sock.Bind(conn)
	if err != nil {
		return err
	}

	return nil
}

func NewDealer(ID string) (*ZDealer, error) {
	dealer, err := zmq4.NewSocket(zmq4.DEALER)
	if err != nil {
		return &ZDealer{}, err
	}

	err = dealer.SetIdentity(ID)
	if err != nil {
		return &ZDealer{}, errors.New(fmt.Sprintf("Tried setting identity but got error: %s", err))
	}

	log.Println("this is new dealer")
	in := make(chan []byte)
	out := make(chan []byte)

	return &ZDealer{ID: ID, In: in, Out: out, Sock: dealer}, nil
}
