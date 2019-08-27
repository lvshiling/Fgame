package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/chess/chess"
	"fgame/fgame/game/chess/pbutil"
	playerchess "fgame/fgame/game/chess/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	chessManager := pl.GetPlayerDataManager(playertypes.PlayerChessDataManagerType).(*playerchess.PlayerChessDataManager)
	playerChessMap := chessManager.GetChessMap()
	logList := chess.GetChessService().GetLogByTime(0)
	if len(logList) > 10 {
		logList = logList[len(logList)-10:]
	}
	scChessInfoGet := pbutil.BuildSCChessInfoGet(playerChessMap, logList)
	pl.SendMsg(scChessInfoGet)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
