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
	playerproperty "fgame/fgame/game/property/player"
	alliancechargearenapvpassistreturntemplate "fgame/fgame/game/welfare/alliance/charge_arenapvp_assist_return/template"
	alliancechargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/charge_arenapvp_assist_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeWuLian, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshInfo))
}

func refreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	sendEmail(obj)
	syncTodayCost(obj)
	return
}

// 刷新今日数据
func syncTodayCost(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*alliancechargearenapvpassistreturntypes.FeedbackChargeArenapvpAssistReturnInfo)
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		// 同步今日消费
		pl := obj.GetPlayer()
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
			info.CostNum = propertyManager.GetTodayCostNum()
			welfareManager.UpdateObj(obj)
		}
	}
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
	info := obj.GetActivityData().(*alliancechargearenapvpassistreturntypes.FeedbackChargeArenapvpAssistReturnInfo)
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
	groupTemp := groupInterface.(*alliancechargearenapvpassistreturntemplate.GroupTemplateChargeArenapvpAssistReturn)

	costGold := info.CostNum
	if costGold < 1 {
		return
	}

	// TODO 获得比武大会名次
	rank := info.RankType.GetNumber()
	if rank == 0 {
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
	rewItemMap[constanttypes.BindGoldItem] = returnGoldNum
	emaillogic.AddEmail(pl, mailTitle, mailContent, rewItemMap)

	info.IsEmail = true
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.UpdateObj(obj)
	return
}
