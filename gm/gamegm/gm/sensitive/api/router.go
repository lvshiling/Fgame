package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	sensitivePath = "/sensitive"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(sensitivePath).Subrouter()

	// //以下需要加入权限

	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleAddSensitive)))
	sr.Path("/get").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatMinitor, http.HandlerFunc(handleGetSensitive)))
	// sr.Path("/list").Handler(http.HandlerFunc(handlePlatformList))
	// sr.Path("/all").Handler(http.HandlerFunc(handleAllPlatformList))
}
