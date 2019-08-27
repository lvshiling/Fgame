package listener

import (
	"fgame/fgame/core/event"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	playerbagua "fgame/fgame/game/bagua/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//八卦秘境挑战结果
func baGuaChallengeResult(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	sucessful, ok := data.(bool)
	if !ok {
		return
	}
	if !sucessful {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerBaGuaDataManagerType).(*playerbagua.PlayerBaGuaDataManager)
	level := manager.GetLevel()

	err = baGuaReachLevel(pl, level)
	if err != nil {
		return
	}
	return
}

//通关八卦秘境第x层
func baGuaReachLevel(pl player.Player, level int32) (err error) {
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeBaGuaMiJingLevel, 0, level)
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaResult, event.EventListenerFunc(baGuaChallengeResult))
}
