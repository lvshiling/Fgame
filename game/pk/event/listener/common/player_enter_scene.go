package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
)

//玩家玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	s := p.GetScene()

	//没有特殊处理的
	//TODO 验证pk模式
	//设置pk模式
	pkState := s.MapTemplate().GetPkState()
	p.SwitchPkState(pkState, pktypes.PkCommonCampDefault)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
