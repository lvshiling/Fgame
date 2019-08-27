package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	chatlogic "fgame/fgame/game/chat/logic"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardschargetemplate "fgame/fgame/game/welfare/rewards/charge/template"
	rewardschargetypes "fgame/fgame/game/welfare/rewards/charge/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家充值领奖（每多少领奖）
func playerChargeRewards(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeRewards
	subType := welfaretypes.OpenActivityRewardsSubTypeCharge
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	chargeTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range chargeTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		welfareManager.RefreshActivityDataByGroupId(groupId)
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)

		// 第一天走refresh，同步今日累计充值
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			continue
		}

		info := obj.GetActivityData().(*rewardschargetypes.ChargeRewInfo)
		info.GoldNum += goldNum

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*rewardschargetemplate.GroupTemplateRewardsCharge)
		convertRate := groupTemp.GetChargeRewardsConvertRate()
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
			econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.EmailOpenActivityChargeRewardsContent), acName, goldNumText, convertRate)
			newItemDataList := welfarelogic.ConvertToItemDataWithWelfareData(newRewItemDataList, openTemp.GetExpireType(), openTemp.GetExpireTime())
			emaillogic.AddEmailItemLevel(pl, title, econtent, endTime, newItemDataList)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeRewards))
}
