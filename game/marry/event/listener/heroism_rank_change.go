package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	pbuitl "fgame/fgame/game/marry/pbutil"
	marryscene "fgame/fgame/game/marry/scene"
	"fgame/fgame/game/player"
)

//豪气值排行榜变化
func heroismRankChange(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(marryscene.MarrySceneData)
	if !ok {
		return
	}
	heroismList := sd.GetHeroismList()
	scMarryHeroismTopThree := pbuitl.BuildSCMarryHeroismTopThree(heroismList)

	sceneAllPlayer := sd.GetScene().GetAllPlayers()
	for _, spl := range sceneAllPlayer {
		pl, ok := spl.(player.Player)
		if !ok {
			continue
		}
		pl.SendMsg(scMarryHeroismTopThree)
	}
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryHeriosmChange, event.EventListenerFunc(heroismRankChange))
}
