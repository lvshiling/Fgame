package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/arenapvp/arenapvp"
	arenapvplogic "fgame/fgame/game/arenapvp/logic"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//加载完成后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	resultGuess(pl)
	return
}

// 结算竞猜
func resultGuess(pl player.Player) {
	attendList := arenapvp.GetArenapvpService().GetPlayerGuessRecordList(pl.GetId())
	for _, attendObj := range attendList {
		if attendObj.GetStatus() == arenapvptypes.ArenapvpGuessStateInit {
			continue
		}
		if attendObj.GetStatus() == arenapvptypes.ArenapvpGuessStateResult {
			arenapvplogic.GuessResult(pl, attendObj)
		}
		if attendObj.GetStatus() == arenapvptypes.ArenapvpGuessStateReturn {
			arenapvplogic.GuessReturn(pl, attendObj)
		}
		arenapvp.GetArenapvpService().RemovePlayerGuessRecord(pl.GetId(), attendObj.GetRaceNumber(), attendObj.GetGuessType())
	}
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
