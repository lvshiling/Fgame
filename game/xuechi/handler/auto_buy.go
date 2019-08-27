package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fgame/fgame/game/xuechi/pbutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XUECHI_AUTO_BUY_TYPE), dispatch.HandlerFunc(handleXueChiAutoBuy))
}

//处理血池生命瓶自动购买
func handleXueChiAutoBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("xuechi:处理血池生命瓶自动购买")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXueChiAutoBuy := msg.(*uipb.CSXueChiAutoBuy)
	itemId := csXueChiAutoBuy.GetItemId()
	num := csXueChiAutoBuy.GetItemNum()
	err = xueChiAutoBuy(tpl, itemId, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
				"error":    err,
			}).Error("xuechi:处理血池生命瓶自动购买,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"num":      num,
		}).Debug("xuechi:处理血池生命瓶自动购买完成")
	return nil
}

//处理血池生命瓶自动购买逻辑
func xueChiAutoBuy(pl player.Player, itemId int32, num int32) (err error) {
	// manager := pl.GetPlayerDataManager(types.PlayerXueChiDataManagerType).(*playerxuechi.PlayerXueChiDataManager)
	itemTempalte := item.GetItemService().GetItem(int(itemId))
	if itemTempalte == nil {
		return
	}
	if itemTempalte.GetItemType() != itemtypes.ItemTypeLifeOrigin {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("xueChi:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	costGold := int32(0)
	costBindGold := int32(0)
	costSilver := int64(0)
	if num > 0 {
		if !shop.GetShopService().ShopIsSellItem(itemId) {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xueChi:商铺没有该道具,无法自动购买")
			playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
			return
		}

		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, itemId, num)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xueChi:购买血瓶物品失败")
			playerlogic.SendSystemMessage(pl, lang.ShopXueChiAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		costGold += int32(shopNeedGold)
		costBindGold += int32(shopNeedBindGold)
		costSilver += shopNeedSilver
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xuechi:银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xuechi:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够元宝
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("xuechi:元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	goldReason := commonlog.GoldLogReasonXueChiAutoBuy
	goldReasonText := fmt.Sprintf(goldReason.String(), itemId, num)
	silverReason := commonlog.SilverLogReasonXueChiAutoBuy
	silverReasonText := fmt.Sprintf(silverReason.String(), itemId, num)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonXueChiAutoBuy, goldReasonText, costSilver, commonlog.SilverLogReasonXueChiAutoBuy, silverReasonText)
	if !flag {
		panic(fmt.Errorf("xueChi: xueChiAutoBuy Cost should be ok"))
	}

	//同步元宝
	propertylogic.SnapChangedProperty(pl)
	addBlood := int64(itemTempalte.TypeFlag1) * int64(num)

	scMsg := pbutil.BuildSCXueChiAutoBuy(itemId, num, addBlood)
	pl.SendMsg(scMsg)

	pl.AddBlood(addBlood)
	scXueChiBlood := pbutil.BuildSCXueChiBlood(pl.GetBlood())
	pl.SendMsg(scXueChiBlood)
	return
}
