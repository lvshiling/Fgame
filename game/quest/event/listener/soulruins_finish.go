package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	soulruinseventtypes "fgame/fgame/game/soulruins/event/types"
)

//帝陵遗迹完成
func soulRuinsFinsh(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*soulruinseventtypes.SoulRuinsFinishEventData)
	if !ok {
		return
	}
	num := eventData.GetNum()
	soulRuinsId := eventData.GetSoulRuinsId()
	if num <= 0 {
		return
	}

	err = challengeSoulRuinsSucess(pl, num)
	if err != nil {
		return
	}

	err = challengeSpecialSoulRuinsSucess(pl, soulRuinsId, num)
	if err != nil {
		return
	}
	return
}

//帝魂遗迹副本通关X次
func challengeSoulRuinsSucess(pl player.Player, num int32) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSoulRuins, 0, num)
}

//通关指定帝魂副本X次
func challengeSpecialSoulRuinsSucess(pl player.Player, soulRuinsId int32, num int32) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSpecifiedSoulRuins, soulRuinsId, num)
}

func init() {
	gameevent.AddEventListener(soulruinseventtypes.EventTypeSoulruinsFinish, event.EventListenerFunc(soulRuinsFinsh))
}
