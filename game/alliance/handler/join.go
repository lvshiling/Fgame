package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_JOIN_APPLY_TYPE), dispatch.HandlerFunc(handleAllianceJoinApply))
}

//处理仙盟加入
func handleAllianceJoinApply(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟加入")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAllianceJoinApply := msg.(*uipb.CSAllianceJoinApply)
	allianceId := csAllianceJoinApply.GetAllianceId()

	err = allianceJoin(tpl, allianceId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"allianceId": allianceId,
				"error":      err,
			}).Error("alliance:处理仙盟加入,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"allianceId": allianceId,
		}).Debug("alliance:处理仙盟加入,完成")
	return nil

}

//仙盟加入
func allianceJoin(pl player.Player, allianceId int64) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeAlliance) {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"allianceId": allianceId,
			}).Warn("alliance:处理仙盟加入,仙盟功能未开放")

		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	_, err = alliancelogic.HandleApplyJoinAlliance(pl, allianceId)
	if err != nil {
		return
	}

	scAllianceJoinApply := pbutil.BuildSCAllianceJoinApply(allianceId)
	pl.SendMsg(scAllianceJoinApply)
	return
}
