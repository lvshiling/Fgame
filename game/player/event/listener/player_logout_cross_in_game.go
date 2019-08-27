package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func playerLogoutCrossInGame(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色场景中退出跨服中")
	ctx := scene.WithPlayer(context.Background(), p)
	p.Post(message.NewScheduleMessage(onPlayerLogoutCrossInGame, ctx, nil, nil))
	return nil
}

func onPlayerLogoutCrossInGame(ctx context.Context, result interface{}, err error) (rerr error) {

	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色场景中退出跨服")
	crossType := p.GetCrossType()
	crosslogic.OnPlayerExitCross(p, crossType)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogoutCrossInGame, event.EventListenerFunc(playerLogoutCrossInGame))
}