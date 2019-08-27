package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	tulongequipeventtypes "fgame/fgame/game/tulongequip/event/types"
	playertulongequip "fgame/fgame/game/tulongequip/player"
)

//穿戴元神金装
func tuLongEquipPutOn(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*tulongequipeventtypes.PlayerTuLongEquipChangedEventData)
	if !ok {
		return
	}
	putOnItemId := eventData.GetItemId()
	suitType := eventData.GetSuitType()

	tulongequipManager := pl.GetPlayerDataManager(types.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	putOnItemTemp := item.GetItemService().GetItem(int(putOnItemId))
	tuLongEquipTemplate := putOnItemTemp.GetTuLongEquipTemplate()
	putOnGroupId := tuLongEquipTemplate.SuitGroup

	curGroupMap := tulongequipManager.GetTuLongEquipGroupNumByType(suitType)
	curPutGroupEquipNum := curGroupMap[putOnGroupId]
	oldPutGroupEquipNum := curPutGroupEquipNum - 1
	curPutGroupSkillList := tuLongEquipTemplate.GetTuLongEquipGroupSuitSkill(curPutGroupEquipNum)
	oldPutGroupSkillList := tuLongEquipTemplate.GetTuLongEquipGroupSuitSkill(oldPutGroupEquipNum)

	for _, skill := range curPutGroupSkillList {
		isActive := false
		for _, oldSkill := range oldPutGroupSkillList {
			if skill == oldSkill {
				isActive = true
				break
			}
		}
		if !isActive {
			//添加技能
			skilllogic.TempSkillChange(pl, 0, skill)
		}
	}

	return nil
}
func init() {
	gameevent.AddEventListener(tulongequipeventtypes.EventTypeTuLongEquipPutOn, event.EventListenerFunc(tuLongEquipPutOn))
}
