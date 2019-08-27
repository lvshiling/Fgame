package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家功能开启
func playerFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	funcType := data.(funcopentypes.FuncOpenType)

	if funcType != funcopentypes.FuncOpenTypeShiHunFan {
		return
	}

	manager := p.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shiHunFanObj := manager.GetShiHunFanInfo()

	//加载技能
	temp := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(shiHunFanObj.AdvanceId))
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.SkillId)
	return
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpen))
}
