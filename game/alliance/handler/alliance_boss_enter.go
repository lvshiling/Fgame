package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_BOSS_ENTER_TYPE), dispatch.HandlerFunc(handleAllianceBossEnter))
}

//处理进入仙盟boss
func handleAllianceBossEnter(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理进入仙盟boss信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceBossEnter(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理进入仙盟boss信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理进入仙盟boss信息,完成")
	return nil

}

//处理进入仙盟boss
func allianceBossEnter(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	flag := playerlogic.CheckCanEnterScene(pl)
	if !flag {
		return
	}
	if !s.MapTemplate().IsWorld() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:您当前不在世界地图,无法召唤仙盟boss")
		return
	}
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:用户不在仙盟")
		playerlogic.SendSystemMessage(pl, lang.AllianceUserNotInAlliance)
		return
	}

	err = alliance.GetAllianceService().AllianceBossEnter(pl)
	if err != nil {
		return
	}
	return
}
