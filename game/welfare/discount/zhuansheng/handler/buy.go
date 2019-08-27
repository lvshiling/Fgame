package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
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
	discountzhuanshengtemplate "fgame/fgame/game/welfare/discount/zhuansheng/template"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_DISCOUNT_ZHUANSHENG_BUY_TYPE), dispatch.HandlerFunc(handlerDiscountZhuanShengBuy))
}

//处理购买转生大礼包
func handlerDiscountZhuanShengBuy(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理购买转生大礼包请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityDiscountZhuanShengBuy)
	groupId := csMsg.GetGroupId()
	giftType := csMsg.GetTyp()
	buyNum := csMsg.GetNum()

	err = buyDiscountZhuanSheng(tpl, groupId, giftType, buyNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("welfare:处理购买转生大礼包请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("welfare:处理购买转生大礼包请求完成")

	return
}

//购买转生大礼包请求逻辑
func buyDiscountZhuanSheng(pl player.Player, groupId int32, giftType int32, buyNum int32) (err error) {
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
				"giftType": giftType,
			}).Warn("welfare:购买转生大礼包请求,不是活动时间")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
		return
	}

	groupTemp := welfaretemplate.GetWelfareTemplateService().GetDiscountZhuanShengGroupTemplate(groupId)
	if groupTemp == nil {
		return
	}
	discountTemp := groupTemp.GetDiscountZhuanShengTemplateByType(pl.GetRole(), pl.GetSex(), giftType)
	if discountTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
				"giftType": giftType,
			}).Warn("welfare:购买转生大礼包请求,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	if buyNum > discountTemp.MaxCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      buyNum,
			}).Warn("welfare:购买转生大礼包请求,超过单次最大购买数量")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
	if info.ChargeNum < int64(discountTemp.NeedChongZhi) {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"num":          buyNum,
				"NeedChongZhi": info.ChargeNum,
			}).Warn("welfare:购买转生大礼包请求,充值条件不足")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityChargeNotEnoughCondition)
		return
	}

	if discountTemp.BuyMax > 0 {
		curBuyTimes := buyNum
		curBuyTimes += info.BuyRecord[giftType]
		if curBuyTimes > discountTemp.BuyMax {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"buyMax":   discountTemp.BuyMax,
					"num":      buyNum,
				}).Warn("welfare:购买转生大礼包请求,超过购买上限")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityDiscountZhuanShengLimit)
			return
		}
	}

	//元宝是否足够
	isOneCondition := false
	needGold := buyNum * discountTemp.UseGold
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if needGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needGold), false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"needGold": needGold,
				}).Warn("welfare:购买转生大礼包请求，当前元宝不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
		isOneCondition = true
	}

	//消耗物品
	useItemMap := make(map[int32]int32)
	for itemId, num := range discountTemp.GetUseItemMap() {
		_, ok := useItemMap[itemId]
		if ok {
			useItemMap[itemId] += num * buyNum
		} else {
			useItemMap[itemId] = num * buyNum
		}
	}
	if len(useItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(useItemMap) {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"groupId":    groupId,
					"buyNum":     buyNum,
					"useItemMap": useItemMap,
				}).Warn("welfare:购买转生大礼包请求，物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		isOneCondition = true
	}

	needPoint := buyNum * discountTemp.UsePoint
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemplate := groupInterface.(*discountzhuanshengtemplate.GroupTemplateZhaunSheng)
	if needPoint > 0 {
		_, leftPoint := groupTemplate.GetTotalAndRemainPoint(info.ChargeNum, info.UsePoint)
		if leftPoint < needPoint {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"needPoint": needPoint,
				}).Warn("welfare:购买转生大礼包请求，当前充值积分不足")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityDiscountChargePointNoEnough)
			return
		}
		isOneCondition = true
	}

	if !isOneCondition {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("welfare:购买转生大礼包请求，不满足购买条件")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCostNotEnoughCondition)
		return
	}

	// //判断背包空间
	// inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// newItemMap := make(map[int32]int32)
	// for itemId, num := range discountTemp.GetItemMap() {
	// 	_, ok := newItemMap[itemId]
	// 	if ok {
	// 		newItemMap[itemId] += num * buyNum
	// 	} else {
	// 		newItemMap[itemId] = num * buyNum
	// 	}
	// }
	// itemDataList := welfarelogic.ConvertToItemData(newItemMap)
	// if !inventoryManager.HasEnoughSlotsOfItemLevel(itemDataList) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":   pl.GetId(),
	// 			"newItemMap": newItemMap,
	// 		}).Warn("welfare:背包空间不足，请清理后再购买")
	// 	playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
	// 	return
	// }

	addItemMap, success := welfarelogic.BuyZhuanShengGift(pl, obj, map[int32]int32{giftType: buyNum})
	if !success {
		// log.WithFields(
		// 	log.Fields{
		// 		"playerId": pl.GetId(),
		// 	}).Warn("welfare:购买转生大礼包请求，购买失败")
		// playerlogic.SendSystemMessage(pl, lang.OpenActivityBuyFaild)
		return
	}

	//消耗元宝
	goldReason := commonlog.GoldLogReasonBuyDiscountUse
	goldReasonText := fmt.Sprintf(goldReason.String(), typ, subType)
	flag := propertyManager.CostGold(int64(needGold), false, goldReason, goldReasonText)
	if !flag {
		panic("welfare: buy discount gift use gold should be ok")
	}

	if len(useItemMap) > 0 {
		useItemReason := commonlog.InventoryLogReasonOpenActivityUse
		useItemReasonText := fmt.Sprintf(useItemReason.String(), typ, subType)
		flag := inventoryManager.BatchRemove(useItemMap, useItemReason, useItemReasonText)
		if !flag {
			panic("welfare:抽奖批量消耗物品应该成功")
		}
	}

	//使用充值积分
	if needPoint > 0 {
		info.UsePoint += needPoint
		//welfareManager.UpdateObj(obj)
	}

	// //添加物品
	// itemGetReason := commonlog.InventoryLogReasonOpenActivityRew
	// itemReasonText := fmt.Sprintf(itemGetReason.String(), typ, subType)
	// flag = inventoryManager.BatchAddOfItemLevel(itemDataList, itemGetReason, itemReasonText)
	// if !flag {
	// 	panic("welfare: buy discount add item should be ok")
	// }

	// //更新信息
	// _, ok := info.BuyRecord[giftType]
	// if !ok {
	// 	info.BuyRecord[giftType] = buyNum
	// } else {
	// 	info.BuyRecord[giftType] += buyNum
	// }
	welfareManager.UpdateObj(obj)

	//购买事件
	eventData := welfareeventtypes.CreatePlayerAllianceCheerEventData(groupId, giftType, needGold)
	gameevent.Emit(welfareeventtypes.EventTypeDiscountBuyZhuanShengGift, pl, eventData)

	//公告
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(addItemMap)
	if len(itemNameLinkStr) > 0 {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		acName := chatlogic.FormatModuleNameNoticeStr(timeTemp.Name)
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)

		// 使用元宝购买的公告
		if needGold > 0 && len(useItemMap) < 1 && discountTemp.DaZhe < 10 && groupTemplate.GetTimeTemplate().IsChuanYin() {
			costGlodNum := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGold))
			yuanGold := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), discountTemp.YuanGold))
			discount := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonDiscountString), discountTemp.DaZhe))
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityDiscountGiftNotice), plName, acName, costGlodNum, itemNameLinkStr, yuanGold, discount, link)
			chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
			noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
		}

		//使用物品购买的公告
		if needGold <= 0 && len(useItemMap) > 0 && groupTemplate.GetTimeTemplate().IsChuanYin() {
			useItemStr := welfarelogic.RewardsItemNoticeStr(useItemMap)
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityDiscountGiftUseItemNotice), plName, acName, useItemStr, itemNameLinkStr, link)
			chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
			noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
		}
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityDiscountZhuanShengBuy(addItemMap, giftType, groupId, buyNum, info.UsePoint)
	pl.SendMsg(scMsg)
	return
}

