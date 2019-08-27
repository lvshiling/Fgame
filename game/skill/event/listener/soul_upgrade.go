package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	souleventtypes "fgame/fgame/game/soul/event/types"
	playersoul "fgame/fgame/game/soul/player"
)

//帝魂升级结果
func playerSoulUpgarde(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	eventData, ok := data.(*souleventtypes.SoulUpgradeEventData)
	if !ok {
		return
	}
	soulTag := eventData.GetSoulTag()
	manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := manager.IfSoulTagEmemded(soulTag)
	if !flag {
		return
	}

	// soulObj := manager.GetSoulInfoByTag(soulTag)
	// if soulObj == nil {
	// 	return
	// }
	oldOrder := eventData.GetOldOrder()
	newOrder := eventData.GetNewOrder()
	newSkillId := manager.GetSkillId(soulTag, newOrder)
	oldSkillId := manager.GetSkillId(soulTag, oldOrder)

	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulUpgrade, event.EventListenerFunc(playerSoulUpgarde))
}
