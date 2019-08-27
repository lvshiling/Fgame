package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/transportation/pbutil"
	"fgame/fgame/game/transportation/transpotation"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeWorld {
		return
	}

	//玩家上线时重新同步下数据
	biaoChe := transpotation.GetTransportService().GetTransportation(pl.GetId())
	if biaoChe == nil {
		return
	}
	obj := biaoChe.GetTransportationObject()

	scTransportBriefInfoNotice := pbutil.BuildSCTransportBriefInfoNotice(obj, biaoChe)
	pl.SendMsg(scTransportBriefInfoNotice)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
