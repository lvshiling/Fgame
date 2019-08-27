package listener

import (
	"fgame/fgame/core/event"
	playereventtypes "fgame/fgame/cross/player/event/types"
	"fgame/fgame/cross/teamcopy/teamcopy"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//跨服登陆
func crossPlayerLogout(target event.EventTarget, data event.EventData) (err error) {
	//TODO 修改
	pl := target.(scene.Player)

	//成员
	teamcopy.GetTeamCopyService().TeamCopyMemberOffline(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypeCrossPlayerLogout, event.EventListenerFunc(crossPlayerLogout))
}
