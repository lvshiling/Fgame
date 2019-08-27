package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	juexueeventtypes "fgame/fgame/game/juexue/event/types"
	"fgame/fgame/game/juexue/juexue"
	playerjuexue "fgame/fgame/game/juexue/player"
	juexuetypes "fgame/fgame/game/juexue/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家绝学升级
func playerJueXueUpgrade(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	juexueType := data.(juexuetypes.JueXueType)

	juexueManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjuexue.PlayerJueXueDataManager)

	curType := juexueManager.GetJueXueUseTyp()
	if curType != juexueType {
		return
	}
	jueXueObj := juexueManager.GetJueXueByTyp(juexueType)

	newSkillId := juexue.GetJueXueService().GetSkillId(jueXueObj.Insight, juexueType, jueXueObj.Level)
	oldSkillId := juexue.GetJueXueService().GetSkillId(jueXueObj.Insight, juexueType, jueXueObj.Level-1)

	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(juexueeventtypes.EventTypeJueXueUpgrade, event.EventListenerFunc(playerJueXueUpgrade))
}
