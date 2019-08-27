package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	gamecommon "fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_AGREE_JOIN_APPLY_BATCH_TYPE), dispatch.HandlerFunc(handleAllianceAgreeJoinApplyBatch))
}

//批量处理仙盟加入申请
func handleAllianceAgreeJoinApplyBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:批量处理仙盟加入申请")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAllianceAgreeJoinApplyBatch := msg.(*uipb.CSAllianceAgreeJoinApplyBatch)
	agree := csAllianceAgreeJoinApplyBatch.GetAgree()
	err = allianceAgreeJoinBatch(tpl, agree)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"agree":    agree,
				"error":    err,
			}).Error("alliance:批量处理仙盟加入申请,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"agree":    agree,
		}).Debug("alliance:批量处理仙盟加入申请,完成")
	return nil

}

func allianceAgreeJoinBatch(pl player.Player, agree bool) (err error) {
	joinList, err := alliance.GetAllianceService().GetJoinApplyList(pl.GetId())
	newJoinList := make([]*alliance.AllianceJoinApplyObject, len(joinList))
	copy(newJoinList, joinList)
	for _, joinObj := range newJoinList {
		err = agreeJoin(pl, joinObj.GetJoinId(), agree)
		if err != nil {
			terr, ok := err.(gamecommon.Error)
			if !ok {
				return
			}
			playerlogic.SendSystemMessage(pl, terr.Code())
		} 
	}

	scMsg := pbutil.BuildSCAllianceAgreeJoinApplyBatch()
	pl.SendMsg(scMsg)
	return
}
