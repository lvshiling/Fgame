package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	playerxinfa "fgame/fgame/game/xinfa/player"
	xinfatypes "fgame/fgame/game/xinfa/types"

	xinafaeventtypes "fgame/fgame/game/xinfa/event/types"
	"fgame/fgame/game/xinfa/xinfa"
)

//玩家心法升级
func playerXinFaUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	xinfaType := data.(xinfatypes.XinFaType)
	manager := pl.GetPlayerDataManager(types.PlayerXinFaDataManagerType).(*playerxinfa.PlayerXinFaDataManager)
	level := manager.GetXinFaLevelByTyp(xinfaType)
	newSkillId := xinfa.GetXinFaService().GetSkillId(xinfaType, level)
	oldSkillId := xinfa.GetXinFaService().GetSkillId(xinfaType, level-1)
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(xinafaeventtypes.EventTypeXinFaUpgrade, event.EventListenerFunc(playerXinFaUpgrade))
}
