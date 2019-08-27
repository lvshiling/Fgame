package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//仙盟仓库变更
func depotChanged(target event.EventTarget, data event.EventData) (err error) {
	al, ok := target.(*alliance.Alliance)
	if !ok {
		return
	}
	changedItemList, ok := data.([]*alliance.AllianceDepotItemObject)
	if !ok {
		return
	}

	//变更推送
	scMsg := pbutil.BuildSCAllianceDepotChangedNotice(changedItemList)
	for _, mem := range al.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			continue
		}
		p.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDepotChanged, event.EventListenerFunc(depotChanged))
}
