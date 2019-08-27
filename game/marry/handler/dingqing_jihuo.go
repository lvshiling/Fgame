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
	marrylogic "fgame/fgame/game/marry/logic"
	marryservice "fgame/fgame/game/marry/marry"
	pbutil "fgame/fgame/game/marry/pbutil"
	marryplayer "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_MarryDingQingJiHuo), dispatch.HandlerFunc(handleMarryDingQingJiHuo))
}

//处理定情信物激活
func handleMarryDingQingJiHuo(s session.Session, msg interface{}) (err error) {
	log.Debug("定情信物激活")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csDingQing := msg.(*uipb.CSMarryDingQingJiHuoMsg)
	suitId := csDingQing.GetSuitId()
	posId := csDingQing.GetPosId()
	autoFlag := csDingQing.GetAutoBuyFlag()
	err = dingQingJiHuo(tpl, suitId, posId, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
				"autoFlag": autoFlag,
				"err":      err,
			}).Error("marry:定情信物激活,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"suitId":   suitId,
			"posId":    posId,
			"autoFlag": autoFlag,
		}).Debug("定情信物激活,成功")
	return
}

func dingQingJiHuo(pl player.Player, suitId int32, posId int32, autoFlag bool) (err error) {
	item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(suitId, posId)
	if item == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
				"autoFlag": autoFlag,
			}).Warn("marry:定情信物激活,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuNotExists)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*marryplayer.PlayerMarryDataManager)
	flag := marryManager.ExistsDingQing(suitId, posId)
	if flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"suitId":   suitId,
				"posId":    posId,
				"autoFlag": autoFlag,
			}).Warn("marry:定情信物激活,已经存在")
		playerlogic.SendSystemMessage(pl, lang.MarryXinWuAlreadyExists)
		return
	}

	//判读是否满足条件
	if !inventoryManager.HasEnoughItem(item.GetItemId(), 1) {
		if !autoFlag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"suitId":   suitId,
					"posId":    posId,
					"autoFlag": autoFlag,
				}).Error("marry:定情信物激活,物品不足")
			playerlogic.SendSystemMessage(pl, lang.MarryXinWuItemNotEnough)
			return
		}
		flag = autoBuyDingQingItem(pl, item.GetItemId(), 1)
		if !flag {
			return
		}
	} else {
		text := commonlog.InventoryLogReasonMarryDingQingJiHuo.String()
		flag = inventoryManager.UseItem(item.GetItemId(), 1, commonlog.InventoryLogReasonMarryDingQingJiHuo, text)
		inventorylogic.SnapInventoryChanged(pl)
		if !flag {
			panic(fmt.Errorf("marry:定情信物激活,消耗物品应该成功"))
		}
	}

	marrylogic.AddPlayerDingQing(pl, suitId, posId)

	//发送消息
	spId := pl.GetSpouseId()
	exFlag := false
	if spId != 0 { //有结婚
		if marryservice.GetMarryService().ExistsSpouseDingQing(pl.GetId(), suitId, posId) { //有定情信物
			exFlag = true
		}
	}
	msg := pbutil.BuildSCMarryDingQingJiHuoMsg(spId, exFlag, suitId, posId)
	pl.SendMsg(msg)
	return
}

func autoBuyDingQingItem(pl player.Player, itemId int32, itemNum int32) bool {
	costGold := int32(0)
	costBindGold := int32(0)
	costSilver := int64(0)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !shop.GetShopService().ShopIsSellItem(itemId) {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"itemNum":  itemNum,
		}).Warn("marry:定情信物激活购买商铺没有该道具,无法自动购买")
		playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		return false
	}

	isEnoughBuyTimes, shopIdMap := shoplogic.MaxBuyTimesForPlayer(pl, itemId, itemNum)
	if !isEnoughBuyTimes {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"itemNum":  itemNum,
		}).Warn("marry:定情信物激活购买物品失败,自动购买已停止")
		playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
		return false
	}

	shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
	costGold += int32(shopNeedGold)
	costBindGold += int32(shopNeedBindGold)
	costSilver += shopNeedSilver

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
			}).Warn("marry:定情信物激活购买物品银两不足,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return false
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
			}).Warn("marry:定情信物激活购买物品元宝不足,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return false
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"itemNum":  itemNum,
			}).Warn("marry:定情信物激活购买物品元宝不足,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return false
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonMarryDingQingJiHuo.String()
	reasonSliverText := commonlog.SilverLogReasonMarryDingQingJiHuo.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonMarryDingQingJiHuo, reasonGoldText, costSilver, commonlog.SilverLogReasonMarryDingQingJiHuo, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("marry: 定情信物自动购买,应该成功"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}
	return true
}
