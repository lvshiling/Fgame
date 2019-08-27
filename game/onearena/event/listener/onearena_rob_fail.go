package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	onearenalogic "fgame/fgame/game/onearena/logic"
	playeronearena "fgame/fgame/game/onearena/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//灵池抢夺成功
func oneArenaRobFail(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*onearenaeventtypes.OneArenaRobFailEventData)
	if !ok {
		return
	}
	peerArenaData := eventData.GetPeerOneArenaData()

	playerId := eventData.GetPlayerId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	ownerId := peerArenaData.GetPlayerId()
	level := peerArenaData.GetLevel()
	pos := peerArenaData.GetPos()
	if ownerId != 0 {
		err = onearenalogic.PeerRobbedRecord(peerArenaData, pl.GetName(), false)
		if err != nil {
			return
		}
	}

	manager := pl.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*playeronearena.PlayerOneArenaDataManager)
	//写抢夺记录
	manager.RobOneArenaRecord(level, pos)
	return
}

func init() {
	gameevent.AddEventListener(onearenaeventtypes.EventTypePlayerOneArenaFail, event.EventListenerFunc(oneArenaRobFail))
}
