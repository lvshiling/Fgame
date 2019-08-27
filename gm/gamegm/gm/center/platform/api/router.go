package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	platformPath = "/platform"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(platformPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterPlatformManage, http.HandlerFunc(handleAddCenterPlatform)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterPlatformManage, http.HandlerFunc(handleUpdateCenterPlatform)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterPlatformManage, http.HandlerFunc(handleDeleteCenterPlatform)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterPlatformManage, http.HandlerFunc(handleCenterPlatformList)))

	sr.Path("/marrylist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleCenterPlatformMarrySetList)))
	sr.Path("/marryflag").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleCenterPlatformMarrySetFlag)))
	sr.Path("/marrycontent").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleCenterPlatformMarrySetContent)))
	sr.Path("/marryloglist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleMarrySetLogList)))
	sr.Path("/marrylogsend").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleMarrySetLogSend)))

	sr.Path("/setting").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleCenterPlatformSetting)))
	sr.Path("/savesetting").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleCenterPlatformSettingSave)))
	sr.Path("/settinglist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypePlatformManageMarry, http.HandlerFunc(handleCenterPlatformSettingList)))
}
