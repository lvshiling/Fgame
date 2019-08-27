package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	loginlogic "fgame/fgame/cross/login/logic"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"

	playereventtypes "fgame/fgame/cross/player/event/types"

	log "github.com/Sirupsen/logrus"
)

func playerBeforeLogout(target event.EventTarget, data event.EventData) (err error) {
	p := target.(*player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家登出前")
	//退出场景
	ctx := scene.WithPlayer(context.Background(), p)
	global.GetGame().GetGlobalRunner().Post(message.NewScheduleMessage(onPlayerBeforeLogout, ctx, nil, nil))
	return nil
}

func onPlayerBeforeLogout(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(*player.Player)
	loginlogic.Logout(p)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypeCrossPlayerBeforeLogout, event.EventListenerFunc(playerBeforeLogout))
}
