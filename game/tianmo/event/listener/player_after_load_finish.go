package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	tianMoManager := p.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianMoObj := tianMoManager.GetTianMoInfo()

	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeTianMo) {
		return
	}

	//天魔信息
	scTianMoGet := pbutil.BuildSCTianMoInfo(p.GetTianMoTiInfo())
	p.SendMsg(scTianMoGet)

	//加载技能
	temp := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(tianMoObj.AdvanceId)
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.SkillId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
