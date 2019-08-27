package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenalogic "fgame/fgame/cross/arena/logic"
	arenascene "fgame/fgame/cross/arena/scene"
	gameevent "fgame/fgame/game/event"

	log "github.com/Sirupsen/logrus"
)

//竞技场进入下一场匹配
func arenaNextMatch(target event.EventTarget, data event.EventData) (err error) {
	t := target.(*arenascene.ArenaTeam)

	log.WithFields(
		log.Fields{
			"team": t.String(),
		}).Debug("arena:匹配成功下一场")

	arenalogic.OnArenaNextMatch(t)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaNextMatch, event.EventListenerFunc(arenaNextMatch))
}
