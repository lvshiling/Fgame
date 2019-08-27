package listener

import (
	"fgame/fgame/core/event"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//防护罩生成
func allianceSceneInitProtect(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(alliancescene.AllianceSceneData)
	if !ok {
		return
	}
	n, ok := data.(scene.NPC)
	if !ok {
		return
	}

	// 加无敌buff
	warTemp := alliancetemplate.GetAllianceTemplateService().GetWarTemplate()
	scenelogic.AddBuff(n, warTemp.ProtectBuffId, n.GetId(), int64(common.MAX_RATE))

	// 驱逐区域玩家
	for _, p := range sd.GetScene().GetAllPlayers() {
		if !warTemp.IsOnProtectArea(p.GetPos()) {
			continue
		}

		scenelogic.FixPosition(p, warTemp.GetProtectFixPos())
	}

	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneInitProtect, event.EventListenerFunc(allianceSceneInitProtect))
}
