package listener

import (
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	skilllogic "fgame/fgame/game/skill/logic"
)

//宝宝玩具激活
func babyUseToy(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*babyeventtypes.PlayerBabyToyChangedEventData)
	if !ok {
		return
	}
	putOnItemId := eventData.GetItemId()
	curSuitNumMap := eventData.GetSuitNumMap()

	itemTemp := item.GetItemService().GetItem(int(putOnItemId))
	toyTemplate := itemTemp.GetBabyToyTemplate()
	suitGroupId := toyTemplate.SuitGroup

	curSuitNum := curSuitNumMap[suitGroupId]
	oldSuitNum := curSuitNum - 1
	curSuitSkillList := toyTemplate.GetBabyToyGroupSuitSkill(curSuitNum)
	oldSuitSkillList := toyTemplate.GetBabyToyGroupSuitSkill(oldSuitNum)

	for _, skill := range curSuitSkillList {
		isActive := false
		for _, oldSkill := range oldSuitSkillList {
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
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyUseToy, event.EventListenerFunc(babyUseToy))
}
