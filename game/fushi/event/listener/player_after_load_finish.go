package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/fushi/pbutil"
	playerfushi "fgame/fgame/game/fushi/player"
	fushitemplate "fgame/fgame/game/fushi/template"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}

func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	// 八卦符石信息
	fushiManager := pl.GetPlayerDataManager(playertypes.PlayerFuShiDataManagerType).(*playerfushi.PlayerFuShiDataManager)
	fushiInfo := fushiManager.GetFushiInfo()
	scMsg := pbutil.BuildSCFuShiInfo(fushiInfo)
	pl.SendMsg(scMsg)

	// 加载技能
	for typ, info := range fushiInfo {
		level := info.GetFushiLevel()
		temp := fushitemplate.GetFuShiTemplateService().GetFuShiLevelByFuShiTypeAndLevel(typ, level)
		if temp == nil {
			continue
		}
		//TODO:修改
		skilllogic.TempSkillChangeNoUpdate(pl, 0, temp.SkillId)
	}

	return
}
