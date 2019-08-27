package listener

import (
	"fgame/fgame/core/event"
	lianyulogic "fgame/fgame/cross/lianyu/logic"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	"fgame/fgame/game/lianyu/lianyu"
	"fgame/fgame/game/lianyu/pbutil"
	lianyuscene "fgame/fgame/game/lianyu/scene"
)

//无间炼狱场景结束
func lianYuSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(lianyuscene.LianYuSceneData)
	if !ok {
		return
	}
	//无间炼狱结果
	allPlayerItemMap := sd.GetItemMap()
	allPlayer := sd.GetScene().GetAllPlayers()
	for _, pl := range allPlayer {
		if pl == nil {
			continue
		}
		playerId := pl.GetId()
		itemMap := allPlayerItemMap[playerId]
		scLianYuResult := pbutil.BuildSCLianYuResult(itemMap)
		pl.SendMsg(scLianYuResult)
	}
	lineList := lianyu.GetLianYuService().GetAllLineUpList()
	lianyulogic.BroadLineYuFinishToLineUp(lineList)
	lianyu.GetLianYuService().LianYuSceneFinish()
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuSceneFinish, event.EventListenerFunc(lianYuSceneFinish))
}
