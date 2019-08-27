package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//同步玩家
func playerPkCheck(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	s := p.GetScene()
	if s == nil {
		return
	}
	if !p.IsGuaJi() {
		return
	}
	bo := p.GetAttackTarget()
	if bo == nil {
		//TODO: 切换回来
		return
	}
	if p.IsEnemy(bo) {
		return
	}
	//城战不切换pk
	if s.MapTemplate().GetMapType() == scenetypes.SceneTypeChengZhan {
		return
	}
	switch tbo := bo.(type) {
	case scene.Player:
		//TODO:修复
		//切换全体
		if s.MapTemplate().LimitPkMode&pktypes.PkStateAll.Mask() != 0 {
			tbo.SwitchPkState(pktypes.PkStateAll, pktypes.PkCommonCampDefault)
			return
		}

		if s.MapTemplate().LimitPkMode&pktypes.PkStateBangPai.Mask() != 0 {
			if p.GetAllianceId() != 0 {
				tbo.SwitchPkState(pktypes.PkStateBangPai, pktypes.PkCommonCampDefault)
				return
			}
		}
		break
	}
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPkCheck, event.EventListenerFunc(playerPkCheck))
}
