package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	mybosslogic "fgame/fgame/game/myboss/logic"
	"fgame/fgame/game/myboss/pbutil"
	mybosstemplate "fgame/fgame/game/myboss/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"

	mybossplayer "fgame/fgame/game/myboss/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MY_BOSS_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerMyBossChallenge))
}

//个人boss挑战
func handlerMyBossChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("myboss:处理个人boss挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSMyBossChallenge)
	biologyId := csMsg.GetBiologyId()

	err = mybossChallenge(tpl, biologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"biologyId": biologyId,
				"err":       err,
			}).Error("myboss:处理个人boss挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  tpl.GetId(),
			"biologyId": biologyId,
		}).Debug("myboss：处理个人boss挑战请求完成")

	return
}

//个人boss挑战逻辑
func mybossChallenge(pl player.Player, biologyId int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	mybossTemp := mybosstemplate.GetMyBossTemplateService().GetMyBossTemplate(biologyId)
	if mybossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("myboss:个人boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	mybossManager := pl.GetPlayerDataManager(types.PlayerMyBossDataManagerType).(*mybossplayer.PlayerMyBossDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	needVipLevel := mybossTemp.NeedVipLevel
	curVipLevel := pl.GetVip()
	if curVipLevel < needVipLevel {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"needVipLevel": needVipLevel,
				"curVipLevel":  curVipLevel,
			}).Warn("myboss:个人boss挑战请求，vip等级不足")
		levelStr := fmt.Sprintf("%s", needVipLevel)
		playerlogic.SendSystemMessage(pl, lang.MyBossBaseVipNotEnough, levelStr)
		return
	}

	mybossManager.RefreshTimes()

	//是否免费次数
	hadAttendTimes := mybossManager.GetAttendTimes(biologyId)
	freeTimes := mybossTemp.FreeTimes
	if hadAttendTimes >= freeTimes {
		//挑战次数是否足够
		limitTimes := mybossTemp.TimesCount
		if hadAttendTimes >= limitTimes {
			log.WithFields(
				log.Fields{
					"playerId":       pl.GetId(),
					"biologyId":      biologyId,
					"hadAttendTimes": hadAttendTimes,
				}).Warn("myboss:个人boss挑战请求，次数不足，无法挑战")
			playerlogic.SendSystemMessage(pl, lang.MyBossBaseTimesNotEnough)
			return
		}

		attendNeedItemId := mybossTemp.UseItem
		attendNeedItemNum := mybossTemp.UseCount

		//挑战所需物品是否足够
		if !inventoryManager.HasEnoughItem(attendNeedItemId, attendNeedItemNum) {
			log.WithFields(
				log.Fields{
					"playerId":          pl.GetId(),
					"attendNeedItemId":  attendNeedItemId,
					"attendNeedItemNum": attendNeedItemNum,
				}).Warn("myboss:个人boss挑战请求，道具不足，无法挑战")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//消耗挑战所需物品
		itemReason := commonlog.InventoryLogReasonMyBossUse
		if flag := inventoryManager.UseItem(attendNeedItemId, attendNeedItemNum, itemReason, itemReason.String()); !flag {
			panic(fmt.Errorf("myboss: mybossChallenge use item should be ok"))
		}
	}

	//更新玩家挑战记录
	mybossManager.AddAttendTimes(biologyId)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	//进入场景
	flag := mybosslogic.PlayerEnterBoss(pl, biologyId)
	if !flag {
		panic("enter myboss scene should be ok!")
	}

	scMsg := pbutil.BuildSCMyBossChallenge()
	pl.SendMsg(scMsg)
	return
}
