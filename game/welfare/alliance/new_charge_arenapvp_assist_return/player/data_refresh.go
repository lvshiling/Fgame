package player

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	alliancenewchargearenapvpassistreturntemplate "fgame/fgame/game/welfare/alliance/new_charge_arenapvp_assist_return/template"
	alliancenewchargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/new_charge_arenapvp_assist_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeNewWuLian, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshInfo))
}

func refreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	sendEmail(obj)
	return
}

func sendEmail(obj *playerwelfare.PlayerOpenActivityObject) {
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
	info := obj.GetActivityData().(*alliancenewchargearenapvpassistreturntypes.FeedbackNewChargeArenapvpAssistReturnInfo)
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
	groupTemp := groupInterface.(*alliancenewchargearenapvpassistreturntemplate.GroupTemplateChargeNewArenapvpAssistReturn)

	rank := info.RankType.GetNumber()
	if rank == 0 {
		return
	}
	costGold := info.CostNum
	if costGold <= 0 {
		return
	}
	returnGoldNum, rateShow, maxGoldNum := groupTemp.GetReturnGoldNum(rank, costGold)

	mailTitle := groupTemp.GetActivityName()
	activityName := chatlogic.FormatMailKeyWordNoticeStr(groupTemp.GetActivityName())
	levelText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), costGold))
	rankText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(info.RankType.GetRankLangCode())))
	rateText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonRate), rateShow))
	extralText := coreutils.FormatColor(chattypes.ColorTypeEmailRedWord, fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackChargeArenapvpAssistReturnContentExtral), maxGoldNum))
	mailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackChargeArenapvpAssistReturnContent), activityName, rankText, rateText, levelText, extralText)
	rewItemMap := make(map[int32]int32)
	rewItemMap[constanttypes.GoldItem] = returnGoldNum
	emaillogic.AddEmail(pl, mailTitle, mailContent, rewItemMap)

	info.IsEmail = true
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.UpdateObj(obj)
	return
}
