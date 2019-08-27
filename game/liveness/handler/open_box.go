package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playerinventory "fgame/fgame/game/inventory/player"
	livenesslogic "fgame/fgame/game/liveness/logic"
	"fgame/fgame/game/liveness/pbutil"
	playerliveness "fgame/fgame/game/liveness/player"
	livenesstemplate "fgame/fgame/game/liveness/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LIVENESS_OPEN_TYPE), dispatch.HandlerFunc(handleLivenessOpen))
}

//处理活跃度奖励信息
func handleLivenessOpen(s session.Session, msg interface{}) (err error) {
	log.Debug("liveness:处理活跃度奖励消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLivenessOpen := msg.(*uipb.CSLivenessOpen)
	openBox := csLivenessOpen.GetBoxId()

	err = livenessOpen(tpl, openBox)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"openBox":  openBox,
				"error":    err,
			}).Error("liveness:处理活跃度奖励消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("liveness:处理活跃度奖励消息完成")
	return nil
}

//活跃度奖励信息的逻辑
func livenessOpen(pl player.Player, openBox int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLivenessDataManagerType).(*playerliveness.PlayerLivenessDataManager)
	flag := manager.IfLivenessBoxRew(openBox)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"openBox":  openBox,
			}).Warn("liveness:条件不足或已领取玩")
		playerlogic.SendSystemMessage(pl, lang.LivenessOpenBoxNoEnough)
		return
	}

	starTemplate := livenesstemplate.GetHuoYueTempalteService().GetHuoYueBoxTemplate(openBox)
	if starTemplate == nil {
		return
	}

	//判断背包是否足够
	rewItemMap := starTemplate.GetRewItemMap()
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(rewItemMap) != 0 {
		flag = inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("liveness:背包空间不足,清理后再来")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	//宝箱奖励
	isReturn := livenesslogic.GiveLivenessBoxReward(pl, starTemplate)
	if isReturn {
		return
	}
	flag = manager.LivenessOpenRew(openBox)
	if !flag {
		panic(fmt.Errorf("liveness: livenessOpen LivenessOpenRew should be ok"))
	}
	liveness := manager.GetLiveness()
	scLivenessOpen := pbutil.BuildSCLivenessOpen(liveness, openBox)
	pl.SendMsg(scLivenessOpen)
	return
}
