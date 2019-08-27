package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	itemskilleventtypes "fgame/fgame/game/itemskill/event/types"
	playeritemskill "fgame/fgame/game/itemskill/player"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家系统技能升级
func playerItemSkillUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	eventData := data.(*itemskilleventtypes.ItemSkillUpgradeEventData)
	typ := eventData.GetType()
	oldLev := eventData.GetOldLev()
	manager := pl.GetPlayerDataManager(types.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)

	level := manager.GetItemSkillLevelByTyp(typ)
	newSkillId := itemskilltemplate.GetItemSkillTemplateService().GetSkillId(typ, level)
	oldSkillId := itemskilltemplate.GetItemSkillTemplateService().GetSkillId(typ, oldLev)
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(itemskilleventtypes.EventTypeItemSkillUpgrade, event.EventListenerFunc(playerItemSkillUpgrade))
}
