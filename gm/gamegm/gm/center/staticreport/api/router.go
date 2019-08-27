package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	userPath = "/report"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(userPath).Subrouter()

	//以下需要加入权限
	sr.Path("/online").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeOnLineReport, http.HandlerFunc(handleOnLineStatic)))
	sr.Path("/ngonline").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeOnLineReport, http.HandlerFunc(handleNeiGuaOnLineStatic)))
	sr.Path("/onlinetotal").Handler(http.HandlerFunc(handleLastOnLineStatic))
	sr.Path("/goldchange").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGoldChange, http.HandlerFunc(handleGoldChangeStatic)))
	sr.Path("/goldchangetype").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGoldChange, http.HandlerFunc(handleGoldChangeType)))
	sr.Path("/newbindgold").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGoldChange, http.HandlerFunc(handleNewBindGoldStatic)))
	sr.Path("/newgoldchange").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeGoldChange, http.HandlerFunc(handleNewGoldChangeStatic)))
	sr.Path("/retention").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRetention, http.HandlerFunc(handleRetentionStatic)))
	sr.Path("/tradeitem").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeCenterServerTradeItem, http.HandlerFunc(handleTradeItemList)))
	sr.Path("/recycle").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeOnLineReport, http.HandlerFunc(handleRecycleStatic)))
}
