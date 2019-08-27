package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	platformPath = "/redeem"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(platformPath).Subrouter()

	//以下需要加入权限
	sr.Path("/add").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRedeem, http.HandlerFunc(handleRedeemAdd)))
	sr.Path("/list").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRedeem, http.HandlerFunc(handleRedeemList)))
	sr.Path("/delete").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRedeem, http.HandlerFunc(handleRedeemDelete)))
	sr.Path("/code").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRedeem, http.HandlerFunc(handleRedeemCode)))
	sr.Path("/codelist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRedeem, http.HandlerFunc(handleRedeemCodeList)))
	sr.Path("/codelistexport").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeRedeem, http.HandlerFunc(handleRedeemCodeListExport)))
}
