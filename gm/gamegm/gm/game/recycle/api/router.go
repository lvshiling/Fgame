package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	userPath = "/recycle"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()

	//以下需要加入权限
	sr.Path("/recyclegold").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRecycle, http.HandlerFunc(handlerecycleGold)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRecycle, http.HandlerFunc(handleRecycleList)))
}
