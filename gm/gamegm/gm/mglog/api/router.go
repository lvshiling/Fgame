package api

import (
	"fgame/fgame/gm/gamegm/gm/middleware"
	"fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	logPath = "/log"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(logPath).Subrouter()

	//以下需要加入权限
	sr.Path("/get").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGameLog, http.HandlerFunc(handleGetMongoLog)))
	sr.Path("/meta").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGameLog, http.HandlerFunc(handleGetLogMeta)))
	sr.Path("/metamsglist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGameLog, http.HandlerFunc(handleGetLogMsgMetaList)))

	sr.Path("/getchatlog").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGameLog, http.HandlerFunc(handleGetMongoChatLog)))
	sr.Path("/playerstats").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGameLog, http.HandlerFunc(handlePlayerStatsList)))
}
