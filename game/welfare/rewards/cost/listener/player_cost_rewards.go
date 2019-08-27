package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardscosttemplate "fgame/fgame/game/welfare/rewards/cost/template"
	rewardscosttypes "fgame/fgame/game/welfare/rewards/cost/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家消费领奖（每多少领奖）
func playerCostRewards(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int64)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeRewards
	subType := welfaretypes.OpenActivityRewardsSubTypeCost
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	costTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range costTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*rewardscosttypes.CostRewInfo)

		now := global.GetGame().GetTimeService().Now()
		isSame, _ := timeutils.IsSameDay(now, obj.GetStartTime())
		//是否第一天
		if isSame {
			propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
			todayCostNum := propertyManager.GetTodayCostNum()
			if info.GoldNum != todayCostNum {
				addNum := todayCostNum - info.GoldNum
				if addNum > 0 {
					info.GoldNum += addNum
					info.LeftConvertNum += addNum
				}
			}
		} else {
			info.GoldNum += goldNum
			info.LeftConvertNum += goldNum
		}

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*rewardscosttemplate.GroupTemplateRewardsCost)
		convertRate := groupTemp.GetCostRewardsConvertRate()
		addTimes := info.CountLeftTimes(convertRate)
		welfareManager.UpdateObj(obj)

		// 发邮件
		if addTimes > 0 {
			openTemp := groupTemp.GetFirstOpenTemp()
			if openTemp == nil {
				continue
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

		// 消费推送
		scMsg := pbutil.BuildSCOpenActivityCostRewardsCostInfo(groupId, info.GoldNum)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCost, event.EventListenerFunc(playerCostRewards))
}
