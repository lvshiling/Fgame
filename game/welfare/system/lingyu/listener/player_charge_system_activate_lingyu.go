package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	systemlingyutypes "fgame/fgame/game/welfare/system/lingyu/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeSystemLingYu(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*chargeeventtypes.PlayerChargeSuccessEventData)
	if !ok {
		return
	}
	chargeGold := eventData.GetChargeGold()

	now := global.GetGame().GetTimeService().Now()

	typ := welfaretypes.OpenActivityTypeSystemActivate
	subType := welfaretypes.OpenActivitySystemActivateSubTypeLingYu
	//领域单笔充值
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		// if !welfarelogic.IsOnActivityTime(groupId) {
		// 	continue
		// }
		groupTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
		if len(groupTemp) != 1 {
			panic("领域活动激活模板应该只有一条")
		}
		for _, temp := range groupTemp {

			continuedTime := int64(temp.Value2) * int64(common.DAY)

			obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
			info := obj.GetActivityData().(*systemlingyutypes.SystemLingYuInfo)

			//验证是否在活动时间
			if info.StartTime+continuedTime < now {
				continue
			}

			if info.IsActivate {
				continue
			}

			if info.MaxSingleChargeGold >= chargeGold {
				continue
			}

			info.MaxSingleChargeGold = chargeGold
			welfareManager.UpdateObj(obj)
			scMsg := pbutil.BuildSCMergeActivitySingleChargeNotice(groupId, chargeGold)
			pl.SendMsg(scMsg)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeSuccess, event.EventListenerFunc(playerChargeSystemLingYu))
}
