package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	"fgame/fgame/game/fourgod/fourgod"
	pbuitl "fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//四神遗迹结束
func fourGodSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(fourgodscene.FourGodSubSceneData)
	if !ok {
		return
	}

	fourgod.GetFourGodService().FourGodSceneFinish()

	s := sd.GetScene()
	for _, p := range s.GetAllPlayers() {
		pl := p.(player.Player)
		manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
		exp, itemMap := manager.GetExpAndItemMap()
		scFourGodTotal := pbuitl.BuildSCFourGodTotal(exp, itemMap)
		pl.SendMsg(scFourGodTotal)
	}

	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodSceneFinish, event.EventListenerFunc(fourGodSceneFinish))
}
