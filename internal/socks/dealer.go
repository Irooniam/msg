package socks

import (
	"errors"
	"fmt"
	"os"

	"github.com/Irooniam/msg/conf"
	"github.com/pebbe/zmq4"
)

type ZDealer struct {
	ID   string
	In   chan []byte
	Out  chan []byte
	Sock *zmq4.Socket
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

func (d *ZDealer) Listen(connstr string) error {
	conn := fmt.Sprintf("tcp://%s", connstr)
	err := d.Sock.Bind(conn)
	if err != nil {
		return err
	}

	return nil
}

// host/port is that of router
func (d *ZDealer) Connect(connstr string) error {
	conn := fmt.Sprintf("tcp://%s", connstr)
	err := d.Sock.Connect(conn)
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

	return &ZDealer{ID: ID, In: in, Out: out, Sock: dealer}, nil
}
