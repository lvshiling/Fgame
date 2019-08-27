package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	xuedunlogic "fgame/fgame/game/xuedun/logic"
	"fgame/fgame/game/xuedun/pbutil"
	playerxuedun "fgame/fgame/game/xuedun/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XUEDUN_UPGRADE_TYPE), dispatch.HandlerFunc(handleXueDunUpgrade))
}

//处理血盾升阶信息
func handleXueDunUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("xuedun:处理血盾升阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = xueDunUpgrade(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xuedun:处理血盾升阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xuedun:处理血盾升阶完成")
	return nil
}

//血盾升阶的逻辑
func xueDunUpgrade(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerXueDunDataManagerType).(*playerxuedun.PlayerXueDunDataManager)
	xueDunInfo := manager.GetXueDunInfo()
	beforeNumber := xueDunInfo.GetNumber()
	beforeStar := xueDunInfo.GetStar()
	curBlood := xueDunInfo.GetBlood()
	bloodShieldTemplate, flag := manager.IsFull()
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xuedun:血盾系统已达满阶")
		playerlogic.SendSystemMessage(pl, lang.XueDunUpgradeReachFull)
		return
	}

	needYinliang := int64(bloodShieldTemplate.UseMoney)
	useBlood := int64(bloodShieldTemplate.UseBlood)
	needItems := bloodShieldTemplate.GetUseItemTemplate()

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	flag = propertyManager.HasEnoughSilver(needYinliang)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xuedun:银两不足，无法升阶")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	if curBlood < useBlood {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xuedun:血炼值不足，无法升阶")
		playerlogic.SendSystemMessage(pl, lang.XueDunUpgradeNoBlood)
		return
	}

	//升阶需要物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xuedun:道具不足，无法升阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//消耗钱
	if needYinliang != 0 {
		reasonSliverText := commonlog.SilverLogReasonXueDunUpgrade.String()
		flag := propertyManager.CostSilver(needYinliang, commonlog.SilverLogReasonXueDunUpgrade, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("xuedun: xueDunUpgrade Cost should be ok"))
		}
	}

	//消耗物品
	if len(needItems) != 0 {
		inventoryReason := commonlog.InventoryLogReasonXueDunUpgrade
		reasonText := inventoryReason.String()
		flag := inventoryManager.BatchRemove(needItems, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("xuedun: xueDunUpgrade BatchRemove item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = manager.XueDunSubBloodChanged(useBlood)
	if !flag {
		panic(fmt.Errorf("xuedun: xueDunUpgrade XueDunSubBloodChanged should be ok"))
	}

	//血盾升阶判断
	pro, _, sucess := xuedunlogic.XueDunUpgrade(xueDunInfo.GetStarNum(), xueDunInfo.GetStarPro(), bloodShieldTemplate)
	flag = manager.Upgrade(bloodShieldTemplate, pro, sucess)
	if !flag {
		panic(fmt.Errorf("xuedun: xueDunUpgrade should be ok"))
	}
	if sucess {
		xueDunReason := commonlog.XueDunLogReasonUpgrade
		reasonText := xueDunReason.String()
		evetData := xueduneventtypes.CreatePlayerXueDunUpgradeLogEventData(beforeNumber, beforeStar, xueDunInfo.GetNumber(), xueDunInfo.GetStar(), xueDunReason, reasonText)
		gameevent.Emit(xueduneventtypes.EventTypeXueDunUpgradeLog, pl, evetData)
		xuedunlogic.XueDunPropertyChanged(pl)
	}
	scXueDunUpgrade := pbutil.BuildSCXueDunUpgrade(sucess, xueDunInfo)
	pl.SendMsg(scXueDunUpgrade)
	return
}
