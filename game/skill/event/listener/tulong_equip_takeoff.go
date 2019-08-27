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

//卸下元神金装
func tuLongEquipTakeOff(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*tulongequipeventtypes.PlayerTuLongEquipChangedEventData)
	if !ok {
		return
	}
	takeOffItemId := eventData.GetItemId()
	suitType := eventData.GetSuitType()

	tulongequipManager := pl.GetPlayerDataManager(types.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	takeOffOnItemTemp := item.GetItemService().GetItem(int(takeOffItemId))
	tuLongEquipTemplate := takeOffOnItemTemp.GetTuLongEquipTemplate()
	takeOffOnGroupId := tuLongEquipTemplate.SuitGroup

	curGroupMap := tulongequipManager.GetTuLongEquipGroupNumByType(suitType)
	curPutGroupEquipNum := curGroupMap[takeOffOnGroupId]
	oldPutGroupEquipNum := curPutGroupEquipNum + 1
	curPutGroupSkillList := tuLongEquipTemplate.GetTuLongEquipGroupSuitSkill(curPutGroupEquipNum)
	oldPutGroupSkillList := tuLongEquipTemplate.GetTuLongEquipGroupSuitSkill(oldPutGroupEquipNum)

	for _, skill := range oldPutGroupSkillList {
		isActive := false
		for _, curSkill := range curPutGroupSkillList {
			if skill == curSkill {
				isActive = true
				break
			}
		}
		if !isActive {
			//卸下技能
			skilllogic.TempSkillChange(pl, skill, 0)
		}
	}

	return nil
}

func init() {
	gameevent.AddEventListener(tulongequipeventtypes.EventTypeTuLongEquipTakeOff, event.EventListenerFunc(tuLongEquipTakeOff))
}
