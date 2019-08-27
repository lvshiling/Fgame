package api

import (
	centerRouter "fgame/fgame/gm/gamegm/gm/center/api"
	reportrouter "fgame/fgame/gm/gamegm/gm/center/staticreport/api"
	channelRouter "fgame/fgame/gm/gamegm/gm/channel/api"
	feedbackfeeapi "fgame/fgame/gm/gamegm/gm/feedbackfee/api"
	allianceRouter "fgame/fgame/gm/gamegm/gm/game/alliance/api"
	playerRouter "fgame/fgame/gm/gamegm/gm/game/player/api"
	recycleapi "fgame/fgame/gm/gamegm/gm/game/recycle/api"
	singleserverapi "fgame/fgame/gm/gamegm/gm/game/singleserver/api"
	mailRouter "fgame/fgame/gm/gamegm/gm/manage/mail/api"
	noticerouter "fgame/fgame/gm/gamegm/gm/manage/notice/api"
	dailyrouter "fgame/fgame/gm/gamegm/gm/manage/serverdaily/api"
	sp "fgame/fgame/gm/gamegm/gm/manage/serversupportpool/api"
	supportplayer "fgame/fgame/gm/gamegm/gm/manage/supportplayer/api"
	logRouter "fgame/fgame/gm/gamegm/gm/mglog/api"
	platformRouter "fgame/fgame/gm/gamegm/gm/platform/api"
	sensitive "fgame/fgame/gm/gamegm/gm/sensitive/api"
	userRouter "fgame/fgame/gm/gamegm/gm/user/api"

	"github.com/gorilla/mux"
)

const (
	gmPath = "/gm"
)

func Router(r *mux.Router) {
	sr := r.PathPrefix(gmPath).Subrouter()
	userRouter.Router(sr)
	channelRouter.Router(sr)
	platformRouter.Router(sr)
	centerRouter.Router(sr)
	playerRouter.Router(sr)
	sensitive.Router(sr)
	logRouter.Router(sr)
	allianceRouter.Router(sr)
	mailRouter.Router(sr)
	supportplayer.Router(sr)
	sp.Router(sr)
	reportrouter.Router(sr)
	noticerouter.Router(sr)
	dailyrouter.Router(sr)
	recycleapi.Router(sr)
	feedbackfeeapi.Router(sr)
	singleserverapi.Router(sr)
}
