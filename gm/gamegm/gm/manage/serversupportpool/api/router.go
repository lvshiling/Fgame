package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	managePath = "/manage"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(managePath).Subrouter()

	//以下需要加入权限
	sr.Path("/addserversppool").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleAddSupportPool)))
	sr.Path("/updateserversppool").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleUpdateSupportPool)))
	sr.Path("/deleteserversppool").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleDeleteSupportPool)))
	sr.Path("/serversppoollist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleServerSupportPoolList)))

	sr.Path("/addplatformpool").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleAddPlatformSupportPoolSet)))
	sr.Path("/updateplatformpool").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleUpdatePlatformSupportPoolSet)))
	sr.Path("/deleteplatformpool").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handleDeletePlatformSupportPoolSet)))
	sr.Path("/platformpoollist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPool, http.HandlerFunc(handlePlatformSupportPoolSetList)))
}
