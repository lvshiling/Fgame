package player

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesinglemaxrewtemplate "fgame/fgame/game/welfare/cycle/charge_single_max_rew/template"
	cyclechargesinglemaxrewtypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(cycleSingleChargeMaxRewRefreshInfo))
}

//每日单笔充值-刷新
func cycleSingleChargeMaxRewRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//重置
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		//发未领取奖励邮件
		cycleSingleMaxRewCrossDay(obj)
	}

	return
}

//每日单笔充值变更
func cycleSingleMaxRewCrossDay(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	lastUpdateTime := obj.GetUpdateTime()
	info := obj.GetActivityData().(*cyclechargesinglemaxrewtypes.CycleSingleChargeMaxRewInfo)
	rewCycDay := info.CycleDay

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*cyclechargesinglemaxrewtemplate.GroupTemplateCycleSingleMaxRew)
	// 未领取的奖励
	tempList := groupTemp.GetCurDayTempDescList(rewCycDay)
	for _, temp := range tempList {
		needNum := temp.Value2
		if utils.ContainInt32(info.CanRewRecord, needNum) {
			title := temp.Label
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", needNum))
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleSingleChargeMailContent), acName, chargeText)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)

			break
		}
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.MaxSingleChargeNum = 0
	info.ReceiveRewRecord = []int32{}
	info.CanRewRecord = []int32{}
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	welfareManager.UpdateObj(obj)

	return
}
