package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	discountdiscounttemplate "fgame/fgame/game/welfare/discount/discount/template"
	discountdiscounttypes "fgame/fgame/game/welfare/discount/discount/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MERGE_ACTIVITY_DISCOUNT_BUY_TYPE), dispatch.HandlerFunc(handlerDiscountBuy))
}

//处理购买限时礼包
func handlerDiscountBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理购买限时礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMergeActivityDiscountBuy := msg.(*uipb.CSMergeActivityDiscountBuy)
	groupId := csMergeActivityDiscountBuy.GetGroupId()
	discountId := csMergeActivityDiscountBuy.GetDiscountId()
	buyNum := csMergeActivityDiscountBuy.GetNum()

	err = buyDiscount(tpl, groupId, discountId, buyNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理购买限时礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理购买限时礼包请求完成")

	return
}

//购买限时礼包请求逻辑
func buyDiscount(pl player.Player, groupId, discountId, buyNum int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeCommon

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	curDiscountDay := welfarelogic.CountCurActivityDay(groupId)
	discountTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountTemplate(discountId)
	if curDiscountDay != discountTemp.DayGroup {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"curDiscountDay": curDiscountDay,
				"discountId":     discountId,
			}).Warn("welfare:购买限时礼包请求，折扣时间错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if buyNum > discountTemp.MaxCount {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"discountId": discountId,
				"num":        buyNum,
			}).Warn("welfare:购买限时礼包请求,超过单次最大购买数量")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	curBuyTimes := buyNum
	info := obj.GetActivityData().(*discountdiscounttypes.DiscountInfo)
	curBuyTimes += info.BuyRecord[discountTemp.Index]
	if discountTemp.LimitCount != 0 {
		if curBuyTimes > discountTemp.LimitCount {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"discountId": discountId,
					"num":        buyNum,
				}).Warn("welfare:购买限时礼包请求,超过每日购买上限")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityDiscountLimit)
			return
		}
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:购买限时礼包请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	groupTemp := groupInterface.(*discountdiscounttemplate.GroupTemplateDiscount)
	// 全服次数更新
	if groupTemp.IsGlobalTimesLimit() {
		if !welfare.GetWelfareService().IsHadDiscountTimes(groupId, discountTemp.Index, discountTemp.LimitQuanfu, buyNum) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"groupId":    groupId,
					"discountId": discountId,
					"buyNum":     buyNum,
				}).Warn("welfare:购买限时礼包请求，全服次数已领完")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityGlobalNotTimesReceiveDiscount)
			return
		}
	}

	//元宝是否足够
	needGold := buyNum * discountTemp.UseGold
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:购买限时礼包请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//判断背包空间
	itemId := discountTemp.ItemId
	totalItemNum := discountTemp.BuyCount * buyNum
	defaultLevel := int32(0)
	bind := itemtypes.ItemBindTypeUnBind
	expireType := groupTemp.GetExpireType()
	expireTime := groupTemp.GetExpireTime()
	now := global.GetGame().GetTimeService().Now()
	itemData := droptemplate.CreateItemDataWithExpire(itemId, totalItemNum, defaultLevel, bind, expireType, expireTime, now)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotItemLevelWithProperty(itemId, totalItemNum, defaultLevel, bind, expireType, itemData.GetItemGetTime(), itemData.GetExpireTimestamp()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"buyNum":   buyNum,
			}).Warn("welfare:背包空间不足，请清理后再购买")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyDiscountUse
	goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic("welfare: buy discount gift use gold should be ok")
	}

	//添加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)

	flag = inventoryManager.AddItemLevel(itemData, itemGetReason, itemReasonText)
	if !flag {
		panic("welfare: buy discount add item should be ok")
	}

	//更新信息
	_, ok := info.BuyRecord[discountTemp.Index]
	if !ok {
		info.BuyRecord[discountTemp.Index] = buyNum
	} else {
		info.BuyRecord[discountTemp.Index] += buyNum
	}
	welfareManager.UpdateObj(obj)

	// 全服次数更新
	if groupTemp.IsGlobalTimesLimit() {
		welfare.GetWelfareService().AddDiscountTimes(groupId, discountTemp.Index, buyNum)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	timesList := welfare.GetWelfareService().GetDiscountTimes(groupId)
	scMsg := pbutil.BuildSCMergeActivityDiscountBuy(groupId, discountId, buyNum, itemId, totalItemNum, timesList)
	pl.SendMsg(scMsg)
	return
}
