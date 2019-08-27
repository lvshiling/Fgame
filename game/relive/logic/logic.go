package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/relive/pbutil"
	"fgame/fgame/game/scene/scene"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func SyncReliveInfo(pl scene.Player) {
	culTime := pl.GetCulReliveTime()
	lastReliveTime := pl.GetLastReliveTime()
	scReliveInfo := pbutil.BuildSCReliveInfo(culTime, lastReliveTime)
	pl.SendMsg(scReliveInfo)
}

//一般的正常复活
func Relive(tpl scene.Player, autoBuy bool) (sucess bool) {
	pl, ok := tpl.(player.Player)
	if !ok {
		return
	}

	maxLevel := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveNoNeedItemsBeforeLevel)
	if pl.GetLevel() < maxLevel {
		pl.Reborn(pl.GetPosition())
		return
	}

	pl.RefreshReliveTime()

	culTime := pl.GetCulReliveTime()
	culTime += 1
	flag := PlayerReliveTimeCost(pl, culTime, autoBuy)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("relive:原地复活,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	pl.Reborn(pl.GetPosition())
	pl.Relive()
	inventorylogic.SnapInventoryChanged(pl)
	if autoBuy {
		propertylogic.SnapChangedProperty(pl)
	}
	sucess = true
	return
}

func PlayerReliveTimeCost(pl player.Player, culTime int32, autoBuy bool) bool {
	mi := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeItemNumAddEveryRelive)) / float64(common.MAX_RATE)
	first := float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeFirstReliveItemNum))
	reliveItemId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveItemId)
	reliveItemLimit := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeResurrectionDanLimit)

	//判断消耗
	itemCount := int32(math.Ceil(first * math.Pow(float64(culTime), mi)))
	if itemCount > reliveItemLimit {
		itemCount = reliveItemLimit
	}

	costItemCount := itemCount

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	totalNum := inventoryManager.NumOfItems(int32(reliveItemId))
	costGold := int64(0)
	costBindGold := int64(0)
	costSilver := int64(0)
	if totalNum < itemCount {
		if autoBuy == false {
			return false
		}

		needBuyNum := itemCount - totalNum
		itemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(reliveItemId)
		// if shopTemplate == nil {
		// 	return false
		// }
		// shopNeedGold, shopNeedBindGold, shopNeedSilver := shopTemplate.GetConsumeData(needBuyNum)

		isEnoughBuyTimes := true
		shopIdMap := make(map[int32]int32)
		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(reliveItemId) {
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return false
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, reliveItemId, needBuyNum)
			if !isEnoughBuyTimes {
				playerlogic.SendSystemMessage(pl, lang.ShopReliveAutoBuyItemFail)
				return false
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += int64(shopNeedGold)
			costBindGold += int64(shopNeedBindGold)
			costSilver += shopNeedSilver
		}

		// costGold = int64(shopNeedGold)
		// costBindGold = int64(shopNeedBindGold)
		// costSilver = shopNeedSilver
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
		if costGold != 0 {
			flag := propertyManager.HasEnoughGold(int64(costGold), false)
			if !flag {
				return false
			}
			goldReason := commonlog.GoldLogReasonReliveAutoBuy
			goldReasonText := fmt.Sprintf(goldReason.String(), needBuyNum)
			flag = propertyManager.CostGold(costGold, true, goldReason, goldReasonText)
			if !flag {
				panic(fmt.Errorf("relive:花钱买复活丹应该成功"))
			}
		}
		if costBindGold != 0 {
			flag := propertyManager.HasEnoughGold(int64(costBindGold), true)
			if !flag {
				return false
			}
			goldReason := commonlog.GoldLogReasonReliveAutoBuy
			goldReasonText := fmt.Sprintf(goldReason.String(), needBuyNum)
			flag = propertyManager.CostGold(costBindGold, true, goldReason, goldReasonText)
			if !flag {
				panic(fmt.Errorf("relive:花钱买复活丹应该成功"))
			}
		}
		if costSilver != 0 {
			flag := propertyManager.HasEnoughGold(int64(costBindGold), true)
			if !flag {
				return false
			}
			silverReason := commonlog.SilverLogReasonReliveAutoBuy
			silverReasonText := fmt.Sprintf(silverReason.String(), needBuyNum)
			flag = propertyManager.CostSilver(costSilver, silverReason, silverReasonText)
			if !flag {
				panic(fmt.Errorf("relive:花钱买复活丹应该成功"))
			}
		}
		costItemCount -= needBuyNum

		//更新自动购买每日限购次数
		if len(shopIdMap) != 0 {
			shoplogic.ShopDayCountChanged(pl, shopIdMap)
		}
	}

	if costItemCount > 0 {
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonRelive.String(), culTime)
		flag := inventoryManager.UseItem(reliveItemId, costItemCount, commonlog.InventoryLogReasonRelive, reasonText)
		if !flag {
			panic(fmt.Errorf("relive:使用复活道具应该成功"))
		}
	}
	return true
}
