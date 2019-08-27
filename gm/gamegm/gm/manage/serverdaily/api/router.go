package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	serverDailyPath = "/serverdaily"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(serverDailyPath).Subrouter()

	//以下需要加入权限
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerDaily, http.HandlerFunc(handleServerDailyReport)))
}
