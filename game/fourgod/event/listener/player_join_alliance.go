package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	gameevent "fgame/fgame/game/event"
	fourgodscene "fgame/fgame/game/fourgod/scene"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//玩家加入仙盟
func playerJoinAlliance(target event.EventTarget, data event.EventData) (err error) {
	memObj, ok := data.(*alliance.AllianceMemberObject)
	if !ok {
		return
	}
	playerId := memObj.GetMemberId()

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}
	_, ok = sd.(fourgodscene.FourGodWarSceneData)
	if !ok {
		return
	}
	p, ok := pl.(scene.Player)
	if !ok {
		return
	}
	p.SwitchPkState(pktypes.PkStateBangPai, pktypes.PkCommonCampDefault)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMemberJoin, event.EventListenerFunc(playerJoinAlliance))
}
