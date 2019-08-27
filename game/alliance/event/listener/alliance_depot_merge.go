package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//仙盟仓库整理
func depotMerge(target event.EventTarget, data event.EventData) (err error) {
	al, ok := target.(*alliance.Alliance)
	if !ok {
		return
	}
	itemList, ok := data.([]*alliance.AllianceDepotItemObject)
	if !ok {
		return
	}

	//变更推送
	scMsg := pbutil.BuildSCAllianceDepotMergeNotice(itemList)
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
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDepotMerge, event.EventListenerFunc(depotMerge))
}
