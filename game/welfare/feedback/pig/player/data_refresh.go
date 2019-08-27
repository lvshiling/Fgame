package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackpigtemplate "fgame/fgame/game/welfare/feedback/pig/template"
	feedbackpigtypes "fgame/fgame/game/welfare/feedback/pig/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldPig, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackGoldPigRefreshInfo))
}

//养金猪刷新
func feedbackGoldPigRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("feedbackGoldPig:活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*feedbackpigtypes.FeedbackGoldPigInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("feedbackGoldPig:活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	feedbackGoldPigEnd(obj)

	return
}

//养金猪结束
func feedbackGoldPigEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*feedbackpigtypes.FeedbackGoldPigInfo)
	costGoldNum := info.CostGold
	if costGoldNum == 0 {
		info.IsEmail = true
	}

	if !info.IsEmail {
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			return
		}
		groupTemp := groupInterface.(*feedbackpigtemplate.GroupTemplateFeedbackGoldPig)
		feedbackRate := groupTemp.GetGoldPigReturnRate(info.CurCondition)
		feedbackNum := int32(math.Ceil(float64(costGoldNum) * float64(feedbackRate) / float64(common.MAX_RATE)))
		rewMap := make(map[int32]int32)
		rewMap[constanttypes.BindGoldItem] = feedbackNum

		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		title := timeTemp.Name
		acName := chatlogic.FormatMailKeyWordNoticeStr(timeTemp.Name)
		costGold := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), costGoldNum))
		chargeGold := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), info.ChargeGold))
		feebackRate := fmt.Sprintf("%d%s", feedbackRate/100, "%")
		feedbackGold := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d绑元", feedbackNum))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityMailFeedbackGoldPigContent), acName, chargeGold, feebackRate, costGold, feedbackGold)
		emaillogic.AddEmailDefinTime(pl, title, econtent, endTime, rewMap)
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
