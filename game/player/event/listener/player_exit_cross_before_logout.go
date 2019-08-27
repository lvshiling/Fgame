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

func playerExitCrossBeforeLogout(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色正在登出,退出跨服前")
	s := p.GetCrossSession()
	if s != nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Info("player:玩家角色正在登出,关闭跨服对话")
		s.Close(true)
	} else {
		ctx := scene.WithPlayer(context.Background(), p)
		global.GetGame().GetGlobalRunner().Post(message.NewScheduleMessage(onPlayerExitCrossBeforeLogout, ctx, nil, nil))
	}

	return nil
}

func onPlayerExitCrossBeforeLogout(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色正在登出,退出跨服")

	//保存数据 退出服务器
	loginlogic.Logout(p)
	return nil
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerExitCrossBeforeLogout, event.EventListenerFunc(playerExitCrossBeforeLogout))
}
