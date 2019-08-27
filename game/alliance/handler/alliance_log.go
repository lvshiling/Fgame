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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_LOG_TYPE), dispatch.HandlerFunc(handleAllianceLogList))
}

//请求仙盟日志列表
func handleAllianceLogList(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:请求仙盟日志列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = allianceLogList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:请求仙盟日志列表,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:请求仙盟日志列表,完成")
	return nil

}

func allianceLogList(pl player.Player) (err error) {
	allianceLogList, err := alliance.GetAllianceService().GetAllianceLogList(pl.GetId())
	if err != nil {
		return
	}
	scAllianceList := pbutil.BuildSCAllianceLogList(allianceLogList)
	pl.SendMsg(scAllianceList)
	return
}
