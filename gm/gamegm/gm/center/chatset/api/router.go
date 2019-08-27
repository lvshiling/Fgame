package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	chatSetPath = "/chatset"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(chatSetPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleAddChatSet)))
	sr.Path("/update").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleUpdateChatSet)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleDeleteChatSet)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleChatSetList)))

	sr.Path("/addplatform").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleAddChatSetPlatform)))
	sr.Path("/updateplatform").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleUpdateChatSetPlatform)))
	sr.Path("/deleteplatform").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleDeleteChatSetPlatform)))
	sr.Path("/listplatform").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleChatSetListPlatform)))
}
