package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	userPath = "/channel"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChannelManage, http.HandlerFunc(handleAddChannel)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChannelManage, http.HandlerFunc(handleUpdateChannel)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChannelManage, http.HandlerFunc(handleDeleteChannel)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChannelManage, http.HandlerFunc(handleChannelList)))
	sr.Path("/all").Handler(http.HandlerFunc(handleAllChannelList))
}
