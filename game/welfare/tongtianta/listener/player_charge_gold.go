package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	tongtiantatypes "fgame/fgame/game/welfare/tongtianta/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeDrew))
}

//玩家充值元宝
func playerChargeDrew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeTongTianTa
	//遍历所有类型
	for subType := welfaretypes.MinTongTianTaSubType; subType <= welfaretypes.MaxTongTianTaSubType; subType++ {
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		drewTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
		for _, timeTemp := range drewTimeTempList {
			groupId := timeTemp.Group
			if !welfarelogic.IsOnActivityTime(groupId) {
				continue
			}

			obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
			now := global.GetGame().GetTimeService().Now()
			info := obj.GetActivityData().(*tongtiantatypes.TongTianTaInfo)
			diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
			// 第一天走refresh同步今日
			if diff != 0 {
				info.ChargeNum += goldNum
				welfareManager.UpdateObj(obj)
			}

			chargeNum := info.ChargeNum
			scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, int64(chargeNum))
			pl.SendMsg(scMsg)

		}
	}

	return
}
