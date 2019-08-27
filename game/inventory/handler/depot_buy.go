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
	processor.Register(codec.MessageType(uipb.MessageType_CS_DEPOT_BUY_SLOTS_TYPE), dispatch.HandlerFunc(handleDepotBuySlot))
}

//处理仓库位置购买
func handleDepotBuySlot(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理仓库位置购买")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csDepotBuySlots := msg.(*uipb.CSDepotBuySlots)
	buyNum := csDepotBuySlots.GetBuyNum()

	err = depotBuySlots(tpl, buyNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"buyNum":   buyNum,
				"error":    err,
			}).Error("inventory:处理仓库位置购买,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"buyNum":   buyNum,
		}).Debug("inventory:处理仓库位置购买,完成")
	return nil
}

//购买槽位
func depotBuySlots(pl player.Player, buyNum int32) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	singleSlotNeedGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDepotUnlockSlotNeedGold)
	totalNeedGold := int64(singleSlotNeedGold * buyNum)
	flag := propertyManager.HasEnoughGold(int64(totalNeedGold), true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"buyNum":        buyNum,
				"totalNeedGold": totalNeedGold,
			}).Warn("inventory:处理获取仓库购买槽位,元宝不足")

		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	goldUseReason := commonlog.GoldLogReasonBuyDepotSlots
	flag = propertyManager.CostGold(totalNeedGold, true, goldUseReason, goldUseReason.String())
	if !flag {
		panic(fmt.Errorf("inventory:花费元宝应该成功"))
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag = inventoryManager.IfCanAddDepotSlots(buyNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"buyNum":   buyNum,
			}).Warn("inventory:处理仓库购买槽位,不能购买槽位了")
			
		playerlogic.SendSystemMessage(pl, lang.InventoryCanNotAddSlot)
		return
	}
	flag = inventoryManager.AddDepotSlots(buyNum)
	if !flag {
		panic(fmt.Errorf("inventory:add depot slots should be ok"))
	}

	propertylogic.SnapChangedProperty(pl)

	slotsNum := inventoryManager.GetDepotSlots()
	scDepotBuySlots := pbutil.BuildSCDepotBuySlots(slotsNum)
	pl.SendMsg(scDepotBuySlots)
	return
}
