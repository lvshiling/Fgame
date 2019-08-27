package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/arenapvp"
	arenapvpeventtypes "fgame/fgame/cross/arenapvp/event/types"
	"fgame/fgame/cross/arenapvp/pbutil"
	arenapvpscene "fgame/fgame/cross/arenapvp/scene"
	gameevent "fgame/fgame/game/event"
)

//竞技场场景结束
func arenapvpBatlleSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(arenapvpscene.ArenapvpBattleSceneData)
	if !ok {
		return
	}

	s := sd.GetScene()
	pvpTemp := sd.GetPvpTemp()
	pvpType := pvpTemp.GetArenapvpType()

	//发送结果
	scMsg := pbutil.BuildSCArenapvpBattleEnd(sd.GetWinnerId(), pvpType)

	battlePl1, battlePl2 := sd.GetBattlePlayer()
	if battlePl1 != nil {
		win := sd.GetWinnerId() == battlePl1.PlayerId
		isMsg := pbutil.BuildISArenapvpResultBattle(win, int32(pvpType))

		pl1 := s.GetPlayer(battlePl1.PlayerId)
		if pl1 != nil {
			pl1.SendMsg(isMsg)
			pl1.SendMsg(scMsg)
		}
	}

	if battlePl2 != nil {
		win := sd.GetWinnerId() == battlePl2.PlayerId
		isMsg := pbutil.BuildISArenapvpResultBattle(win, int32(pvpType))

		pl2 := s.GetPlayer(battlePl2.PlayerId)
		if pl2 != nil {
			pl2.SendMsg(isMsg)
			pl2.SendMsg(scMsg)
		}
	}

	arenapvp.GetArenapvpService().ArenapvpBattleFinish(s)

	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpBattleSceneFinish, event.EventListenerFunc(arenapvpBatlleSceneFinish))
}
