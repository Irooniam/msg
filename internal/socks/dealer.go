package socks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Irooniam/msg/conf"
	zmq "github.com/pebbe/zmq4/draft"
)

type ZDealer struct {
	ID   string
	In   chan [][]byte
	Out  chan [][]byte
	Done chan bool
	sock *zmq.Socket
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
	log.Println("Starting loop to receive/send on dealer socket...")
	for {
		select {
		case out := <-d.Out:
			log.Printf("sending message out to router payload: %v", out)
			d.sendMsg(out[0], out[1])
		case <-d.RecvMsg():
			log.Println("received message on dealer socket.")
		case <-d.Done:
			log.Println("looks like we are done...")
			return
		case <-time.After(time.Millisecond * 10):
			continue
		}

	}
}

func (d *ZDealer) RecvMsg() <-chan [][]byte {
	msg, err := d.sock.RecvMessageBytes(zmq.DONTWAIT)
	if err != nil {
		if err.Error() != "resource temporarily unavailable" {
			log.Printf("error on router recvmsg function '%s' - '%s'", msg, err)
		}
		return d.In
	}

	d.In <- [][]byte{msg[0], msg[1]}
	return d.In
}

/*
*

	Responsible for parsing all incoming messages from
	dealer socket

*
*/
func (r *ZDealer) ParseIn() {
	for {
		msg := <-r.In
		log.Println("Parse incoming message ", msg)
	}
}

func (d *ZDealer) sendMsg(action []byte, msg []byte) {
	x, err := d.sock.SendBytes(action, zmq.SNDMORE)
	log.Println(x, err)

	x, err = d.sock.SendBytes(msg, 0)
	log.Println(x, err)
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
	dealer, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		return &ZDealer{}, err
	}

	err = dealer.SetIdentity(ID)
	if err != nil {
		return &ZDealer{}, errors.New(fmt.Sprintf("Tried setting identity but got error: %s", err))
	}

	in := make(chan [][]byte)
	out := make(chan [][]byte)
	done := make(chan bool)
	return &ZDealer{ID: ID, In: in, Out: out, Done: done, sock: dealer}, nil
}
