package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	exchangePath = "/exchange"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(exchangePath).Subrouter()
	sr.Path("/code").HandlerFunc(http.HandlerFunc(handleExchangeCode))
	sr.Path("/expire").HandlerFunc(http.HandlerFunc(handleExchangeExpire))
	sr.Path("/exchange").HandlerFunc(http.HandlerFunc(handleExchange))

}
