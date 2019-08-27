package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	"fgame/fgame/game/center/center"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ALLIANCE_DEPOT_TYPE), dispatch.HandlerFunc(handleMergeDepot))
}

//处理仙盟仓库整理
func handleMergeDepot(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟仓库整理")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSAllianceDepotMerge)
	indexList := csMsg.GetIndexList()
	err = mergeAllianceDepot(tpl, indexList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": indexList,
				"error":     err,
			}).Error("alliance:处理仙盟仓库整理,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"indexList": indexList,
		}).Debug("alliance:处理仙盟仓库整理,完成")
	return nil

}

//仙盟仓库整理
func mergeAllianceDepot(pl player.Player, indexList []int32) (err error) {
	if !center.GetCenterService().IsAllianceOpen() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("alliance:仙盟仓库关闭中")
		playerlogic.SendSystemMessage(pl, lang.AllianceDepotClose)
		return
	}
	if coreutils.IfRepeatElementInt32(indexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": indexList,
			}).Warn("alliance:处理仙盟仓库整理,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//整理
	err = alliance.GetAllianceService().MergeDepot(pl, indexList)
	if err != nil {
		return
	}

	scMsg := pbutil.BuildSCAllianceDepotMerge()
	pl.SendMsg(scMsg)
	return
}
