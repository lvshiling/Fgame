package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	couponPath = "/coupon"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(couponPath).Subrouter()
	sr.Path("/exchange").HandlerFunc(http.HandlerFunc(handleExchangeCoupon))
}
