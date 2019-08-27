package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/funcopen/funcopen"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	welfaresceneeventtypes "fgame/fgame/game/welfarescene/event/types"
	"fgame/fgame/game/welfarescene/pbutil"
	welfarescenescene "fgame/fgame/game/welfarescene/scene"
)

func welfareSceneRefresh(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(welfarescenescene.WelfareQiYuBossSceneData)
	if !ok {
		return
	}

	collectNum := sd.GetCollectNum()
	npcMap := sd.GetBossMap()
	tempId := int32(sd.GetWelfareSceneTemp().Id)
	scMsg := pbutil.BuildSCWelfareSceneDataChangedNotice(tempId, npcMap, collectNum)
	sd.GetScene().BroadcastMsg(scMsg)

	funcTemp := funcopen.GetFuncOpenService().GetFuncOpenTemplate(funcopentypes.FuncOpenTypeDuanWuQiYuDao)
	if funcTemp == nil {
		return
	}
	refreshScMsg := pbutil.BuildSCWelfareSceneRefersh(sd.GetGroupId())
	alPlayerList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range alPlayerList {
		if pl.GetLevel() < funcTemp.OpenedLevel {
			continue
		}

		pl.SendMsg(refreshScMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(welfaresceneeventtypes.EventTypeWelfareSceneRefresh, event.EventListenerFunc(welfareSceneRefresh))
}
