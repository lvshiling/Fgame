package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	serverPath = "/notice"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(serverPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleAddLoginNotice)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterLoginNotice)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleDeleteLoginNotice)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleLoginNoticeList)))
	sr.Path("/defaultadd").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateDefaultNotice)))
	sr.Path("/defaultinfo").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleDefaultNoticeInfo)))
	sr.Path("/refresh").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleCenterNoticeRefresh)))
}
