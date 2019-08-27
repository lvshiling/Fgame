package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	arenaeventtypes "fgame/fgame/cross/arena/event/types"
	arenalogic "fgame/fgame/cross/arena/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

//竞技场匹配成功
func arenaMatched(target event.EventTarget, data event.EventData) (err error) {
	s := target.(scene.Scene)
	eventData := data.(*arena.MatchEventData)
	team1 := eventData.GetTeam1()
	team2 := eventData.GetTeam2()
	log.WithFields(
		log.Fields{
			"team1": team1.String(),
			"team2": team2.String(),
		}).Debug("arena:匹配成功成功")

	arenalogic.OnArenaMatched(s, team1, team2)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaMatched, event.EventListenerFunc(arenaMatched))
}
