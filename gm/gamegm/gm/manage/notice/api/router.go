package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	noticePath = "/notice"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(noticePath).Subrouter()

	//以下需要加入权限
	sr.Path("/notice").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeNotice, http.HandlerFunc(handleNotice)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeNotice, http.HandlerFunc(handleNoticeList)))
}
