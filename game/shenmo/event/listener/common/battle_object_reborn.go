package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/shenmo/pbutil"
	shenmotemplate "fgame/fgame/game/shenmo/template"
)

//神魔战场大旗重生
func battleObjectReborn(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}

	s := n.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossShenMo {
		return
	}

	biologyId := int32(n.GetBiologyTemplate().TemplateId())
	shenMoTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoConstantTemplate()
	if shenMoTemplate.DaQiBiologyId != biologyId {
		return
	}

	bcMsg := pbutil.BuildSCShenMoBioBroadcast(n)
	s.BroadcastMsg(bcMsg)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(battleObjectReborn))
}
