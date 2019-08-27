package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	playerarena "fgame/fgame/game/arena/player"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerring "fgame/fgame/game/ring/player"
	ringtypes "fgame/fgame/game/ring/types"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shop/pbutil"
	playershop "fgame/fgame/game/shop/player"
	"fgame/fgame/game/shop/shop"
	shoptypes "fgame/fgame/game/shop/types"
	shopdiscountlogic "fgame/fgame/game/shopdiscount/logic"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHOP_BUY_TYPE), dispatch.HandlerFunc(handleShopBuy))
}

//处理商店购买道具
func handleShopBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("shop:处理商店购买道具")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShopBuy := msg.(*uipb.CSShopBuy)
	shopId := csShopBuy.GetShopId()
	num := csShopBuy.GetNum()

	err = shopBuy(tpl, shopId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shop:处理商店购买道具,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shop:处理商店购买道具完成")
	return nil

}

//商店购买道具的逻辑
func shopBuy(pl player.Player, shopId int32, num int32) (err error) {
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil || num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if num > shopTemplate.MaxCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:购买数量大于最大购买数量")
		playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
		return
	}

	//购买总数
	totalNum := int32(shopTemplate.BuyCount * num)

	shopManager := pl.GetPlayerDataManager(playertypes.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	isLimitBuy, leftNum := shopManager.LeftDayCount(shopId)
	if isLimitBuy && leftNum < totalNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:购买次数，已达每日限购数量")
		playerlogic.SendSystemMessage(pl, lang.ShoBuyReacheLimit)
		return
	}

	//货币判断
	flag := true
	consumeType := shoptypes.ShopConsumeType(shopTemplate.ConsumeType)
	discountRatio := shopdiscountlogic.GetShopDiscount(pl, consumeType)
	consume := int32(math.Ceil(float64(shopTemplate.ConsumeData1*num) * discountRatio / float64(common.MAX_RATE)))
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	// chuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
	switch consumeType {
	case shoptypes.ShopConsumeTypeBindGold:
		flag = propertyManager.HasEnoughGold(int64(consume), true)
		break
	case shoptypes.ShopConsumeTypeGold:
		flag = propertyManager.HasEnoughGold(int64(consume), false)
		break
	case shoptypes.ShopConsumeTypeSliver:
		flag = propertyManager.HasEnoughSilver(int64(consume))
		break
	case shoptypes.ShopConsumeTypeGongXun:
		gongXunNum := pl.GetShenMoGongXunNum()
		if gongXunNum < consume {
			flag = false
		} else {
			flag = true
		}
	case shoptypes.ShopConsumeTypeItem:
		{
			flag = inventoryManager.HasEnoughItem(shopTemplate.ConsumeItemId, consume)
			break
		}
	case shoptypes.ShopConsumeTypeArenaJiFen:
		{
			arenaObj := arenaManager.GetPlayerArenaObjectByRefresh()
			flag = arenaObj.IfEnoughPoint(consume)
		}
	case shoptypes.ShopConsumeTypeArenapvpJiFen:
		{
			arenapvpObj := arenapvpManager.GetPlayerArenapvpObj()
			flag = arenapvpObj.IfEnoughJiFen(consume)
		}
	case shoptypes.ShopConsumeTypeChuangShiJiFen:
		{
			// chuangShiObj := chuangShiManager.GetPlayerChuangShiInfo()
			// flag = chuangShiObj.IfEnoughJiFen(int64(consume))
		}
	case shoptypes.ShopConsumeTypeRingXunBaoJiFen:
		{
			ringObj := ringManager.GetPlayerBaoKuObject(ringtypes.BaoKuTypeRing)
			flag = ringObj.IfEnoughJiFen(consume)
		}
	default:
		break
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
				"consume":  consume,
			}).Warn("shop:元宝不足，无法完成购买")

		switch consumeType {
		case shoptypes.ShopConsumeTypeGongXun:
			playerlogic.SendSystemMessage(pl, lang.PlayerGongXunNoEnough)
		case shoptypes.ShopConsumeTypeItem:
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		case shoptypes.ShopConsumeTypeArenaJiFen:
			playerlogic.SendSystemMessage(pl, lang.PlayerArenaPointNoEnough)
		case shoptypes.ShopConsumeTypeArenapvpJiFen:
			playerlogic.SendSystemMessage(pl, lang.PlayerArenapvpPointNoEnough)
		case shoptypes.ShopConsumeTypeChuangShiJiFen:
			playerlogic.SendSystemMessage(pl, lang.ChuangShiJiFenNoEnough)
		case shoptypes.ShopConsumeTypeRingXunBaoJiFen:
			playerlogic.SendSystemMessage(pl, lang.RingAttendPointsNotEnough)
		default:
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		}
		return
	}

	//判断背包空间
	itemId := shopTemplate.ItemId
	flag = inventoryManager.HasEnoughSlot(itemId, totalNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:背包空间不足，请清理后再购买")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗货币
	switch consumeType {
	case shoptypes.ShopConsumeTypeBindGold:
		goldReason := commonlog.GoldLogReasonShopBuyItem
		reasonText := fmt.Sprintf(goldReason.String(), shopId, shopTemplate.Name, num)
		flag = propertyManager.CostGold(int64(consume), true, goldReason, reasonText)
		break
	case shoptypes.ShopConsumeTypeGold:
		goldReason := commonlog.GoldLogReasonShopBuyItem
		reasonText := fmt.Sprintf(goldReason.String(), shopId, shopTemplate.Name, num)
		flag = propertyManager.CostGold(int64(consume), false, goldReason, reasonText)
		break
	case shoptypes.ShopConsumeTypeSliver:
		silverReason := commonlog.SilverLogReasonShopBuyItem
		reasonText := fmt.Sprintf(silverReason.String(), shopId, shopTemplate.Name, num)
		flag = propertyManager.CostSilver(int64(consume), silverReason, reasonText)
		break
	case shoptypes.ShopConsumeTypeGongXun:
		gongXunNum := pl.GetShenMoGongXunNum()
		leftGongXunNum := gongXunNum - consume
		if leftGongXunNum < 0 {
			flag = false
		} else {
			flag = true
			pl.SetShenMoGongXunNum(leftGongXunNum)
		}
	case shoptypes.ShopConsumeTypeItem:
		{
			useItemReason := commonlog.InventoryLogReasonShopBuyItem
			useItemReasonText := fmt.Sprintf(useItemReason.String(), shopId, num)
			flag = inventoryManager.UseItem(shopTemplate.ConsumeItemId, consume, useItemReason, useItemReasonText)
			break
		}
	case shoptypes.ShopConsumeTypeArenaJiFen:
		{
			flag = arenaManager.UsePoint(consume)
			break
		}
	case shoptypes.ShopConsumeTypeArenapvpJiFen:
		{
			flag = arenapvpManager.UseJiFen(consume)
			break
		}
	case shoptypes.ShopConsumeTypeChuangShiJiFen:
		{
			// flag = chuangShiManager.UseJiFen(int64(consume))
			break
		}
	case shoptypes.ShopConsumeTypeRingXunBaoJiFen:
		{
			flag = ringManager.UseJiFen(ringtypes.BaoKuTypeRing, consume)
			break
		}

	default:
		break
	}
	if !flag {
		panic("shop: cost Gold/Silver/GongXun/Item/arenaPoint/ChuangShiJiFen should be ok")
	}

	//添加物品
	reasonText := commonlog.InventoryLogReasonShopBuy.String()
	flag = inventoryManager.AddItem(itemId, totalNum, commonlog.InventoryLogReasonShopBuy, reasonText)
	if !flag {
		panic(fmt.Errorf("shop: shopBuy add item should be ok"))
	}

	//同步元宝
	propertylogic.SnapChangedProperty(pl)
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	//更新当日购买次数
	dayCount := int32(0)
	shopManager.UpdateObject(shopId, totalNum, false)
	if shopTemplate.LimitCount != 0 {
		dayCount = shopManager.GetShopBuyByShopId(shopId).DayCount
	}
	scShopBuy := pbutil.BuildSCShopBuy(shopId, num, dayCount)
	pl.SendMsg(scShopBuy)
	return
}
