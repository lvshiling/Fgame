package player

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/utils"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	investnewsevendaytypes "fgame/fgame/game/welfare/invest/new_sevenday/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, playerwelfare.ActivityObjInfoRefreshHandlerFunc(refreshNewInvestInfo))
}

func refreshNewInvestInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	refreshInvestNewSevendayInfoAboutCharge(obj)
	newSevenDayInvestRewEnd(obj)
	return
}

//发未领取奖励邮件
func newSevenDayInvestRewEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		return
	}

	info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)
	if info.IsEmail {
		return
	}

	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()

	title := ""
	rewItemMap := make(map[int32]int32)

	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		investType := temp.Value1
		day := temp.Value2
		if info.IsCanReceiveAboutEmail(investType, day) {
			utils.MergeMap(rewItemMap, temp.GetEmailRewItemMap())
		}

		title = temp.Label
	}

	if len(rewItemMap) != 0 {
		acName := chatlogic.FormatMailKeyWordNoticeStr(title)
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityNewSevenDayInvestEmail), acName)
		emaillogic.AddEmail(pl, acName, econtent, rewItemMap)
		info.UpdateNewSevenDayInvestAboutEmail()
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsEmail = true
	welfareManager.UpdateObj(obj)

	return
}

// 同步充值
func refreshInvestNewSevendayInfoAboutCharge(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//同步今天最大单笔充值
	if !welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
	if diff == 0 {
		pl := obj.GetPlayer()
		info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)
		chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		if info.MaxSingleChargeNum != int32(chargeManager.GetTodayMaxSingleCharge()) {
			info.MaxSingleChargeNum = int32(chargeManager.GetTodayMaxSingleCharge())
			welfareManager.UpdateObj(obj)
		}
	}
	return
}
