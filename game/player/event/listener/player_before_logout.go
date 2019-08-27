package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	loginlogic "fgame/fgame/game/login/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"

	playereventtypes "fgame/fgame/game/player/event/types"

	log "github.com/Sirupsen/logrus"
)

func playerBeforeLogout(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家角色登出前")
	crossSession := p.GetCrossSession()
	if crossSession != nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"state":    p.CurrentState(),
			}).Info("player:玩家角色登出前,跨服对话关闭")
		crossSession.Close(true)
	}
	//退出场景
	ctx := scene.WithPlayer(context.Background(), p)
	global.GetGame().GetGlobalRunner().Post(message.NewScheduleMessage(onPlayerBeforeLogout, ctx, nil, nil))
	return nil
}

func onPlayerBeforeLogout(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	loginlogic.Logout(p)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerBeforeLogout, event.EventListenerFunc(playerBeforeLogout))
}
