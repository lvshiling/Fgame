package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marryscene "fgame/fgame/game/marry/scene"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//婚宴开始玩家进入场景
func playerEnterMarryScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	sd, ok := data.(marryscene.MarrySceneData)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	status := sd.GetStatus()
	if status != marryscene.MarrySceneStatusTypeInit {
		manager.RefreshHeriosm()
	}

	period := sd.GetPeriod()
	playerId, name, spouseId, spouseName := sd.GetBothName()
	heroismList := sd.GetHeroismList()
	scMarryBanquet := pbuitl.BuildSCMarryBanquet(period, playerId, name, spouseId, spouseName, heroismList)
	pl.SendMsg(scMarryBanquet)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerEnterMarryScene, event.EventListenerFunc(playerEnterMarryScene))
}
