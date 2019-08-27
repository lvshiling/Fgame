package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/gem/pbutil"
	playergem "fgame/fgame/game/gem/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerGemMineAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	gemManager := p.GetPlayerDataManager(playertypes.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)
	mine := gemManager.GetMine()
	scGemMineGet := pbutil.BuildSCGemMineGet(mine)
	p.SendMsg(scGemMineGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerGemMineAfterLoadFinish))
}
