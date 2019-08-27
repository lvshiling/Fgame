package api

import (
	"net/http"

	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"

	"github.com/gorilla/mux"
)

const (
	userPath = "/user"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()
	sr.Path("/login").Handler(http.HandlerFunc(handleLogin))
	sr.Path("/logout").Handler(http.HandlerFunc(handleLoginOut))
	sr.Path("/get_info").Handler(http.HandlerFunc(handleInfo))

	//以下需要加入权限
	sr.Path("/get_list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeUserManage, http.HandlerFunc(handleUserList)))
	sr.Path("/saveuser").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeUserManage, http.HandlerFunc(handleSaveUser)))
	sr.Path("/deleteuser").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeUserManage, http.HandlerFunc(handleDeleteUser)))
	sr.Path("/changepwd").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeUserManage, http.HandlerFunc(handleChangePassWord)))
	sr.Path("/childprivilege").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeUserManage, http.HandlerFunc(handleChildPrivilege)))
}
