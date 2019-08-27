package handler

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	cyclechargesinglemaxrewmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/template"
	cyclechargesinglemaxrewmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/types"
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
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple, welfare.ReceiveHandlerFunc(cycleSingleMaxRewMultipleReceive))
}

//每日单笔充值领取奖励请求逻辑
func cycleSingleMaxRewMultipleReceive(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeCycleCharge
	subType := welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:领取每日单笔充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupId := openTemp.Group

	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取每日单笔充值奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*cyclechargesinglemaxrewmultipletypes.CycleSingleChargeMaxRewMultipleInfo)

	cycDay := openTemp.Value1
	if cycDay != info.CycleDay {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"rewId":     rewId,
				"cycDay":    cycDay,
				"curCycDay": info.CycleDay,
			}).Warn("welfare:领取每日单笔充值奖励请求,充值日类型错误")
		playerlogic.SendSystemMessage(pl, lang.OpenActivityCycleTypeNotEqual)
		return
	}

	needGoldNum := openTemp.Value2
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取每日单笔充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupTemp := groupInterface.(*cyclechargesinglemaxrewmultipletemplate.GroupTemplateCycleSingleMaxRewMultiple)
	descTempList := groupTemp.GetCurDayTempDescList(info.CycleDay)
	for _, temp := range descTempList {
		canRewGold := temp.Value2
		if canRewGold != needGoldNum {
			continue
		}

		//领取条件
		if !info.IsCanReceiveRewards(canRewGold) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"needGoldNum":  needGoldNum,
					"curMaxCharge": info.MaxSingleChargeNum,
				}).Warn("welfare:领取每日单笔充值奖励请求，不满足领取条件")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
			return
		}
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddReceiveRecord(needGoldNum)
	welfareManager.UpdateObj(obj)

	//公告
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		acName := chatlogic.FormatModuleNameNoticeStr(openTemp.Label)
		glodNum := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGoldNum))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleChargeRewardsNotice), plName, acName, glodNum, itemNameLinkStr, link)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRewCycleSingleMaxRewMultiple(rewId, groupId, totalRewData, rewItemMap, info.NewReceiveRewRecord)
	pl.SendMsg(scMsg)

	return
}