func checkRelateGroupBuyGift(pl player.Player, groupId int32) (flag bool) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
		if !welfarelogic.IsOnActivityTime(relateGroupId) {
			continue
		}

		//城战助威
		relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
		if relateTimeTemp.IsAllianceCheer() {
			activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeAlliance)
			if activityTemp == nil {
				return
			}

			// 是否活动日
			now := global.GetGame().GetTimeService().Now()
			openTime := welfare.GetWelfareService().GetServerStartTime()  //global.GetGame().GetServerTime()
			mergeTime := welfare.GetWelfareService().GetServerMergeTime() //merge.GetMergeService().GetMergeTime()

			activityTimeTemp := activityTemp.GetOnDateTimeTemplate(now, openTime, mergeTime)
			if activityTimeTemp == nil {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"groupId":  groupId,
					}).Warn("welfare:购买转生大礼包请求,不是九霄城战日")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityNotOnTime)
				return
			}

			// 是否活动结束
			beginTime, _ := activityTimeTemp.GetBeginTime(now)
			if now >= beginTime {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"groupId":  groupId,
					}).Warn("welfare:购买转生大礼包请求,城战助威已开启")
				playerlogic.SendSystemMessage(pl, lang.OpenActivityAllianceCheerBuyFaild)
				return
			}
		}
	}

	flag = true
	return
}
