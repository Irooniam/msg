package socks

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Irooniam/msg/conf"
	"github.com/pebbe/zmq4"
)

type ZDealer struct {
	ID   string
	In   chan []byte
	Out  chan []byte
	Done chan bool
	sock *zmq4.Socket
}

// centralize config checking here
func ChkDealerConf() error {
	if os.Getenv(conf.MSG_DIR_HOST) == "" {
		return errors.New(fmt.Sprintf("env var for directory host %s is not set", conf.MSG_DIR_HOST))
	}

	if os.Getenv(conf.MSG_DIR_PORT) == "" {
		return errors.New(fmt.Sprintf("env var for directory port %s is not set", conf.MSG_DIR_PORT))
	}

	if os.Getenv(conf.MSG_DEALER_HOST) == "" {
		return errors.New(fmt.Sprintf("env var for dealer host %s is not set", conf.MSG_DEALER_HOST))
	}

	if os.Getenv(conf.MSG_DEALER_PORT) == "" {
		return errors.New(fmt.Sprintf("env var for dealer port %s is not set", conf.MSG_DEALER_PORT))
	}

	if os.Getenv(conf.MSG_DEALER_ID) == "" {
		return errors.New(fmt.Sprintf("env var for dealer ID %s is not set", conf.MSG_DEALER_ID))
	}

	return nil
}

/*
we cant share socket across goroutines
all send/recv biz rules has to be in
one goroutine
*/
func (d *ZDealer) Run() {
	log.Println("in the run...")
	for {
		log.Println("in the for loop")
		select {
		case msg := <-d.In:
			fmt.Print("***** dealer IN channel: send message to router", msg)
			d.sock.SendBytes([]byte(d.ID), zmq4.SNDMORE)
			d.sock.SendBytes([]byte(msg), 0)
			log.Println("dealer IN channel send to router - complete")
		case msg := <-d.RecvMsg():
			log.Println("||||| recived message from router ", string(msg))
		}

	}
}

func (d *ZDealer) RecvMsg() <-chan []byte {
	msg, err := d.sock.RecvMessageBytes(zmq4.DONTWAIT)
	if err != nil {
		log.Println(" receive error ", msg, err)
		return d.Out
	}

	log.Println("RecvMsg function pre", msg, err)
	go func() {
		d.Out <- msg[0]
	}()
	log.Println("post RecvMsg function")
	return d.Out
}

func (d *ZDealer) Listen(connstr string) error {
	conn := fmt.Sprintf("tcp://%s", connstr)
	err := d.sock.Bind(conn)
	if err != nil {
		return err
	}

	return nil
}

// host/port is that of router
func (d *ZDealer) Connect(connstr string) error {
	conn := fmt.Sprintf("tcp://%s", connstr)
	log.Println("Dealer connecting to router ", conn)
	err := d.sock.Connect(conn)
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

	in := make(chan []byte)
	out := make(chan []byte)
	return &ZDealer{ID: ID, In: in, Out: out, sock: dealer}, nil
}
