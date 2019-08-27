package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tulongeventtypes "fgame/fgame/game/tulong/event/types"
	"fgame/fgame/game/tulong/pbutil"
)

//跨服屠龙任务开始
func tuLongActivityStart(target event.EventTarget, data event.EventData) (err error) {

	for _, al := range alliance.GetAllianceService().GetAllianceList() {
		memberList := al.GetMemberList()
		for _, member := range memberList {
			playerId := member.GetMemberId()
			pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
			if pl == nil {
				continue
			}
			scTuLongStart := pbutil.BuildSCTuLongStart()
			pl.SendMsg(scTuLongStart)
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(tulongeventtypes.EventTypeTuLongActivityStart, event.EventListenerFunc(tuLongActivityStart))
}
