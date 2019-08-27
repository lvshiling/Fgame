package player

import (
	"fgame/fgame/common/lang"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardscosttemplate "fgame/fgame/game/welfare/rewards/cost/template"
	rewardscosttypes "fgame/fgame/game/welfare/rewards/cost/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

// 每消费奖励
func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCost, playerwelfare.ActivityObjInfoRefreshHandlerFunc(rewardsCostRefreshInfo))
}

func rewardsCostRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	info := obj.GetActivityData().(*rewardscosttypes.CostRewInfo)

	pl := obj.GetPlayer()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	todayCostNum := propertyManager.GetTodayCostNum()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(obj.GetGroupId())
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*rewardscosttemplate.GroupTemplateRewardsCost)
	if !welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameDay(now, obj.GetStartTime())
	//是否第一天
	if !isSame {
		return
	}

	convertRate := groupTemp.GetCostRewardsConvertRate()
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	if info.GoldNum != todayCostNum {
		addNum := todayCostNum - info.GoldNum
		if addNum > 0 {
			info.GoldNum += addNum
			info.LeftConvertNum += addNum
			addTimes := info.CountLeftTimes(convertRate)
			welfareManager.UpdateObj(obj)
			// 发邮件
			if addTimes > 0 {
				openTemp := groupTemp.GetFirstOpenTemp()
				if openTemp == nil {
					return
				}

				newRewItemDataList := openTemp.GetEmailRewItemDataListWithRatio(addTimes)
				endTime := global.GetGame().GetTimeService().Now()
				title := openTemp.Label
				acName := chatlogic.FormatMailKeyWordNoticeStr(openTemp.Label)
				goldNumText := chatlogic.FormatMailKeyWordNoticeStr(fmt.Sprintf(lang.GetLangService().ReadLang(lang.OpenActivityEmailCommonGoldString), info.GoldNum))
				econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityCostRewardsContent), acName, goldNumText, convertRate)
				newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(newRewItemDataList, openTemp.GetExpireType(), openTemp.GetExpireTime())
				emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
			}
		}
	}
	return
}
