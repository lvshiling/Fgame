package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	serverPath = "/set"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(serverPath).Subrouter()
	//以下需要加入权限
	sr.Path("/clientverionset").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerMetaSet, http.HandlerFunc(handleClientVersionSet)))
	sr.Path("/clientverionget").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerMetaSet, http.HandlerFunc(handleClientVersionGet)))
	sr.Path("/platformserverconfigget").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerMetaSet, http.HandlerFunc(handleplatformServerConfigGet)))
	sr.Path("/platformserverconfigset").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerMetaSet, http.HandlerFunc(handlePlatformServerConfigSet)))
}
