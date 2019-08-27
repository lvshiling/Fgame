package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	gameevent "fgame/fgame/game/event"
	funcopeneventtypes "fgame/fgame/game/funcopen/event/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//玩家功能开启
func playerFuncOpen(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	funcType := data.(funcopentypes.FuncOpenType)

	if funcType != funcopentypes.FuncOpenTypeAnQi {
		return
	}

	//暗器信息
	scAnqiGet := pbutil.BuildSCAnqiInfo(p.GetAnqiInfo())
	p.SendMsg(scAnqiGet)

	//加载技能
	anqiManager := p.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiObj := anqiManager.GetAnqiInfo()
	temp := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(anqiObj.AdvanceId))
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.Skill)

	return
}

func init() {
	gameevent.AddEventListener(funcopeneventtypes.EventTypeFuncOpen, event.EventListenerFunc(playerFuncOpen))
}
