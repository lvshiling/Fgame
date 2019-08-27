package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	lingyuManager := p.GetPlayerDataManager(playertypes.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuObj := lingyuManager.GetLingyuInfo()
	lingyuOtherMap := lingyuManager.GetLingyuOtherMap()
	scLingyuGet := pbutil.BuildSCLingyuGet(lingyuObj, lingyuOtherMap)
	p.SendMsg(scLingyuGet)

	//加载技能
	skillId := lingyutemplate.GetLingyuTemplateService().GetLingyuSkill(int32(lingyuObj.AdvanceId))
	err = skilllogic.TempSkillChange(p, 0, skillId)
	if err != nil {
		return
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
