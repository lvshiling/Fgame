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

//buff移除
func buffRemove(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)
	buffObject := data.(buffcommon.BuffObject)

	buffId := buffObject.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if bo.GetScene() != nil {
		scObjectBuffRemove := pbutil.BuildSCObjectBuffRemove(bo, buffObject)
		scenelogic.BroadcastNeighborIncludeSelf(bo, scObjectBuffRemove)
	}

	//同步属性
	if buffTemplate.GetTouchType() == scenetypes.BuffTouchTypeImmediate {
		bufflogic.UpdateBattleProperty(bo)
	}
	//判断是否是变形
	if buffTemplate.ModelId != 0 {
		switch tbo := bo.(type) {
		case scene.Player:
			if buffTemplate.ModelId == tbo.GetModel() {
				tbo.SetModel(0)
			}
			break
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffRemove, event.EventListenerFunc(buffRemove))
}
