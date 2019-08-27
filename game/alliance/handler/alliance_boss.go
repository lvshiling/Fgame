package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BOSS_TYPE), dispatch.HandlerFunc(handleAllianceBoss))
}

//处理仙盟boss
func handleAllianceBoss(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟boss信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceBoss(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理仙盟boss信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟boss信息,完成")
	return nil

}

//处理仙盟boss
func allianceBoss(pl player.Player) (err error) {
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:用户不在仙盟内")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	status, level, exp, summonTime, err := alliance.GetAllianceService().AllianceBossInfo(allianceId)
	if err != nil {
		return
	}

	scAllianceBoss := pbutil.BuildSCAllianceBoss(int32(status), level, exp, summonTime)
	pl.SendMsg(scAllianceBoss)
	return
}
