package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	shihunfanManager := p.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanObj := shihunfanManager.GetShiHunFanInfo()

	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShiHunFan) {
		return
	}

	//噬魂幡信息
	scShiHunFanGet := pbutil.BuildSCShiHunFanGet(shihunfanObj)
	p.SendMsg(scShiHunFanGet)

	//加载技能
	temp := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(int32(shihunfanObj.AdvanceId))
	if temp == nil {
		return
	}
	skilllogic.TempSkillChange(p, 0, temp.SkillId)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
