package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	serverPath = "/server"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(serverPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleAddCenterServer)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServer)))

	sr.Path("/updateparent").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerParent)))
	sr.Path("/updateparentarray").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerParentArray)))
	sr.Path("/updatejiaoyizhanqu").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerJiaoYiZhanQu)))
	sr.Path("/updatejiaoyizhanquarray").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerJiaoYiZhanQuArray)))
	sr.Path("/updatepingtaifu").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerPingTaiFu)))
	sr.Path("/updatepingtaifuarray").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerPingTaiFuArray)))
	sr.Path("/updatechengzhanfu").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerChengZhan)))
	sr.Path("/updatechengzhanfuarray").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleUpdateCenterServerChengZhanArray)))

	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleDeleteCenterServer)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleCenterServerList)))
	sr.Path("/serverlisttype").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleCenterTypeServerList)))
	sr.Path("/refresh").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handleCenterRefresh)))

	sr.Path("/simplelist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerSimpleList, http.HandlerFunc(handleSimpleCenterServerList)))
	sr.Path("/simpleupdate").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerSimpleList, http.HandlerFunc(handleUpdateCenterServerName)))
	sr.Path("/zhanqulist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerSimpleList, http.HandlerFunc(handleCenterServerZhanQuList)))
	sr.Path("/zhanqulistexport").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerSimpleList, http.HandlerFunc(handleCenterServerZhanQuListExport)))

	sr.Path("/ping").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerManage, http.HandlerFunc(handlePingCenterServer)))
}
