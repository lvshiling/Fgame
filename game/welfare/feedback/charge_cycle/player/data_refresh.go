package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargecycletemplate "fgame/fgame/game/welfare/feedback/charge_cycle/template"
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCycleCharge, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackCycleRefreshInfo))
}

//连续充值-刷新
func feedbackCycleRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//跨天
	groupId := obj.GetGroupId()
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	if !isSame {
		playerFeedbackCycleCrossDay(obj)
	}

	// 刷新今日数据
	info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		// 同步今日充值
		chargeManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
		todayChargeNum := int32(chargeManager.GetTodayChargeNum())
		if info.CurDayChargeNum != todayChargeNum {
			groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
			if groupInterface != nil {
				groupTemp := groupInterface.(*feedbackchargecycletemplate.GroupTemplateCycleCharge)
				info.CycleDay = welfarelogic.CountCurActivityDay(groupId)
				needGold := groupTemp.GetDayRewCondition(info.CycleDay)

				pl := obj.GetPlayer()
				welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
				info.UpdateTodayCharge(todayChargeNum, needGold)
				welfareManager.UpdateObj(obj)
			}
		}
	}

	// 活动结束
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}

	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("feedbackGoldPig:活动结束发放邮件,活动未结束")
		return
	}

	if info.IsEndMail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	playerFeedbackCycleEnd(obj)
	return
}

//连续充值跨天
func playerFeedbackCycleCrossDay(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargecycletemplate.GroupTemplateCycleCharge)

	needCharge := groupTemp.GetDayRewCondition(info.CycleDay)
	if info.IsCanReceiveToday(needCharge) {
		rewTemp := groupTemp.GetCrossDayRewTemp(info.CycleDay)
		if rewTemp == nil {
			return
		}

		title := rewTemp.Label
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackCycleChargeDuanWuContent), title)
		lastUpdateTime := obj.GetUpdateTime()
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(rewTemp.GetEmailRewItemDataList(), rewTemp.GetExpireType(), rewTemp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.CurDayChargeNum = 0
	info.IsReceiveDayRew = false
	welfareManager.UpdateObj(obj)
	return
}

//连续充值结束
func playerFeedbackCycleEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargecycletemplate.GroupTemplateCycleCharge)
	for needDay, temp := range groupTemp.GetEndRewTempMap() {
		if !info.IsCanReceiveCountDay(needDay) {
			continue
		}

		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackCycleChargeEndMailContent), acName)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
		info.ReceiveCountDay(needDay)
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEndMail = true
	welfareManager.UpdateObj(obj)

	return
}
