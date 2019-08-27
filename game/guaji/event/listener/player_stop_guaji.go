package listener

import (
	"fgame/fgame/core/event"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	"fgame/fgame/game/player"

	log "github.com/Sirupsen/logrus"
)

//玩家退出场景
func guaJiStop(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	pl.StopGuaJi()
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("guaji:停止挂机")

	//跨服中退出跨服
	if pl.IsCross() {
		crosslogic.PlayerExitCross(pl)
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().IsWorld() {
		return
	}
	//返回世界地图
	pl.BackLastScene()
	return
}

func init() {
	gameevent.AddEventListener(guajieventtypes.GuaJiEventTypeGuaJiStop, event.EventListenerFunc(guaJiStop))
}
