package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesinglemaxrewmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/template"
	cyclechargesinglemaxrewmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple, playerwelfare.ActivityObjInfoRefreshHandlerFunc(cycleSingleChargeMaxRewMultipleRefreshInfo))
}

//每日单笔充值,多次-刷新
func cycleSingleChargeMaxRewMultipleRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//重置
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		//发未领取奖励邮件
		cycleSingleMaxRewMultipleCrossDay(obj)
	}

	return
}

//每日单笔充值邮件变更,多次
func cycleSingleMaxRewMultipleCrossDay(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	lastUpdateTime := obj.GetUpdateTime()
	info := obj.GetActivityData().(*cyclechargesinglemaxrewmultipletypes.CycleSingleChargeMaxRewMultipleInfo)
	rewCycDay := info.CycleDay

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*cyclechargesinglemaxrewmultipletemplate.GroupTemplateCycleSingleMaxRewMultiple)
	// 未领取的奖励
	leftM := info.LeftCanReceiveRewards()
	for chargeNum, times := range leftM {
		temp := groupTemp.GetCurDayChargeNumTemp(rewCycDay, chargeNum)
		if temp == nil {
			continue
		}
		title := temp.Label
		acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
		chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", chargeNum))
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleSingleChargeMailContent), acName, chargeText)
		newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataListWithRatio(times), temp.GetExpireType(), temp.GetExpireTime())
		emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.MaxSingleChargeNum = 0
	info.NewReceiveRewRecord = map[int32]int32{}
	info.NewCanRewRecord = map[int32]int32{}
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	welfareManager.UpdateObj(obj)

	return
}
