package player

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargereturnmultipletemplate "fgame/fgame/game/welfare/feedback/charge_return_multiple/template"
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	timeutils "fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, playerwelfare.ActivityObjInfoRefreshHandlerFunc(chargeReturnMultipleRefreshInfo))
}

//累充-刷新
func chargeReturnMultipleRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	// 活动是否结束
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("ChargeReturnMultiple:活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	//同步第一天充值
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			pl := obj.GetPlayer()

			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			if info.PeriodChargeNum != int32(chargeManager.GetTodayChargeNum()) {
				diffNum := int32(chargeManager.GetTodayChargeNum()) - info.PeriodChargeNum
				if diffNum > 0 {

					info.AddPeriodCharge(diffNum)
					welfareManager.UpdateObj(obj)
				}
			}
		}
	}

	groupTemp := groupInterface.(*feedbackchargereturnmultipletemplate.GroupTemplateChargeReturnMultiple)

	totalCnt := info.PeriodChargeNum / groupTemp.GetPerChargeNum()
	if totalCnt <= info.RewardCnt {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,奖励已经领完")
		return
	}
	if groupTemp.GetRewardLimitCnt() <= info.RewardCnt {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("ChargeReturnMultiple:活动结束发放邮件,奖励次数上限")
		return
	}

	//发送未领取奖励邮件
	playerChargeReturnMultipleEnd(obj)
	return
}

//期间累计充值奖励结束
func playerChargeReturnMultipleEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargereturnmultipletemplate.GroupTemplateChargeReturnMultiple)
	totalCnt := info.PeriodChargeNum / groupTemp.GetPerChargeNum()
	leftCnt := totalCnt - info.RewardCnt
	rewardLimitCnt := groupTemp.GetRewardLimitCnt()
	if totalCnt > rewardLimitCnt {
		leftCnt = rewardLimitCnt - info.RewardCnt
	}
	if leftCnt <= 0 {
		return
	}

	info.AddRewardCnt(leftCnt)
	reachVal := info.RewardCnt * groupTemp.GetPerChargeNum()
	for _, temp := range groupTemp.GetOpenTempMap() {
		title := coreutils.FormatNoticeStr(temp.Label)
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", reachVal))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackChargeReturnMultipleEndMailContent), acName, chargeText)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataListWithRatio(leftCnt), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.UpdateObj(obj)
	return
}
