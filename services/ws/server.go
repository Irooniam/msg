package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Irooniam/msg/internal/socks"
	"github.com/coder/websocket"
	"golang.org/x/net/context"
)

type WS struct {
	Server *http.Server
	DIn    chan [][]byte
	DOut   chan [][]byte
	RIn    chan [][]byte
	ROut   chan [][]byte
}

func (ws *WS) handle(w http.ResponseWriter, r *http.Request) {
	rt := context.WithValue(r.Context(), "websocket-key", r.Header.Get("Sec-WebSocket-Key"))

	c, err := websocket.Accept(w, r.WithContext(rt), &websocket.AcceptOptions{
		Subprotocols:       []string{"echo"},
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println("error in handling websocket ", err)
		log.Println(err)
		return
	}

	//before doing anything - check the token from query string
	token := r.URL.Query().Get("token")
	if token == "" {
		log.Println("token query string empty")
		c.Close(websocket.StatusPolicyViolation, "missing token")
		return
	}
}

func NewWS(dealer *socks.ZDealer, router *socks.ZRouter) (*WS, error) {
	host := "127.0.0.1"
	port := 9080
	connuri := fmt.Sprintf("%s:%d", host, port)

	wServer := &WS{}
	wServer.RIn = router.In
	wServer.ROut = router.Out
	wServer.DIn = dealer.In
	wServer.DOut = dealer.Out

	server := &http.Server{}
	server.Handler = http.HandlerFunc(wServer.handle)
	server.Addr = connuri

	return wServer, nil
}
