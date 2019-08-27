package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	juexueeventtypes "fgame/fgame/game/juexue/event/types"
	"fgame/fgame/game/juexue/juexue"
	playerjuexue "fgame/fgame/game/juexue/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家绝学顿悟
func playerJueXueInsight(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	eventData := data.(*juexueeventtypes.JueXueInsightEventData)
	juexueType := eventData.GetType()
	oldSkillId := eventData.GetOldSkillId()

	juexueManager := pl.GetPlayerDataManager(types.PlayerJueXueDataManagerType).(*playerjuexue.PlayerJueXueDataManager)

	curType := juexueManager.GetJueXueUseTyp()
	if curType != juexueType {
		return
	}
	jueXueObj := juexueManager.GetJueXueByTyp(curType)

	newSkillId := juexue.GetJueXueService().GetSkillId(jueXueObj.Insight, juexueType, jueXueObj.Level)
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(juexueeventtypes.EventTypeJueXueInsight, event.EventListenerFunc(playerJueXueInsight))
}
