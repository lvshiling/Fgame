package pool

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	poolPath = "/pool"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(poolPath).Subrouter()
	sr.Path("/add").Handler(http.HandlerFunc(handleAddPoolPlayer))
	sr.Path("/adduser").Handler(http.HandlerFunc(handleAddSupportPlayer))
}
