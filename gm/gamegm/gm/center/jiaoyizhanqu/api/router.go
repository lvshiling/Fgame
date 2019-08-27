package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	serverPath = "/zhanqu"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(serverPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeJiaoYiZhanQu, http.HandlerFunc(handleJiaoYiZhanQuAdd)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeJiaoYiZhanQu, http.HandlerFunc(handleJiaoYiZhanQuUpdate)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeJiaoYiZhanQu, http.HandlerFunc(handleJiaoYiZhanQuDelete)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeJiaoYiZhanQu, http.HandlerFunc(handleJiaoYiZhanQuList)))
	sr.Path("/alllist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeJiaoYiZhanQu, http.HandlerFunc(handleJiaoYiZhanQuListAll)))
}
