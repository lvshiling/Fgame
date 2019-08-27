package api

import (
	"net/http"

	chatsetrouter "fgame/fgame/gm/gamegm/gm/center/chatset/api"
	jiaoyizhanqu "fgame/fgame/gm/gamegm/gm/center/jiaoyizhanqu/api"
	centernotice "fgame/fgame/gm/gamegm/gm/center/notice/api"
	orderrouter "fgame/fgame/gm/gamegm/gm/center/order/api"
	centerrouter "fgame/fgame/gm/gamegm/gm/center/platform/api"
	redeemrouter "fgame/fgame/gm/gamegm/gm/center/redeem/api"
	serverrouter "fgame/fgame/gm/gamegm/gm/center/server/api"
	setrouter "fgame/fgame/gm/gamegm/gm/center/set/api"
	centeruser "fgame/fgame/gm/gamegm/gm/center/user/api"

	"github.com/gorilla/mux"
)

const (
	centerPath = "/center"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(centerPath).Subrouter()

	sr.Path("/platform").Handler(http.HandlerFunc(handleAllCenterPlatform))
	sr.Path("/group").Handler(http.HandlerFunc(handleGroup))
	sr.Path("/server").Handler(http.HandlerFunc(handleServer))
	sr.Path("/sdktype").Handler(http.HandlerFunc(handleSdkType))
	sr.Path("/allserver").Handler(http.HandlerFunc(handleAllServerById))
	sr.Path("/alluserserver").Handler(http.HandlerFunc(handleAllCenterServer))
	centerrouter.Router(sr)
	serverrouter.Router(sr)
	chatsetrouter.Router(sr)
	orderrouter.Router(sr)
	redeemrouter.Router(sr)
	centeruser.Router(sr)
	centernotice.Router(sr)
	setrouter.Router(sr)
	jiaoyizhanqu.Router(sr)
}
