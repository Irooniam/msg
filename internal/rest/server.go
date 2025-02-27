package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Irooniam/msg/internal/socks"
	"github.com/julienschmidt/httprouter"
)

type RESTServer struct {
	Mux     *httprouter.Router
	Server  *http.Server
	ZDealer *socks.ZDealer
}

func NewREST(host string, port int, sock *socks.ZDealer) {
	addr := fmt.Sprintf("%s:%d", host, port)
	r := RESTServer{}
	r.Server = &http.Server{
		Addr: addr,
	}

	r.Mux = httprouter.New()
	r.Server.Handler = r.Mux
}

func Foo() {
	x := os.Getenv("MSG_DIR_HOST")
	log.Println("ms dir ", x)
}
