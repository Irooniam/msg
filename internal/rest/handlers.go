package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DefaultH struct {
}

func (h *DefaultH) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("OK"))
}
