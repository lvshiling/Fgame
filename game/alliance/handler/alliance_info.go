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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BRIEF_INFO_TYPE), dispatch.HandlerFunc(handleAllianceBriefInfo))
}

//仙盟信息
func handleAllianceBriefInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:获取仙盟信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = allianceInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:获取仙盟信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:获取仙盟信息,完成")
	return nil

}

func allianceInfo(pl player.Player) (err error) {
	member := alliance.GetAllianceService().GetAllianceMember(pl.GetId())
	if member == nil {
		return
	}

	// TODO xzk27 仙盟增加阵营字段

	scAlliance := pbutil.BuildSCAlliance(member.GetAlliance())
	pl.SendMsg(scAlliance)
	return
}
