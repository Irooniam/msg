package socks

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/states"
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
	if len(msg) == 2 && len(msg[1]) == 0 {
		r.In <- [][]byte{msg[0], []byte(states.ACTIONS["DEALER-EVENT"]), []byte{0x0}}
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
		select {
		case msg := <-r.In:
			//have to receive 3 parts: header, action, payload
			if len(msg) != 3 {
				log.Printf("expected msg received to be 3 parts but is: %d", len(msg))
				continue
			}

			action, err := states.TranslateAction(msg[1])
			if err != nil {
				log.Printf("tried getting action from msg but got %s", err)
				continue
			}

			//all messages go / recv in []byte
			saction := string(action)

			/*
				We are matching what the action message TRANSLATES to not actual msg (DR)
			*/
			switch saction {
			case "DEALER-EVENT": //connect/disconnect
				log.Println("dealer event ", msg)
				states.DealerEvent(msg[0], r.Out)

			default:
				log.Printf("actions is %s - and we dont have case math", action)

			}

		case <-r.PDone:
			log.Println("ParseIn goroutine is done")
			return
		}
	}
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
