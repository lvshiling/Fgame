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

//穿戴元神金装
func goldEquipPutOn(target event.EventTarget, data event.EventData) error {
	pl := target.(player.Player)
	putOnItemId := data.(int32)

	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	putOnItemTemp := item.GetItemService().GetItem(int(putOnItemId))
	goldEquipTemplate := putOnItemTemp.GetGoldEquipTemplate()
	putOnGroupId := goldEquipTemplate.SuitGroup

	curGroupMap := goldequipManager.GetGoldEquipGroupNum()
	curPutGroupEquipNum := curGroupMap[putOnGroupId]
	oldPutGroupEquipNum := curPutGroupEquipNum - 1
	curPutGroupSkillList := goldEquipTemplate.GetGoldEquipGroupSuitSkill(curPutGroupEquipNum)
	oldPutGroupSkillList := goldEquipTemplate.GetGoldEquipGroupSuitSkill(oldPutGroupEquipNum)

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
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipPutOn, event.EventListenerFunc(goldEquipPutOn))
}
