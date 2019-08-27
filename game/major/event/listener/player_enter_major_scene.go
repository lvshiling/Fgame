package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	"fgame/fgame/game/major/pbutil"
	playermajor "fgame/fgame/game/major/player"
	majortemplate "fgame/fgame/game/major/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//玩家进入双修场景
func playerEnterMajorScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	fubenTemplate, ok := data.(majortemplate.MajorTemplate)
	if !ok {
		return
	}

	majorType := fubenTemplate.GetMajorType()
	manager := pl.GetPlayerDataManager(types.PlayerMajorDataManagerType).(*playermajor.PlayerMajorDataManager)
	num := manager.UseMajorNum(majorType)
	scMajorNum := pbutil.BuildSCMajorNum(int32(majorType), num)
	pl.SendMsg(scMajorNum)
	return
}

func init() {
	gameevent.AddEventListener(majoreventtypes.EventTypePlayerEnterMajorScene, event.EventListenerFunc(playerEnterMajorScene))
}
