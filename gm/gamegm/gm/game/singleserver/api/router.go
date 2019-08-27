package api

import (
	middleware "fgame/fgame/gm/gamegm/gm/middleware"
	types "fgame/fgame/gm/gamegm/gm/types"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	singleserverPath = "/singleserver"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(singleserverPath).Subrouter()

	sr.Path("/doublecharge").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeDoubleCharge, http.HandlerFunc(handleDoubleCharge)))
	sr.Path("/doublechargereset").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeDoubleCharge, http.HandlerFunc(handleDoubleChargeReset)))
	sr.Path("/doublechargeloglist").Handler(middleware.PrivilegeHandler(types.PrivilegeTypeDoubleCharge, http.HandlerFunc(handleDoubleChargeLogList)))

}
