package listener

import (
	"fgame/fgame/core/event"
	playerdragon "fgame/fgame/game/dragon/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"

	"fgame/fgame/game/dragon/dragon"
	dragoneventtypes "fgame/fgame/game/dragon/event/types"
)

//玩家神龙进阶
func playerDragonAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	oldSkillId := int32(0)
	newSkillId := int32(0)

	manager := pl.GetPlayerDataManager(types.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	dragonObj := manager.GetDragon()
	beforeStage := dragonObj.StageId - 1

	beTo := dragon.GetDragonService().GetDragonTemplate(beforeStage)
	if beTo != nil {
		newSkillId = beTo.DragonSkill
	}

	skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)

	return
}

func init() {
	gameevent.AddEventListener(dragoneventtypes.EventTypeDragonAdvanced, event.EventListenerFunc(playerDragonAdvanced))
}
