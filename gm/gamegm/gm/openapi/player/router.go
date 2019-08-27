package player

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	poolPath = "/player"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(poolPath).Subrouter()
	sr.Path("/create").Handler(http.HandlerFunc(handlePlayerAddRole))
}
