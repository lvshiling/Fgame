package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	serverPath = "/order"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(serverPath).Subrouter()

	//以下需要加入权限
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleOrderList)))
	sr.Path("/gamelist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleGameOrderList)))
	sr.Path("/gamestatic").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleGameOrderStatic)))
	sr.Path("/centerstatic").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleOrderStatic)))
	sr.Path("/centertotalstatic").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleorderStaticTotal)))
	sr.Path("/centerdatestatic").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleOrderDateStaticList)))
	sr.Path("/centerdateplatformstatic").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerCenterOrderList, http.HandlerFunc(handleOrderDatePlatformStaticList)))
}
