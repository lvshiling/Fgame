package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/xinfa/pbutil"
	playerxinfa "fgame/fgame/game/xinfa/player"
	"fgame/fgame/game/xinfa/xinfa"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	xinfaManager := pl.GetPlayerDataManager(playertypes.PlayerXinFaDataManagerType).(*playerxinfa.PlayerXinFaDataManager)
	xinFaMap := xinfaManager.GetXinFaMap()
	scXinFaGet := pbutil.BuildSCXinFaGet(xinFaMap)
	pl.SendMsg(scXinFaGet)

	for _, obj := range xinFaMap {
		typ := obj.Type
		level := obj.Level
		newSkillId := xinfa.GetXinFaService().GetSkillId(typ, level)
		skilllogic.TempSkillChangeNoUpdate(pl, 0, newSkillId)
	}
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
