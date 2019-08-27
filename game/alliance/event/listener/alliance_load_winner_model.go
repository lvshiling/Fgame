package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	gameevent "fgame/fgame/game/event"
)

//服务器启动加载城战获胜仙盟盟主雕像
func allianceLoadWinnerModel(target event.EventTarget, data event.EventData) (err error) {
	// 获胜仙盟
	winnerAllianceId := alliance.GetAllianceService().GetAllianceHegemon().GetAllianceId()
	if winnerAllianceId == 0 {
		return
	}

	al := alliance.GetAllianceService().GetAlliance(winnerAllianceId)
	alliancelogic.LoadWinnerModel(al.GetAllianceMengZhuId())
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceLoadWinnerModel, event.EventListenerFunc(allianceLoadWinnerModel))
}
