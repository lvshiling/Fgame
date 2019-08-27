package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	onearenalogic "fgame/fgame/game/onearena/logic"
	"fgame/fgame/game/onearena/onearena"
	"fgame/fgame/game/onearena/pbutil"
	playeronearena "fgame/fgame/game/onearena/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	curOneArenaObj := manager.GetOneArena()
	level, pos := onearena.GetOneArenaService().GetPlayerOneArena(pl.GetId())
	if curOneArenaObj.Level != level || curOneArenaObj.Pos != pos {
		//玩家抢夺后个人数据没写入
		if level > curOneArenaObj.Level {
			manager.ReplaceOneArena(level, pos)
		} else {
			oneArenaObj := onearena.GetOneArenaService().GetOneArena(curOneArenaObj.Level, curOneArenaObj.Pos)
			manager.ReplaceOneArenaAfter(level, pos, oneArenaObj.OwnerName, oneArenaObj.UpdateTime)
			firstRecord := manager.GetFirstRecord()
			//不是合服
			if oneArenaObj.OwnerName != "" {
				scOneArenaRobbedPush := pbutil.BuildSCOneArenaRobbedPush(firstRecord.RobName)
				pl.SendMsg(scOneArenaRobbedPush)
			}
		}
	}

	kunMap := manager.GetOneArenaKunMap()
	if len(kunMap) != 0 {
		onearenalogic.PlayerAddKun(pl, kunMap)
	}
	manager.DeleteKunRecord()
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
