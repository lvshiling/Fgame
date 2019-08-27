package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	itemskilleventtypes "fgame/fgame/game/itemskill/event/types"
	playeritemskill "fgame/fgame/game/itemskill/player"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
	itemskilltypes "fgame/fgame/game/itemskill/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家物品技能激活
func playerItemSkillActive(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	typ := data.(itemskilltypes.ItemSkillType)
	manager := pl.GetPlayerDataManager(types.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	level := manager.GetItemSkillLevelByTyp(typ)
	newSkillId := itemskilltemplate.GetItemSkillTemplateService().GetSkillId(typ, level)
	err = skilllogic.TempSkillChange(pl, 0, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(itemskilleventtypes.EventTypeItemSkillActive, event.EventListenerFunc(playerItemSkillActive))
}
