package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	userPath = "/platform"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManage, http.HandlerFunc(handleAddPlatform)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManage, http.HandlerFunc(handleUpdatePlatform)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManage, http.HandlerFunc(handleDeletePlatform)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManage, http.HandlerFunc(handlePlatformList)))
	sr.Path("/refreshsdk").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManage, http.HandlerFunc(handleRefreshSdk)))
	sr.Path("/all").Handler(http.HandlerFunc(handleAllPlatformList))
}
