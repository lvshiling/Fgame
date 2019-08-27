package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	"fgame/fgame/game/fourgod/pbutil"
	playerfourgod "fgame/fgame/game/fourgod/player"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//玩家进入四神遗迹场景
func fourGodPlayerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	sd, ok := data.(fourgodscene.FourGodWarSceneData)
	if !ok {
		return
	}

	npcMap := sd.GetAllNPCS()
	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	keyNum := manager.GetKeyNum()
	scFourGodGet := pbuitl.BuildSCFourGodGet(keyNum, npcMap)
	pl.SendMsg(scFourGodGet)
	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodPlayerEnter, event.EventListenerFunc(fourGodPlayerEnterScene))
}
