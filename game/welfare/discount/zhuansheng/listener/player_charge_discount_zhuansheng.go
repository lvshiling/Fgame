package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeZhuanSheng(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeDiscount
	subType := welfaretypes.OpenActivityDiscountSubTypeZhuanSheng
	//转生礼包
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	drewTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range drewTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
		info.ChargeNum += int64(goldNum)
		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, info.ChargeNum)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeZhuanSheng))
}
