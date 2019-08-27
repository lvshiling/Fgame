package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	systemlingyutypes "fgame/fgame/game/welfare/system/lingyu/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家功能开启
func playerFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	funcType := data.(funcopentypes.FuncOpenType)

	if funcType != funcopentypes.FuncOpenTypeLingYu {
		return
	}

	handleFuncOpen(pl)
	return
}

func handleFuncOpen(pl player.Player) {
	now := global.GetGame().GetTimeService().Now()
	typ := welfaretypes.OpenActivityTypeSystemActivate
	subType := welfaretypes.OpenActivitySystemActivateSubTypeLingYu
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	lingyuTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range lingyuTimeTempList {
		groupId := timeTemp.Group
		// 根据策划要求不去根据活动时间模板来
		// if !welfarelogic.IsOnActivityTime(groupId) {
		// 	continue
		// }

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*systemlingyutypes.SystemLingYuInfo)
		if !info.IsOpen {
			info.StartTime = now
			info.IsOpen = true

			welfareManager.UpdateObj(obj)
		}

	}
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpen))
}
