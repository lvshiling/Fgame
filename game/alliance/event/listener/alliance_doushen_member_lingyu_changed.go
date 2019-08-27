package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//斗神领域变化
func allianceDouShenMemberLingYuChanged(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	memObj := data.(*alliance.AllianceMemberObject)
	lingyuId := memObj.GetLingyuId()
	memId := memObj.GetMemberId()

	pl := player.GetOnlinePlayerManager().GetPlayerById(memId)		
	if pl == nil {
		return
	}

	//广播内容
	scAllianceDouShenLingyuChangedBroadcast := pbutil.BuildSCAllianceDouShenLingyuChangedBroadcast(lingyuId, pl)

	for _, mem := range al.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			continue
		}
		if p.GetId() != memId {
			p.SendMsg(scAllianceDouShenLingyuChangedBroadcast)
		}
		alliancelogic.DoushenChanged(p)
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDouShenMemberLingYuChanged, event.EventListenerFunc(allianceDouShenMemberLingYuChanged))
}
