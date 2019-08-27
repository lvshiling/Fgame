package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tulong/pbutil"
	tulong "fgame/fgame/game/tulong/tulong"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_RANK_TYPE), dispatch.HandlerFunc(handleTuLongRankGet))
}

//处理屠龙排行榜获得
func handleTuLongRankGet(s session.Session, msg interface{}) (err error) {
	log.Debug("tulong:处理屠龙排行榜获得")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = tuLongRankGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("tulong:处理屠龙排行榜获得,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("tulong:处理选择屠龙")
	return nil

}

//处理屠龙排行榜获得
func tuLongRankGet(pl player.Player) (err error) {
	serverId := pl.GetServerId()
	allianceId := pl.GetAllianceId()

	dataList, pos := tulong.GetTuLongService().GetRankList(serverId, allianceId)
	scRankList := pbutil.BuildSCRankList(dataList, pos)
	pl.SendMsg(scRankList)
	return
}
