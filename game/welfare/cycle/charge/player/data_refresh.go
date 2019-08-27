package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	cyclechargetemplate "fgame/fgame/game/welfare/cycle/charge/template"
	cyclechargetypes "fgame/fgame/game/welfare/cycle/charge/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeCharge, playerwelfare.ActivityObjInfoRefreshHandlerFunc(cycleChargeRefreshInfo))
}

//每日充值-刷新
func cycleChargeRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()

	//重置
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	if !isSame {
		playerCycleChanged(obj)
	}

	// 同步今日充值
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		pl := obj.GetPlayer()
		chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
		todayChargeNum := int32(chargeManager.GetTodayChargeNum())
		info := obj.GetActivityData().(*cyclechargetypes.CycleChargeInfo)
		if info.GoldNum != todayChargeNum {
			info.GoldNum = todayChargeNum
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			welfareManager.UpdateObj(obj)
		}
	}
	return
}

//每日首充变更
func playerCycleChanged(obj *playerwelfare.PlayerOpenActivityObject) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*cyclechargetypes.CycleChargeInfo)
	chargeGoldNum := info.GoldNum
	rewCycDay := info.CycleDay
	lastUpdateTime := obj.GetUpdateTime()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*cyclechargetemplate.GroupTemplateCycleCharge)
	tempList := groupTemp.GetCurDayTempList(rewCycDay)
	for _, temp := range tempList {
		needCharge := temp.Value2
		if info.IsCanReceiveRewards(needCharge) {
			title := temp.Label
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", chargeGoldNum))
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityCycleChargeContent), acName, chargeText)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataList(), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
		}
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.GoldNum = 0
	info.RewRecord = []int32{}
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	welfareManager.UpdateObj(obj)
	return
}
