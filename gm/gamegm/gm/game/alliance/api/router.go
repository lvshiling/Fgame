package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	alliancePath = "/alliance"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(alliancePath).Subrouter()

	//以下需要加入权限
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeAlliance, http.HandlerFunc(handleAllianceList)))
	sr.Path("/modifygonggao").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeAlliance, http.HandlerFunc(handleAllianceGongGao)))
	sr.Path("/dismiss").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeAlliance, http.HandlerFunc(handleAllianceDismiss)))

	sr.Path("/serverstate").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleServerOpenList)))
	sr.Path("/serverset").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleRegisterServerSet)))
	sr.Path("/serversetloglist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeChatSet, http.HandlerFunc(handleRegisterServerLog)))

}
