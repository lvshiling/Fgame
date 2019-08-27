package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
)

//玩家功能开启
func playerFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	funcType := data.(funcopentypes.FuncOpenType)

	if funcType != funcopentypes.FuncOpenTypeTianMo {
		return
	}

	tianMoManager := p.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianMoObj := tianMoManager.GetTianMoInfo()

	//加载技能
	temp := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(tianMoObj.AdvanceId)
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.SkillId)
	return
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpen))
}
