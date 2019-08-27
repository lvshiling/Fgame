package listener

import (
	"fgame/fgame/core/event"
	playereventtypes "fgame/fgame/cross/player/event/types"
	"fgame/fgame/cross/teamcopy/teamcopy"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//跨服登陆
func crossPlayerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	//TODO 修改
	pl := target.(scene.Player)

	sd := teamcopy.GetTeamCopyService().GetTeamCopyDataByPlayerId(pl.GetId())
	if sd == nil {
		return
	}
	sceneData := sd.GetTeamCopySceneData()
	teamObj := sceneData.GetTeamObj()
	teamcopy.GetTeamCopyService().TeamCopyMemberOnline(pl)
	pl.SetArenaTeam(teamObj.GetTeamId(), teamObj.GetTeamName(), teamObj.GetTeamPurpose())
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypeCrossPlayerAfterLoad, event.EventListenerFunc(crossPlayerAfterLoad))
}
