package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_DISCOUNT_ZHUANSHENG_BUY_ALL_TYPE), dispatch.HandlerFunc(handlerDiscountZhuanShengBuyAll))
}

//处理全部购买转生大礼包
func handlerDiscountZhuanShengBuyAll(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理全部购买转生大礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityDiscountZhuanShengBuyAll)
	groupId := csMsg.GetGroupId()
	timesInfoList := csMsg.GetBuyInfoList()

	buyGiftMap := make(map[int32]int32)
	for _, timesInfo := range timesInfoList {
		giftType := timesInfo.GetKey()
		buyNum := timesInfo.GetValue()
		_, ok := buyGiftMap[giftType]
		if !ok {
			buyGiftMap[giftType] = buyNum
		} else {
			buyGiftMap[giftType] += buyNum
		}
	}

	err = buyDiscountZhuanShengAll(tpl, groupId, buyGiftMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理全部购买转生大礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理全部购买转生大礼包请求完成")

	return
}

//全部购买转生大礼包请求逻辑
func buyDiscountZhuanShengAll(pl player.Player, groupId int32, giftBuyMap map[int32]int32) (err error) {
	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	if !welfarelogic.IsOnActivityTime(groupId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:购买转生大礼包请求,不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(groupId)
	if groupTemp == nil {
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)

	initClassType := int32(0)
	totalBuyNum := int32(0)
	for giftType, buyNum := range giftBuyMap {
		totalBuyNum += buyNum

		discountTemp := groupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
		if discountTemp == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"groupId":  groupId,
					"giftType": giftType,
				}).Warn("welfare:全部购买转生大礼包请求,模板不存在")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		if initClassType == 0 {
			initClassType = discountTemp.Bargain
		}
		// 不是同一标签页的物品
		if initClassType != 0 && initClassType != discountTemp.Bargain {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"initClassType": initClassType,
					"curClassType":  discountTemp.Bargain,
				}).Warn("welfare:全部购买转生大礼包请求,礼包标签类型错误")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		// if buyNum > discountTemp.MaxCount {
		// 	log.WithFields(
		// 		log.Fields{
		// 			"playerId": pl.GetId(),
		// 			"num":      buyNum,
		// 		}).Warn("welfare:全部购买转生大礼包请求,超过单次最大全部购买数量")
		// 	playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		// 	return
		// }

		if info.ChargeNum < int64(discountTemp.NeedChongZhi) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"num":      buyNum,
				}).Warn("welfare:全部购买转生大礼包请求,充值条件不足")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityChargeNotEnoughCondition)
			return
		}

		curBuyTimes := info.BuyRecord[giftType] + buyNum
		if curBuyTimes > discountTemp.BuyMax {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"buyMax":      discountTemp.BuyMax,
					"curBuyTimes": curBuyTimes,
					"num":         buyNum,
				}).Warn("welfare:全部购买转生大礼包请求,超过全部购买上限")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityDiscountZhuanShengLimit)
			return
		}
	}

	// 折扣
	//classType和购买数量 取折扣
	initDazheRatio := int32(common.MAX_RATE)
	zhuanshengBargainTemp := welfaretemplate.GetWelfareTemplateService().GetZhuanShengCircleBargainTemplate(initClassType, totalBuyNum)
	if zhuanshengBargainTemp != nil {
		initDazheRatio = zhuanshengBargainTemp.Discount
	}

	totalGold := int32(0)
	totalUseItemMap := make(map[int32]int32)
	giftTypeNeedGoldMap := make(map[int32]int32)
	giftTypeUseItemMap := make(map[int32]map[int32]int32)
	for giftType, buyNum := range giftBuyMap {
		discountTemp := groupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
		//元宝
		costGold := int32(math.Ceil(float64(buyNum*discountTemp.UseGold*initDazheRatio) / float64(common.MAX_RATE)))
		totalGold += costGold
		giftTypeNeedGoldMap[giftType] = buyNum * discountTemp.UseGold

		//消耗物品
		useItemMap := make(map[int32]int32)
		for itemId, num := range discountTemp.GetUseItemMap() {
			_, ok := totalUseItemMap[itemId]
			if ok {
				totalUseItemMap[itemId] += num * buyNum
			} else {
				totalUseItemMap[itemId] = num * buyNum
			}

			_, ok = useItemMap[itemId]
			if ok {
				useItemMap[itemId] += num * buyNum
			} else {
				useItemMap[itemId] = num * buyNum
			}
		}
		giftTypeUseItemMap[giftType] = useItemMap
	}

	//元宝是否足够
	isOneCondition := false
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if totalGold > 0 {
		if !propertyManager.HasEnoughGold(int64(totalGold), false) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"totalGold": totalGold,
				}).Warn("welfare:全部购买转生大礼包请求，当前元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
		isOneCondition = true
	}

	if len(totalUseItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(totalUseItemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("welfare:全部购买转生大礼包请求，物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		isOneCondition = true
	}

	if !isOneCondition {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"groupId":         groupId,
				"totalGold":       totalGold,
				"totalUseItemMap": totalUseItemMap,
				"giftBuyMap":      giftBuyMap,
			}).Warn("welfare:全部购买转生大礼包请求，不满足全部购买条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCostNotEnoughCondition)
		return
	}

	addItemMap, success := welfarelogic.BuyZhuanShengGift(pl, obj, giftBuyMap)
	if !success {
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyDiscountUse
	goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
	flag := propertyManager.CostGold(int64(totalGold), false, goldReason, goldReasonText)
	if !flag {
		panic("welfare: buy discount gift use gold should be ok")
	}

	if len(totalUseItemMap) > 0 {
		useItemReason := commonlog.InventoryLogReasonOpenActivityUse
		useItemReasonText := fmt.Sprintf(useItemReason.String(), typ, subType)
		flag := inventoryManager.BatchRemove(totalUseItemMap, useItemReason, useItemReasonText)
		if !flag {
			panic("welfare:抽奖批量消耗物品应该成功")
		}
	}

	welfareManager.UpdateObj(obj)
	// 不是组合折扣购买的公告
	if zhuanshengBargainTemp == nil {

		for giftType, buyNum := range giftBuyMap {
			needGold := giftTypeNeedGoldMap[giftType]
			useItemMap := giftTypeUseItemMap[giftType]

			//全部购买事件
			eventData := welfareeventtypes.CreatePlayerAllianceCheerEventData(groupId, giftType, needGold)
			gameevent.Emit(welfareeventtypes.EventTypeDiscountBuyZhuanShengGift, pl, eventData)

			//公告
			discountTemp := groupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
			getItemMap := make(map[int32]int32)
			for itemId, num := range discountTemp.GetItemMap() {
				_, ok := getItemMap[itemId]
				if ok {
					getItemMap[itemId] += num * buyNum
				} else {
					getItemMap[itemId] = num * buyNum
				}
			}
			itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(getItemMap)
			if len(itemNameLinkStr) > 0 {
				timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
				plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
				acName := chatlogic.FormatModuleNameNoticeStr(timeTemp.Name)
				args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
				link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)

				// 使用元宝购买的公告
				if needGold > 0 && len(useItemMap) < 1 && discountTemp.DaZhe < 10 {
					costGlodNum := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGold))
					yuanGold := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), discountTemp.YuanGold))
					discount := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonDiscountString), discountTemp.DaZhe))
					content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityDiscountGiftNotice), plName, acName, costGlodNum, itemNameLinkStr, yuanGold, discount, link)
					chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
					noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
				}

				//使用物品购买的公告
				if needGold <= 0 && len(useItemMap) > 0 {
					useItemStr := welfarelogic.RewardsItemNoticeStr(useItemMap)
					content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityDiscountGiftUseItemNotice), plName, acName, useItemStr, itemNameLinkStr, link)
					chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
					noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
				}
			}
		}
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityDiscountZhuanShengBuyAll(groupId, addItemMap, giftBuyMap)
	pl.SendMsg(scMsg)
	return
}
