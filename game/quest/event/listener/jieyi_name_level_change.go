package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

// 玩家威名升级
func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeNameUpLev, event.EventListenerFunc(jieYiNameLevelChange))
}

func jieYiNameLevelChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	nameLev := manager.GetNameLevel()

	questlogic.SetQuestData(pl, questtypes.QuestSubTypeXiongDiWeiMing, 0, nameLev)
	return
}
