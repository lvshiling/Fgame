package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//移动
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	//添加无敌buff
	s := p.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().IfChangeSceneProtect() {
		buffId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeProtectBuff)
		scenelogic.AddBuff(p, buffId, p.GetId(), common.MAX_RATE)
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
