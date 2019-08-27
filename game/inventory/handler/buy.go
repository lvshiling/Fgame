package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_BUY_SLOTS_TYPE), dispatch.HandlerFunc(handleInventoryBuySlot))
}

//处理购买槽位
func handleInventoryBuySlot(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理获取背包购买槽位")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryBuySlots := msg.(*uipb.CSInventoryBuySlots)
	buyNum := csInventoryBuySlots.GetBuyNum()

	err = buySlots(tpl, buyNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"buyNum":   buyNum,
				"error":    err,
			}).Error("inventory:处理获取背包购买槽位,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"buyNum":   buyNum,
		}).Debug("inventory:处理获取背包购买槽位,完成")
	return nil
}

//购买槽位
func buySlots(pl player.Player, buyNum int32) (err error) {
	singleSlotGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGoldForSingleSlot)
	openGold := int64(singleSlotGold * buyNum)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(int64(openGold), true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"buyNum":        buyNum,
				"totalNeedGold": openGold,
			}).Warn("inventory:处理获取背包购买槽位,元宝不足")

		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	reasonText := commonlog.GoldLogReasonBuySlots.String()
	flag = propertyManager.CostGold(openGold, true, commonlog.GoldLogReasonBuySlots, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:花费元宝应该成功"))
	}
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag = manager.IfCanAddSlots(buyNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:处理获取背包购买槽位,不能购买槽位了")

		playerlogic.SendSystemMessage(pl, lang.InventoryCanNotAddSlot)
		return
	}
	flag = manager.AddSlots(buyNum)
	if !flag {
		panic(fmt.Errorf("inventory:add slots should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	slotsNum := manager.GetSlots()
	inventoryBuySlots := pbutil.BuildSCInventoryBuySlots(slotsNum)
	pl.SendMsg(inventoryBuySlots)

	return nil
}
