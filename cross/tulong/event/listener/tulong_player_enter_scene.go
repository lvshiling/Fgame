package listener

import (
	"fgame/fgame/core/event"
	tulongeventtypes "fgame/fgame/cross/tulong/event/types"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	"fgame/fgame/cross/tulong/tulong"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//玩家进入屠龙场景
func tuLongPlayerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	sd, ok := data.(tulongscene.TuLongSceneData)
	if !ok {
		return
	}

	allianceId := pl.GetAllianceId()
	playerBiaoShi, flag := tulong.GetTuLongService().GetPlayerBornBiaoShi(allianceId)
	if !flag {
		return
	}

	bigEgg := sd.GetBigEgg()
	status := bigEgg.GetStatus()
	biaoShi := bigEgg.GetBornBiaoShi()
	scTuLongBossStatus := pbutil.BuildSCTuLongBossStatus(int32(status), biaoShi)
	pl.SendMsg(scTuLongBossStatus)
	scTuLongAllianceBiaoShi := pbutil.BuildSCTuLongAllianceBiaoShi(playerBiaoShi)
	pl.SendMsg(scTuLongAllianceBiaoShi)
	return
}

func init() {
	gameevent.AddEventListener(tulongeventtypes.EventTypeTuLongPlayerEnter, event.EventListenerFunc(tuLongPlayerEnterScene))
}
