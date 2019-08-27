package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	juexueeventtypes "fgame/fgame/game/juexue/event/types"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家绝学卸下
func playerJueXueUnload(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	oldSkillId := data.(int32)

	err = skilllogic.TempSkillChange(pl, oldSkillId, 0)
	return
}

func init() {
	gameevent.AddEventListener(juexueeventtypes.EventTypeJueXueUnload, event.EventListenerFunc(playerJueXueUnload))
}
