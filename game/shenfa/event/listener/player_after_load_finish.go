package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	shenfaManager := p.GetPlayerDataManager(playertypes.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfObj := shenfaManager.GetShenfaInfo()
	shenFaOtherMap := shenfaManager.GetShenfaOtherMap()
	scShenfaGet := pbutil.BuildSCShenfaGet(shenfObj, shenFaOtherMap)
	p.SendMsg(scShenfaGet)

	//加载技能
	temp := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(int32(shenfObj.AdvanceId))
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.Skill)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
