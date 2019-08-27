package player

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargesinglemaxrewtemplate "fgame/fgame/game/welfare/feedback/charge_single_max_rew/template"
	feedbackchargesinglemaxrewtypes "fgame/fgame/game/welfare/feedback/charge_single_max_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackSingleMaxRewRefreshInfo))
}

//单笔充值-刷新
func feedbackSingleMaxRewRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	groupId := obj.GetGroupId()
	pl := obj.GetPlayer()
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
	info := obj.GetActivityData().(*feedbackchargesinglemaxrewtypes.FeedbackSingleChargeMaxRewInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": groupId,
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	if info.MaxSingleChargeNum < 1 {
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	// 未领取的奖励
	groupTemp := groupInterface.(*feedbackchargesinglemaxrewtemplate.GroupTemplateSingleChargeMaxRew)
	for _, temp := range groupTemp.GetTempDescList() {
		needNum := temp.Value1
		if utils.ContainInt32(info.CanRewRecord, needNum) {
			title := temp.Label
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", needNum))
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleSingleChargeMailContent), acName, chargeText)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)

			info.AddReceiveRecord(needNum)
		}
	}

	info.IsEmail = true
	welfareManager.UpdateObj(obj)

	return
}
