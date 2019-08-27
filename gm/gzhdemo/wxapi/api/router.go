package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	r.Path("/gzhapi").HandlerFunc(http.HandlerFunc(handlegzhapi))
}
