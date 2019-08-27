package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_LEARN_UPLEVEL_TYPE), dispatch.HandlerFunc(handleBabyLearnUplevel))
}

//处理宝宝读书
func handleBabyLearnUplevel(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理宝宝读书消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSBabyLearnUplevel)
	babyId := csMsg.GetBabyId()
	itemId := csMsg.GetItemId()
	num := csMsg.GetNum()
	isAuto := csMsg.GetIsAuto()

	err = handlerLearnUplevel(tpl, babyId, itemId, num, isAuto)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理宝宝读书消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理宝宝读书消息完成")
	return nil

}

// 读书
func handlerLearnUplevel(pl player.Player, babyId int64, itemId, num int32, isAuto bool) (err error) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("baby:处理宝宝读书,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	baby := babyManager.GetBabyInfo(babyId)
	if baby == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"babyId":   babyId,
			}).Warn("baby:处理宝宝读书, 宝宝不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	pregnantTemplate := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(baby.GetQuality())
	if baby.GetLearnLevel() >= pregnantTemplate.LevelMax {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("baby:处理宝宝读书,读书达最高级")
		playerlogic.SendSystemMessage(pl, lang.BabyLearnFullLevel)
		return
	}

	nextLevel := baby.GetLearnLevel() + 1
	nextLearnTemp := babytemplate.GetBabyTemplateService().GetBabyLearnTemplate(nextLevel)
	if nextLearnTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("baby:处理宝宝读书,下一级读书模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	numOfItems := inventoryManager.NumOfItems(itemId)
	costItemNum := num
	//扣除物品
	if numOfItems < num {
		//不自动购买
		if !isAuto {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("baby:处理宝宝读书,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		costItemNum = numOfItems
		//计算需要购买的
		needBuy := num - numOfItems
		needBindGold := int64(0)
		needGold := int64(0)
		needSilver := int64(0)
		if !shop.GetShopService().ShopIsSellItem(itemId) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("baby:商铺没有该道具,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
			return
		}

		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, itemId, needBuy)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("baby:购买宝宝读书物品失败")
			playerlogic.SendSystemMessage(pl, lang.ShopBabyLearnAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		needGold += shopNeedGold
		needBindGold += shopNeedBindGold
		needSilver += shopNeedSilver

		if !propertyManager.HasEnoughCost(needBindGold, needGold, needSilver) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("baby:处理宝宝读书物品,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//更新自动购买每日限购次数
		if len(shopIdMap) != 0 {
			shoplogic.ShopDayCountChanged(pl, shopIdMap)
		}
		goldReason := commonlog.GoldLogReasonBabyLearnBuyUse
		goldReasonText := fmt.Sprintf(goldReason.String(), itemId, needBuy, babyId, baby.GetLearnLevel(), baby.GetQuality())
		silverReason := commonlog.SilverLogReasonBabyLearnBuyUse
		silverReasonText := fmt.Sprintf(silverReason.String(), itemId, needBuy, babyId, baby.GetLearnLevel(), baby.GetQuality())
		flag := propertyManager.Cost(needBindGold, needGold, goldReason, goldReasonText, needSilver, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("baby:购买宝宝读书物品,应该成功"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	if costItemNum > 0 {
		itemUseReason := commonlog.InventoryLogReasonBabyLearnUse
		itemUseReasonText := fmt.Sprintf(itemUseReason.String(), babyId, baby.GetLearnLevel())
		flag := inventoryManager.UseItem(itemId, costItemNum, itemUseReason, itemUseReasonText)
		if !flag {
			panic(fmt.Errorf("baby:宝宝读书消耗物品应该成功"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	addLearnExp := itemTemplate.TypeFlag1 * num
	isUplevel := babyManager.AddLearnExp(babyId, addLearnExp)
	if isUplevel {
		babylogic.BabyPropertyChanged(pl)
	}

	scMsg := pbutil.BuildSCBabyLearnUplevel(babyId, itemId, num, isAuto, baby)
	pl.SendMsg(scMsg)
	return
}
