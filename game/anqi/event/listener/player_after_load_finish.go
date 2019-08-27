package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	anqiManager := p.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiObj := anqiManager.GetAnqiInfo()

	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeAnQi) {
		return
	}

	//暗器信息
	scAnqiGet := pbutil.BuildSCAnqiInfo(p.GetAnqiInfo())
	p.SendMsg(scAnqiGet)

	//加载技能
	temp := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(anqiObj.AdvanceId))
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.Skill)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
