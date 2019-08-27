package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	playerrealm "fgame/fgame/game/realm/player"
)

//天劫塔挑战结果
func realmChallengeResult(target event.EventTarget, data event.EventData) (err error) {
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
	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	level := manager.GetTianJieTaLevel()
	isFullLevel := manager.IfFullLevel()

	err = realmReachLevel(pl, level)
	if err != nil {
		return
	}

	err = challengeRealm(pl, isFullLevel)
	if err != nil {
		return
	}

	err = challengeRealmSucess(pl, isFullLevel)
	if err != nil {
		return
	}

	return
}

//通关天劫塔第x层
func realmReachLevel(pl player.Player, level int32) (err error) {
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeRealmLevel, 0, level)
}

//进入天劫塔X次
func challengeRealm(pl player.Player, isFullLevel bool) (err error) {
	if isFullLevel {
		return questlogic.FillQuestData(pl, questtypes.QuestSubTypeEnterRealm, 0)
	}
	return
}

//挑战成功X次天劫塔
func challengeRealmSucess(pl player.Player, isFullLevel bool) (err error) {
	if isFullLevel {
		return questlogic.FillQuestData(pl, questtypes.QuestSubTypeRealm, 0)
	} else {
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeRealm, 0, 1)
	}
	return
}

func init() {
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmResult, event.EventListenerFunc(realmChallengeResult))
}
