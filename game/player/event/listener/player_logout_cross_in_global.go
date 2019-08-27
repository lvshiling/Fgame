package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func playerLogoutCrossInGlobal(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色全局退出跨服中")

	ctx := scene.WithPlayer(context.Background(), p)
	global.GetGame().GetGlobalRunner().Post(message.NewScheduleMessage(onPlayerLogoutCrossInGlobal, ctx, nil, nil))

	return nil
}

func onPlayerLogoutCrossInGlobal(ctx context.Context, result interface{}, err error) (rerr error) {

	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色全局退出跨服")
	crossType := p.GetCrossType()
	crosslogic.OnPlayerExitCross(p, crossType)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogoutCrossInGlobal, event.EventListenerFunc(playerLogoutCrossInGlobal))
}
