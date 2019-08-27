package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	souleventtypes "fgame/fgame/game/soul/event/types"
	playersoul "fgame/fgame/game/soul/player"
	soultypes "fgame/fgame/game/soul/types"
)

//帝魂强化
func soulStrengthen(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	soulTag, ok := data.(soultypes.SoulType)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	soulInfo := manager.GetSoulInfoByTag(soulTag)
	if soulInfo == nil {
		return
	}
	strengthenLevel := soulInfo.StrengthenLevel
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSoulStrengthenLevel, int32(soulTag), strengthenLevel)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulStrengthen, event.EventListenerFunc(soulStrengthen))
}
