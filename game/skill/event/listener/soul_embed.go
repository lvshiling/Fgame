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

//帝魂镶嵌结果
func playerSoulEmbed(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	oldSkillId := int32(0)
	newSkillId := int32(0)

	eventData, ok := data.(*souleventtypes.SoulEmbedEventData)
	if !ok {
		return
	}
	oldSoulTag := eventData.GetOldTag()
	newSoulTag := eventData.GetNewTag()
	manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	if oldSoulTag.Valid() {
		soulObj := manager.GetSoulInfoByTag(oldSoulTag)
		if soulObj != nil {
			oldSkillId = manager.GetSkillId(oldSoulTag, soulObj.AwakenOrder)
		}
	}

	soulObj := manager.GetSoulInfoByTag(newSoulTag)
	if soulObj != nil {
		newSkillId = manager.GetSkillId(newSoulTag, soulObj.AwakenOrder)
	}

	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulEmbed, event.EventListenerFunc(playerSoulEmbed))
}
