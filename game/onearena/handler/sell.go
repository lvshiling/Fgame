package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/onearena/pbutil"
	playeronearena "fgame/fgame/game/onearena/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"math"

	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"

	commonlog "fgame/fgame/common/log"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_SELL_TYPE), dispatch.HandlerFunc(handleOneArenaSell))
}

//处理鲲一键卖出信息
func handleOneArenaSell(s session.Session, msg interface{}) (err error) {
	log.Debug("onearena:处理获取鲲一键卖出消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = oneArenaSell(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("onearena:处理获取鲲一键卖出消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("onearena:处理获取鲲一键卖出消息完成")
	return nil
}

//处理鲲一键卖出信息逻辑
func oneArenaSell(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	sellSilver := int64(0)
	tempSellSilver := float64(0)
	sellGold := int64(0)
	sellBindGold := int64(0)
	tempBindGold := float64(0)
	kunMap := inventoryManager.GetAllKun()
	if len(kunMap) == 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("onearena:当前没有鲲可以出售")
		playerlogic.SendSystemMessage(pl, lang.OneArenaKunNoExist)
		return
	}

	for itemId, num := range kunMap {
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		if itemTemplate == nil {
			continue
		}
		saleRate := float64(itemTemplate.SaleRate) / float64(common.MAX_RATE)
		if itemTemplate.BuySilver != 0 {
			tempSellSilver += float64(itemTemplate.BuySilver) * saleRate * float64(num)
			continue
		}
		if itemTemplate.BuyBindgold != 0 {
			tempBindGold += float64(itemTemplate.BuyBindgold) * saleRate * float64(num)
		}
	}

	sellSilver = int64(math.Ceil(tempSellSilver))
	sellBindGold = int64(math.Ceil(tempBindGold))

	inventoryLogReason := commonlog.InventoryLogReasonOneArenaSellKun
	reasonText := inventoryLogReason.String()
	flag := inventoryManager.BatchRemove(kunMap, inventoryLogReason, reasonText)
	if !flag {
		panic(fmt.Errorf("onearena: oneArenaSell BatchRemove should be ok"))
	}
	inventorylogic.SnapInventoryChanged(pl)

	manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	totalSiver, totalBindGold := manager.SellKunAddRes(sellSilver, int64(sellBindGold))

	reasonGoldText := commonlog.GoldLogReasonOneArenaSellKun.String()
	reasonSliverText := commonlog.SilverLogReasonOneArenaSellKun.String()
	flag = propertyManager.AddMoney(sellBindGold, sellGold, commonlog.GoldLogReasonOneArenaSellKun, reasonGoldText, sellSilver, commonlog.SilverLogReasonOneArenaSellKun, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("onearena: oneArenaSell AddMoney should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	scOneArenaSell := pbutil.BuildSCOneArenaSell(totalSiver, totalBindGold)
	pl.SendMsg(scOneArenaSell)
	return
}
