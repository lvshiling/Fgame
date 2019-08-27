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
	welfaretypes "fgame/fgame/game/welfare/types"
	xiuxianbooktypes "fgame/fgame/game/welfare/xiuxianbook/types"
	"fgame/fgame/pkg/timeutils"
)

//累充奖励
func playerCharge(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeXiuxianBook

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	for subType := welfaretypes.MinOpenActivityXiuxianBookSubType; subType <= welfaretypes.MaxOpenActivityXiuxianBookSubType; subType++ {
		err = welfareManager.RefreshActivityData(typ, subType)
		if err != nil {
			continue
		}
		timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
		for _, timeTemp := range timeTempList {
			groupId := timeTemp.Group
			if !welfarelogic.IsOnActivityTime(groupId) {
				continue
			}
			obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
			now := global.GetGame().GetTimeService().Now()
			info := obj.GetActivityData().(*xiuxianbooktypes.XiuxianBookInfo)
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

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerCharge))
}
