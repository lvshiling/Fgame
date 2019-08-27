package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questeventtypes "fgame/fgame/game/quest/event/types"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
)

//开服目标任务完成次数变更
func questKaiFuMuBiaoFinishNumChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	kaiFuDayList := data.([]int32)
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	kaiFuMuBiaoMap := manager.GetKaiFuMuBiaoMap()
	scQuestKaiFuMuBiaoFinishNumChanged := pbutil.BuildSCQuestKaiFuMuBiaoFinishNumChanged(kaiFuMuBiaoMap, kaiFuDayList)
	pl.SendMsg(scQuestKaiFuMuBiaoFinishNumChanged)
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestKaiFuMuBiaoFinishChanged, event.EventListenerFunc(questKaiFuMuBiaoFinishNumChanged))
}
