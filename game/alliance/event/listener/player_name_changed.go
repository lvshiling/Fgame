package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	mem := alliance.GetAllianceService().GetAllianceMember(p.GetId())
	allianceManager := p.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	if mem == nil {
		//退出仙盟
		allianceManager.SyncAlliance(0, 0, "", 0, 0)
		return
	}

	//同步成员信息
	alliance.GetAllianceService().SyncMemberInfo(p.GetId(), p.GetName(), p.GetSex(), p.GetLevel(), p.GetForce(), p.GetZhuanSheng(), p.GetLingyuInfo().AdvanceId, p.GetVip())

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
