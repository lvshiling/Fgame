package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	loginlogic "fgame/fgame/cross/login/logic"
	playereventtypes "fgame/fgame/cross/player/event/types"
	"fgame/fgame/cross/player/player"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func playerExitSceneBeforeLogout(target event.EventTarget, data event.EventData) (err error) {
	p := target.(*player.Player)
	//退出场景
	ctx := scene.WithPlayer(context.Background(), p)
	p.Post(message.NewScheduleMessage(onPlayerSceneExit, ctx, nil, nil))
	return nil
}

func onPlayerSceneExit(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(*player.Player)

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家角色正在登出,退出场景中")
	//退出场景
	scenelogic.PlayerExitScene(p, false)
	loginlogic.Logout(p)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypeCrossPlayerExitSceneBeforeLogout, event.EventListenerFunc(playerExitSceneBeforeLogout))
}
