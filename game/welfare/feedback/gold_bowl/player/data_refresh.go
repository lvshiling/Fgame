package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackgoldbowltemplate "fgame/fgame/game/welfare/feedback/gold_bowl/template"
	feedbackgoldbowltypes "fgame/fgame/game/welfare/feedback/gold_bowl/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldBowl, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackGoldBowlRefreshInfo))
}

//合服聚宝盆-刷新
func feedbackGoldBowlRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,活动未结束")
		return
	}
	info := obj.GetActivityData().(*feedbackgoldbowltypes.FeedbackGoldBowlInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	//发送未领取奖励邮件
	goldBowlEnd(obj)
	return
}

//聚宝盆结束
func goldBowlEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*feedbackgoldbowltypes.FeedbackGoldBowlInfo)
	costGoldNum := info.GoldNum
	if costGoldNum == 0 {
		return
	}

	rewMap := make(map[int32]int32)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackgoldbowltemplate.GroupTemplateGoldBowl)
	feedbackRate := groupTemp.GetGoldBowlRate()
	feedbackNum := int32(math.Ceil(float64(costGoldNum) * float64(feedbackRate) / float64(common.MAX_RATE)))
	rewMap[constanttypes.BindGoldItem] = feedbackNum

	timeTmep := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	title := timeTmep.Name
	acName := chatlogic.FormatMailKeyWordNoticeStr(timeTmep.Name)
	costGold := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), costGoldNum))
	feebackRate := fmt.Sprintf("%d%s", feedbackRate/100, "%")
	feedbackGold := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d绑元", feedbackNum))
	econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityGoldBowlContent), acName, costGold, feebackRate, feedbackGold)
	emaillogic.AddEmailDefinTime(pl, title, econtent, endTime, rewMap)

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)
	return
}
