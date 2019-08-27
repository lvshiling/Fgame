package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/dragon/dragon"
	pbuitl "fgame/fgame/game/dragon/pbutil"
	playerdragon "fgame/fgame/game/dragon/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerDragonAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := p.GetPlayerDataManager(playertypes.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	dragonObj := manager.GetDragon()
	stageId := dragonObj.StageId

	for i := int32(1); i < stageId; i++ {
		to := dragon.GetDragonService().GetDragonTemplate(i)
		if to == nil {
			continue
		}
		if to.DragonSkill != 0 {
			skilllogic.TempSkillChangeNoUpdate(p, 0, to.DragonSkill)
		}
	}
	to := dragon.GetDragonService().GetDragonTemplate(stageId)
	_, eatFull, _ := manager.IfFullAndEatFull()
	if eatFull && to.DragonSkill != 0 {
		skilllogic.TempSkillChangeNoUpdate(p, 0, to.DragonSkill)
	}

	scDragonGet := pbuitl.BuildSCDragonGet(dragonObj)
	p.SendMsg(scDragonGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerDragonAfterLoadFinish))
}
