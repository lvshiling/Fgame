package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家坐骑进阶
func playerMountAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int)
	if !ok {
		return
	}
	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
	if mountTemplate == nil {
		return
	}

	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeMount), int32(advancedId))
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvanced, event.EventListenerFunc(playerMountAdavanced))
}
