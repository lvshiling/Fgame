package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackcosttemplate "fgame/fgame/game/welfare/feedback/cost/template"
	feedbackcosttypes "fgame/fgame/game/welfare/feedback/cost/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCost, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackCostRefreshInfo))
}

//消费返利-刷新
func feedbackCostRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*feedbackcosttypes.FeedbackCostInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	costGold := info.GoldNum
	if costGold < 1 {
		return
	}

	groupTemp := groupInterface.(*feedbackcosttemplate.GroupTemplateCost)
	for _, temp := range groupTemp.GetOpenTempMap() {
		needCost := temp.Value1
		if !info.IsCanReceiveRewards(needCost) {
			continue
		}

		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		costText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), costGold))
		levelText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needCost))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackCostContent), acName, costText, levelText)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(obj.GetPlayer(), title, econtent, endTime, newItemDataList)

		info.AddRecord(needCost)
	}

	pl := obj.GetPlayer()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
