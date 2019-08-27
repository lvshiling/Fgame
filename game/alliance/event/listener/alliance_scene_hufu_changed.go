package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	// alliancescene "fgame/fgame/game/alliance/scene"
	// playerpropertytypes "fgame/fgame/game/property/player/types"
	gameevent "fgame/fgame/game/event"
)

//攻方虎符改变
func allianceSceneHuFuChanged(target event.EventTarget, data event.EventData) (err error) {
	// sd := target.(alliancescene.AllianceSceneData)
	// eventData := data.(*allianceeventtypes.AllianceSceneHuFuChangedEventData)
	// s := sd.GetScene()
	// allianceId := eventData.GetAllianceId()
	// huFu := eventData.GetHuFu()
	// scAllianceSceneHuFuChanged := pbutil.BuildSCAllianceSceneHuFuChanged(huFu)
	// for _, pl := range s.GetAllPlayers() {
	// 	if pl.GetAllianceId() == allianceId {
	// 		tpl := pl.(player.Player)
	// 		tpl.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeHuFu.Mask())
	// 		pl.SendMsg(scAllianceSceneHuFuChanged)
	// 	}
	// }
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneHuFuChanged, event.EventListenerFunc(allianceSceneHuFuChanged))
}
