package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_LIST_TYPE), dispatch.HandlerFunc(handleAllianceList))
}

//处理仙盟列表
func handleAllianceList(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = allianceList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理仙盟列表,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟列表,完成")
	return nil

}

//仙盟列表
func allianceList(pl player.Player) (err error) {
	alList := alliance.GetAllianceService().GetAllianceList()
	scAllianceList := pbutil.BuildSCAllianceList(alList)
	pl.SendMsg(scAllianceList)
	return
}
