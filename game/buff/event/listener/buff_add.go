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

//buff添加
func buffAdd(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.BattleObject)

	buffObject := data.(buffcommon.BuffObject)

	buffId := buffObject.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	if bo.GetScene() != nil {
		//zrc:特殊处理定制称号下马
		if buffTemplate.GetBuffType() == scenetypes.BuffTypeTitleDingZhi {
			switch tbo := bo.(type) {
			case scene.Player:
				if !tbo.IsMountHidden() {
					tbo.MountHidden(true)
				}
				break
			}
		}
		scObjectBuff := pbutil.BuildSCObjectBuff(bo, buffObject)
		scenelogic.BroadcastNeighborIncludeSelf(bo, scObjectBuff)
	}

	//立即触发
	if buffTemplate.GetTouchType() == scenetypes.BuffTouchTypeImmediate {
		bo.TouchBuff(buffId)
	}

	//控制类buff
	if buffTemplate.TypeLimit != 0 {
		//判断是否触发
		bufflogic.TouchBuffByAction(bo, scenetypes.BuffTouchTypeBuff)
	}

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffAdd, event.EventListenerFunc(buffAdd))
}
