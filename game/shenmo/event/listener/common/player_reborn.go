package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	shenmotemplate "fgame/fgame/game/shenmo/template"
)

//玩家复活
func playerReborn(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}

	s := p.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossShenMo {
		return
	}

	//添加无敌buff
	shenMoConstantTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoConstantTemplate()
	if shenMoConstantTemplate == nil {
		return
	}
	scenelogic.AddBuff(p, shenMoConstantTemplate.RebornBuff, p.GetId(), common.MAX_RATE)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(playerReborn))
}
