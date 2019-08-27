package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	xiantaologic "fgame/fgame/game/xiantao/logic"
)

//采集 采集完成
func playerCollectFinishWith(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*collecteventtypes.CollectFinishWithEventData)
	if !ok {
		return
	}

	n := eventData.GetCollectNpc()
	bioType := n.GetBiologyTemplate().GetBiologyScriptType()
	if bioType != scenetypes.BiologyScriptTypeXianTaoQianNianCollect {
		return
	}

	s := pl.GetScene()
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeXianTaoDaHui {
		return
	}

	ppl, ok := pl.(player.Player)
	if !ok {
		return
	}

	//增加采集次数
	sd := s.SceneDelegate()
	xiantaoSd, ok := sd.(xiantaologic.XianTaoSceneData)
	if !ok {
		return
	}
	xiantaoSd.AddPlayerCollectCount(ppl.GetId())
	xiantaologic.PlayerXianTaoInfoChanged(ppl)

	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinishWith, event.EventListenerFunc(playerCollectFinishWith))
}
