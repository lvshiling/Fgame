package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/juexue/juexue"
	"fgame/fgame/game/juexue/pbutil"
	playerjx "fgame/fgame/game/juexue/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerJueXueAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	jxManager := p.GetPlayerDataManager(playertypes.PlayerJueXueDataManagerType).(*playerjx.PlayerJueXueDataManager)
	useId := jxManager.GetJueXueUseId()
	jueXueMap := jxManager.GetJueXueMap()
	scJueXueGet := pbutil.BuildSCJueXueGet(useId, jueXueMap)
	p.SendMsg(scJueXueGet)

	//发送技能
	oldSkillId := int32(0)
	typ := jxManager.GetJueXueUseTyp()
	jueXueObj := jxManager.GetJueXueByTyp(typ)
	if jueXueObj == nil {
		return
	}
	newSkillId := juexue.GetJueXueService().GetSkillId(jueXueObj.Insight, typ, jueXueObj.Level)
	skilllogic.TempSkillChangeNoUpdate(p, oldSkillId, newSkillId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerJueXueAfterLoadFinish))
}
