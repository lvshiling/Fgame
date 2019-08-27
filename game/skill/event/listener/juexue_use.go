package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	juexueeventtypes "fgame/fgame/game/juexue/event/types"
	"fgame/fgame/game/juexue/juexue"
	playerjuexue "fgame/fgame/game/juexue/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家绝学使用
func playerJueXueUse(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	oldSkillId := data.(int32)

	jxManager := pl.GetPlayerDataManager(playertypes.PlayerJueXueDataManagerType).(*playerjuexue.PlayerJueXueDataManager)
	typ := jxManager.GetJueXueUseTyp()
	jueXueObj := jxManager.GetJueXueByTyp(typ)
	if jueXueObj == nil {
		return
	}
	newSkillId := juexue.GetJueXueService().GetSkillId(jueXueObj.Insight, typ, jueXueObj.Level)
	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(juexueeventtypes.EventTypeJueXueUse, event.EventListenerFunc(playerJueXueUse))
}
