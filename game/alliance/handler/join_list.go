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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_JOIN_APPLY_LIST_TYPE), dispatch.HandlerFunc(handleAllianceJoinApplyList))
}

//处理仙盟加入列表
func handleAllianceJoinApplyList(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟加入列表")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceJoinApllyList(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("alliance:处理仙盟加入列表,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟加入列表,完成")
	return nil

}

//仙盟加入
func allianceJoinApllyList(pl player.Player) (err error) {
	//申请加入仙盟
	applyObjList, err := alliance.GetAllianceService().GetJoinApplyList(pl.GetId())
	if err != nil {
		return
	}
	scAllianceJoinApplyList := pbutil.BuildSCAllianceJoinApplyList(applyObjList)
	pl.SendMsg(scAllianceJoinApplyList)
	return
}
