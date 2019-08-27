package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"

	crosseventtypes "fgame/fgame/game/cross/event/types"
	"fgame/fgame/game/cross/pbutil"

	log "github.com/Sirupsen/logrus"
)

func playerCrossHeartbeat(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家跨服心跳")
	siHeartBeat := pbutil.BuildSIHeartBeat()
	p.SendCrossMsg(siHeartBeat)
	return nil
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossHeartbeat, event.EventListenerFunc(playerCrossHeartbeat))
}
