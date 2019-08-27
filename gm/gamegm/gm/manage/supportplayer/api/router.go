package api

import (
	"fgame/fgame/gm/gamegm/gm/middleware"
	"fgame/fgame/gm/gamegm/gm/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	userPath = "/manage"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()

	//以下需要加入权限
	sr.Path("/supportplayerlist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPlayer, http.HandlerFunc(handlePlayerList)))
	sr.Path("/privilegecharge").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPlayer, http.HandlerFunc(handlePrivilegeCharge)))
	sr.Path("/privilegechargemulity").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPlayer, http.HandlerFunc(handlePrivilegeChargeMulity)))
	sr.Path("/privilegeset").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPlayer, http.HandlerFunc(handlePrivilegeSet)))
	sr.Path("/supportplayerlog").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeServerSupportPlayer, http.HandlerFunc(handleSupportPlayerLogList)))
	// sr.Path("/supportplayerlist").Handler(http.HandlerFunc(handlePlayerList))

}

//基础方法全部先写到路由这边

func changeInt64ToString(p_id int64) string {
	return strconv.FormatInt(p_id, 10)
}

func changeStringToInt64(p_id string) int64 {
	rst, _ := strconv.ParseInt(p_id, 10, 64)
	return rst
}
