package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	loginPath = "/login"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(loginPath).Subrouter()
	sr.Path("/visitor").HandlerFunc(http.HandlerFunc(handleVisitorLogin))
	sr.Path("/verify").HandlerFunc(http.HandlerFunc(handleVerify))
	sr.Path("/register").HandlerFunc(http.HandlerFunc(handleRegister))
	sr.Path("/login").HandlerFunc(http.HandlerFunc(handleLogin))
}

type LoginResponse struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
}

type RestResult struct {
	ErrorCode int         `json:"errorCode"`
	Result    interface{} `json:"result"`
}
