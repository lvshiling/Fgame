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
	feedbackchargesinglemaxrewtemplate "fgame/fgame/game/welfare/feedback/charge_single_max_rew/template"
	feedbackchargesinglemaxrewtypes "fgame/fgame/game/welfare/feedback/charge_single_max_rew/types"
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
	welfare.RegisterReceiveHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, welfare.ReceiveHandlerFunc(receiveSingleChargeMaxRew))
}

//处理单笔充值
func receiveSingleChargeMaxRew(pl player.Player, rewId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew
	openTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplate(rewId)
	if openTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"rewId":    rewId,
			}).Warn("welfare:单笔充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
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
			}).Warn("welfare:单笔充值奖励请求，活动不存在")
		return
	}
	info := obj.GetActivityData().(*feedbackchargesinglemaxrewtypes.FeedbackSingleChargeMaxRewInfo)
	needGold := openTemp.Value1
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"groupId":  groupId,
			}).Warn("welfare:领取单笔充值奖励请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	groupTemp := groupInterface.(*feedbackchargesinglemaxrewtemplate.GroupTemplateSingleChargeMaxRew)
	for _, temp := range groupTemp.GetTempDescList() {
		canRewGold := temp.Value1
		if canRewGold != needGold {
			continue
		}

		//领取条件
		if !info.IsCanReceiveRewards(canRewGold) {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"groupId":      groupId,
					"rewId":        rewId,
					"canRewGold":   canRewGold,
					"CanRewRecord": info.CanRewRecord,
				}).Warn("welfare:单笔充值奖励请求，不满足领取条件")
			playerlogic.SendSystemMessage(pl, lang.OpenActivityNotCanReceiveRewards)
			return
		}
	}

	totalRewData, rewItemMap, flag := welfarelogic.AddOpenActivityRewards(pl, openTemp)
	if !flag {
		return
	}

	//更新信息
	info.AddReceiveRecord(needGold)
	welfareManager.UpdateObj(obj)

	//公告
	itemNameLinkStr := welfarelogic.RewardsItemNoticeStr(rewItemMap)
	if len(itemNameLinkStr) > 0 {
		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		plName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
		glodNum := chatlogic.FormatModuleNameNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGold))
		args := []int64{int64(chattypes.ChatLinkTypeOpenView), int64(timeTemp.OpenId)}
		link := coreutils.FormatLink(chattypes.ButtonTypeToGet, args)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MergeActivitySingleChargeRewardsNotice), plName, glodNum, itemNameLinkStr, link)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}

	//同步资源
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCOpenActivityReceiveRew(rewId, groupId, totalRewData, rewItemMap, info.ReceiveRewRecord)
	pl.SendMsg(scMsg)

	return
}
