package listener

import (
	"fgame/fgame/core/event"
	buffcommon "fgame/fgame/game/buff/common"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	bufflogic "fgame/fgame/game/buff/logic"
	bufftemplate "fgame/fgame/game/buff/template"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//buff更新
func buffUpdate(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)
	buffObject := data.(buffcommon.BuffObject)

	buffId := buffObject.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if bo.GetScene() != nil {
		scObjectBuff := pbutil.BuildSCObjectBuff(bo, buffObject)
		scenelogic.BroadcastNeighborIncludeSelf(bo, scObjectBuff)

	}

	//同步属性
	if buffTemplate.GetTouchType() == scenetypes.BuffTouchTypeImmediate {
		bufflogic.UpdateBattleProperty(bo)
		//TODO 更新护盾
	}
	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffUpdate, event.EventListenerFunc(buffUpdate))
}
