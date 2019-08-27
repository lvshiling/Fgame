package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	playergoldequip "fgame/fgame/game/goldequip/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//卸下元神金装
func goldEquipTakeOff(target event.EventTarget, data event.EventData) error {
	pl := target.(player.Player)
	takeOffItemId := data.(int32)

	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	takeOffOnItemTemp := item.GetItemService().GetItem(int(takeOffItemId))
	goldEquipTemplate := takeOffOnItemTemp.GetGoldEquipTemplate()
	takeOffOnGroupId := goldEquipTemplate.SuitGroup

	curGroupMap := goldequipManager.GetGoldEquipGroupNum()
	curPutGroupEquipNum := curGroupMap[takeOffOnGroupId]
	oldPutGroupEquipNum := curPutGroupEquipNum + 1
	curPutGroupSkillList := goldEquipTemplate.GetGoldEquipGroupSuitSkill(curPutGroupEquipNum)
	oldPutGroupSkillList := goldEquipTemplate.GetGoldEquipGroupSuitSkill(oldPutGroupEquipNum)

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
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipTakeOff, event.EventListenerFunc(goldEquipTakeOff))
}
