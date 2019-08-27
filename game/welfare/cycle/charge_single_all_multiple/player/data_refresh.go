package player

import (
	"fgame/fgame/common/lang"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	chargesingleallmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/template"
	chargesingleallmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, playerwelfare.ActivityObjInfoRefreshHandlerFunc(dataRefresh))
}

func dataRefresh(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}

	if !isSame {
		//发未领取奖励邮件
		sendEmailCrossDay(obj)
	}

	//非活动时间内
	if !welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		return
	}
	// 同步今日消费
	pl := obj.GetPlayer()

	// 同步今日消费
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	chargeMap := chargeManager.GetTodayChargeList()
	// chargeMap := make(map[int32]int32)
	// for _, chargeObj := range chargeList {
	// 	chargeMap[chargeObj.GetChargeNum()] += 1
	// }
	info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)
	//同步改变修改数据
	flag := info.SyncCharges(chargeMap)
	if flag {
		welfareManager.UpdateObj(obj)
	}

	return
}

func sendEmailCrossDay(obj *playerwelfare.PlayerOpenActivityObject) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	lastUpdateTime := obj.GetUpdateTime()
	info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)
	rewCycDay := info.CycleDay
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	groupTemp := groupInterface.(*chargesingleallmultipletemplate.GroupTemplateCycleSingleAllMultiple)
	remainTimes := info.GetCanRewRecord()

	tempList := groupTemp.GetCurDayTempDescList(rewCycDay)

	for goldNum, timesNum := range remainTimes {
		for _, temp := range tempList {
			if goldNum < temp.Value2 {
				continue
			}
			title := temp.Label
			chargeNum := goldNum
			acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
			chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", chargeNum))
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleSingleChargeMailContent), acName, chargeText)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataListWithRatio(timesNum), temp.GetExpireType(), temp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)
			break
		}
	}
	info.Receive(remainTimes)

	// maxNeedGold := int32(0)
	// for key, _ := range info.CanRewRecord {
	// 	if key > maxNeedGold {
	// 		maxNeedGold = key
	// 	}
	// }
	//判断是否未领取奖励
	// if maxNeedGold == int32(0) {
	// 	return
	// }

	// temp := groupTemp.GetCurDayChargeNumTemp(rewCycDay, maxNeedGold)
	// if temp == nil {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":    pl.GetId(),
	// 			"maxNeedGold": maxNeedGold,
	// 		}).Warn("welfare:领取每日单笔充值奖励请求，领取模板不存在")
	// }
	// title := temp.Label
	// chargeNum := maxNeedGold
	// times := info.CanRewRecord[maxNeedGold]
	// acName := chatlogic.FormatMailKeyWordNoticeStr(temp.Label)
	// chargeText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf("%d", chargeNum))
	// econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityCycleSingleChargeMailContent), acName, chargeText)
	// newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(temp.GetEmailRewItemDataListWithRatio(times), temp.GetExpireType(), temp.GetExpireTime())
	// emaillogic.AddEmailItemLevel(pl, title, econtent, lastUpdateTime, newItemDataList)

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.SingleChargeRecord = make([]int32, 1)
	info.CanRewRecord = map[int32]int32{}
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	info.NewSingleChargeRecord = make(map[int32]int32)
	info.RewRecord = make(map[int32]int32)
	welfareManager.UpdateObj(obj)
	return
}
