package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Irooniam/msg/conf"
	"github.com/Irooniam/msg/internal/socks"
	"github.com/julienschmidt/httprouter"
)

type RESTServer struct {
	Mux     *httprouter.Router
	Server  *http.Server
	ZDealer *socks.ZDealer
}

func ChkiHTTPConf() error {
	if os.Getenv(conf.MSG_IHTTP_HOST) == "" {
		return errors.New(fmt.Sprintf("env var for ihttp host %s is not set", conf.MSG_IHTTP_HOST))
	}
	if os.Getenv(conf.MSG_IHTTP_PORT) == "" {
		return errors.New(fmt.Sprintf("env var for ihttp port %s is not set", conf.MSG_IHTTP_PORT))
	}

	return nil
}

func (r *RESTServer) SetupHandlers() {
	h := DefaultH{}
	r.Mux.GET("/", h.Index)
}

// lstr = host:port format
func NewREST(lstr string, sock *socks.ZDealer) *RESTServer {
	r := RESTServer{}
	r.Server = &http.Server{
		Addr: lstr,
	}

	r.Mux = httprouter.New()
	r.Server.Handler = r.Mux
	return &r
}

func Foo() {
	x := os.Getenv("MSG_DIR_HOST")
	log.Println("ms dir ", x)
}
