package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargetemplate "fgame/fgame/game/welfare/feedback/charge/template"
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCharge, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackChargeRefreshInfo))
}

//充值返利-刷新
func feedbackChargeRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {

	// 同步今日充值
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			pl := obj.GetPlayer()
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
			if info.GoldNum != int32(chargeManager.GetTodayChargeNum()) {
				info.GoldNum = int32(chargeManager.GetTodayChargeNum())
				welfareManager.UpdateObj(obj)
			}
		}
	}

	//发送未领取奖励邮件
	feedbackChargeEnd(obj)
	return
}

//合服-累充返利结束
func feedbackChargeEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
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
	info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
	if info.IsEmail {
		log.WithFields(
			log.Fields{
				"groupId": obj.GetGroupId(),
			}).Debugf("活动结束发放邮件,邮件已发送")
		return
	}

	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp, ok := groupInterface.(*feedbackchargetemplate.GroupTemplateCharge)
	if !ok {
		log.WithFields(
			log.Fields{
				"type":    obj.GetActivityType(),
				"subType": obj.GetActivitySubType(),
				"groupId": groupId,
			}).Error("welfare:模板类型强转错误")
		return
	}

	chargeGold := info.GoldNum
	if info.GoldNum < 1 {
		return
	}

	for _, temp := range groupTemp.GetOpenTempMap() {
		needCharge := temp.Value1
		if !info.IsCanReceiveRewards(needCharge) {
			continue
		}

		needGoldNum := temp.Value1
		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), chargeGold))
		levelText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), needGoldNum))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityFeedbackChargeContent), acName, chargeText, levelText)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)

		info.AddRecord(needGoldNum)
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)

	return
}
