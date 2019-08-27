package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"

	playereventtypes "fgame/fgame/game/player/event/types"

	log "github.com/Sirupsen/logrus"
)

func playerLogoutBeforeLoaded(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家角色登出在加载完前")
	//退出服务器
	player.GetOnlinePlayerManager().PlayerLeaveServer(p)
	//移除用户
	player.GetOnlinePlayerManager().RemovePlayer(p)
	return nil
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogoutBeforeLoaded, event.EventListenerFunc(playerLogoutBeforeLoaded))
}
