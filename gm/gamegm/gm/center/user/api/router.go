package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	centerUserPath = "/centeruser"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(centerUserPath).Subrouter()

	//以下需要加入权限
	sr.Path("/querylist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterUserQuery, http.HandlerFunc(handleCenterUserList)))
	sr.Path("/neigualist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterUserQuery, http.HandlerFunc(handleCenterNeiGuaUserList)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterUser, http.HandlerFunc(handleCenterUserList)))
	sr.Path("/updategm").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterUser, http.HandlerFunc(handleUpdateGmList)))

	sr.Path("/updateusername").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handleUpdateUserNameList)))
	sr.Path("/userinfo").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handleUserInfo)))
	sr.Path("/updateforbid").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handleUpdateForbid)))
	sr.Path("/updateipforbid").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handleIpUpdateForbid)))
	sr.Path("/getipstate").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handleIpGetState)))
	sr.Path("/updateipunforbid").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlayerSearch, http.HandlerFunc(handleIpUpdateUnForbid)))
}
