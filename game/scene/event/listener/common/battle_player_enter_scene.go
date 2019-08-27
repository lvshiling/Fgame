package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//玩家进入场景
func battlePlayerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)

	//移除称号
	s := p.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().IsHiddenTitle() {
		p.TitleHidden(true)
	} else {
		p.TitleHidden(false)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(battlePlayerEnterScene))
}
