package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegescene "fgame/fgame/game/godsiege/scene"
	"fgame/fgame/game/scene/scene"
)

//神兽攻城生物改变
func godSiegeNpcChange(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(godsiegescene.GodSiegeSceneData)
	if !ok {
		return
	}

	npc, ok := data.(scene.NPC)
	if !ok {
		return
	}

	scMsg := pbutil.BuildSCGodSiegeCollectNpcChanged(npc)
	sd.GetScene().BroadcastMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegeCollectNpcChanged, event.EventListenerFunc(godSiegeNpcChange))
}
