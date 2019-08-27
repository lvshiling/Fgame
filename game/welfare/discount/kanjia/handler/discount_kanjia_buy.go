package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	discountkanjiatypes "fgame/fgame/game/welfare/discount/kanjia/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_KANJIA_BUY_TYPE), dispatch.HandlerFunc(handlerKanJiaBuy))
}

//处理购买砍价礼包
func handlerKanJiaBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理购买砍价礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityKanJiaBuy)
	groupId := csMsg.GetGroupId()
	giftType := csMsg.GetType()

	err = buyKanJia(tpl, groupId, giftType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理购买砍价礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理购买砍价礼包请求完成")

	return
}

//购买砍价礼包请求逻辑
func buyKanJia(pl player.Player, groupId int32, giftType int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeKanJia

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	groupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountBargainShopGroupTemplate(groupId)
	if groupTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:购买砍价礼包请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	discountTemp := groupTemp.GetDiscountBargainTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
	if discountTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"giftType": giftType,
			}).Warn("welfare:购买砍价礼包请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*discountkanjiatypes.DiscountKanJiaInfo)

	//元宝是否足够
	_, discount := info.GetKanJiaInfo(giftType)
	needGold := math.Ceil(float64(discountTemp.YuanGold*discount) / float64(common.MAX_RATE))
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(int64(needGold), false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("welfare:购买砍价礼包请求，当前元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//判断背包空间
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	firsOpenTemp := groupInterface.GetFirstOpenTemp()
	itemDataList := welfarelogic.ConvertToItemData(discountTemp.GetItemMap(), firsOpenTemp.GetExpireType(), firsOpenTemp.GetExpireTime())
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:背包空间不足，请清理后再购买")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonOpenActivityCost
	goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic("welfare: buy discount gift use gold should be ok")
	}

	//添加物品
	itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	itemReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
	flag = inventoryManager.BatchAddOfItemLevel(itemDataList, itemGetReason, itemReasonText)
	if !flag {
		panic("welfare: buy discount add item should be ok")
	}

	//更新信息
	info.AddBuyRecord(giftType)
	welfareManager.UpdateObj(obj)

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityKanJiaBuy(groupId, int32(giftType), info.BuyRecord)
	pl.SendMsg(scMsg)
	return
}
