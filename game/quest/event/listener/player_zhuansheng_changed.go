package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	zhuanShu := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeZhuanSheng, 0, zhuanShu)

	// 激活奇遇任务
	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	questlogic.InitQiYuQuest(pl, feiManager.GetFeiShengLevel())
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
