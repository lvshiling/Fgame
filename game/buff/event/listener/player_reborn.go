package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//移动
func playerReborn(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	//添加无敌buff
	buffId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveProtectBuff)
	scenelogic.AddBuff(p, buffId, p.GetId(), common.MAX_RATE)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(playerReborn))
}
