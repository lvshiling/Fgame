package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	souleventtypes "fgame/fgame/game/soul/event/types"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
)

//帝魂觉醒
func playerSoulAwaken(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	soulTag := data.(soultypes.SoulType)
	manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := manager.IfSoulTagEmemded(soulTag)
	if !flag {
		return
	}

	soulObj := manager.GetSoulInfoByTag(soulTag)
	if soulObj == nil {
		return
	}

	newSkillId := manager.GetSkillId(soulTag, soulObj.AwakenOrder)

	awakenTemplate := soul.GetSoulService().GetSoulAwakenTemplateByOrder(soulTag, soulObj.AwakenOrder)
	oldSkillId := awakenTemplate.UplevelSkillId
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulAwaken, event.EventListenerFunc(playerSoulAwaken))
}
