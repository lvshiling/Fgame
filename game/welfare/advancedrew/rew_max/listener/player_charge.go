package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewrewmaxtypes "fgame/fgame/game/welfare/advancedrew/rew_max/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//充值升阶
func playerChargeAdvancdRewMax(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeAdvancedRew
	subType := welfaretypes.OpenActivityAdvancedRewSubTypeRewMax

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*advancedrewrewmaxtypes.AdvancedRewMaxInfo)
		info.PeriodChargeNum += goldNum
		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityPeriodChargeNotice(groupId, info.PeriodChargeNum)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeAdvancdRewMax))
}
