package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
	"fgame/fgame/game/godsiege/godsiege"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/godsiege/pbutil"
	godsiegescene "fgame/fgame/game/godsiege/scene"
)

//神兽攻城场景结束
func godSiegeSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(godsiegescene.GodSiegeSceneData)
	if !ok {
		return
	}
	godType := sd.GetGodType()
	//神兽攻城结果
	allPlayers := sd.GetScene().GetAllPlayers()
	allPlayerItemMap := sd.GetItemMap()
	for _, pl := range allPlayers {
		if pl == nil {
			continue
		}
		playerId := pl.GetId()
		itemMap := allPlayerItemMap[playerId]
		scGodSiegeResult := pbutil.BuildSCGodSiegeResult(int32(godType), itemMap)
		pl.SendMsg(scGodSiegeResult)
	}

	lineList := godsiege.GetGodSiegeService().GetAllLineUpList(godType)
	godsiegelogic.BroadGodSiegeFinishToLineUp(int32(godType), lineList)
	godsiege.GetGodSiegeService().GodSiegeSceneFinish(godType)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegeSceneFinish, event.EventListenerFunc(godSiegeSceneFinish))
}
