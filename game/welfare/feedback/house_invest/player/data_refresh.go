package player

import (
	"fgame/fgame/common/lang"
	coreutils "fgame/fgame/core/utils"
	playercharge "fgame/fgame/game/charge/player"
	chatlogic "fgame/fgame/game/chat/logic"
	constanttypes "fgame/fgame/game/constant/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbackhouseinvesttemplate "fgame/fgame/game/welfare/feedback/house_invest/template"
	feedbackhouseinvesttypes "fgame/fgame/game/welfare/feedback/house_invest/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseInvest, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackHouseInvestRefreshInfo))
}

//房产投资-刷新
func feedbackHouseInvestRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	//跨天
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(obj.GetUpdateTime(), now)
	if err != nil {
		return err
	}
	if !isSame {
		playerFeedbackHouseInvestCrossDay(obj)
	}

	// 刷新今日数据
	info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		chargeManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
		todayCharge := int32(chargeManager.GetTodayChargeNum())
		if info.CurDayChargeNum != todayCharge {
			// 同步今日充值
			welfareManager := obj.GetPlayer().GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			diff := todayCharge - info.CurDayChargeNum
			if diff > 0 {
				info.ChargeNum += diff
			}
			info.CurDayChargeNum = todayCharge
			welfareManager.UpdateObj(obj)
		}
	}

	// 活动结束
	endTime := obj.GetEndTime()
	if endTime <= 0 {
		return
	}
	if now < endTime {
		return
	}

	if !info.IsActivity {
		return
	}

	if info.IsSell {
		return
	}

	//发送未领取奖励邮件
	playerFeedbackHouseInvestEnd(obj)
	return
}

//房产投资跨天
func playerFeedbackHouseInvestCrossDay(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	if !welfarelogic.IsOnActivityTime(groupId) {
		return
	}

	info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)
	if info.IsSell {
		return
	}

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.CurDayChargeNum = 0
	info.IsCurDayDecor = false
	welfareManager.UpdateObj(obj)

	return
}

//房产投资结束
func playerFeedbackHouseInvestEnd(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	pl := obj.GetPlayer()
	groupId := obj.GetGroupId()
	endTime := obj.GetEndTime()
	info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackhouseinvesttemplate.GroupTemplateHouseInvest)
	openTemp := groupTemp.GetOpenActivityHouseInvest(info.DecorDays)
	if openTemp == nil {
		return
	}

	title := coreutils.FormatNoticeStr(openTemp.Label)
	acName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
	econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityFeedbackHouseInvestEndMailContent), acName)
	//rewItemMap := temp.GetEmailRewItemMap()
	rewItemMap := make(map[int32]int32)
	sellNum := int32(groupTemp.GetOpenActivityHouseInvestSellNum(info.DecorDays))
	rewItemMap[constanttypes.SilverItem] = sellNum
	newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, openTemp.GetExpireType(), openTemp.GetExpireTime())
	emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)

	//重置
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	info.IsSell = true
	welfareManager.UpdateObj(obj)

	return
}
