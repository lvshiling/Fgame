package biaoche

import (
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	transportationtemplate "fgame/fgame/game/transportation/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterAction(scenetypes.BiologyScriptTypeBiaoChe, scene.NPCStateInit, scene.NPCActionHandler(idleAction))
}

//行走
func idleAction(n scene.NPC) {

	// 强制转换
	biaoChe := n.(*biaochenpc.BiaocheNPC)
	moveId := biaoChe.GetTransportationObject().GetTransportMoveId()
	moveTemplate := transportationtemplate.GetTransportationTemplateService().GetTransportationMoveTemplate(moveId)
	nextTemp := moveTemplate.GetNextTemp()
	if nextTemp == nil {
		return
	}

	pos := nextTemp.GetPosition()
	mapId := nextTemp.MapId
	s := scene.GetSceneService().GetWorldSceneByMapId(mapId)
	if s == nil {
		return
	}
	if s == n.GetScene() {
		flag := n.SetDestPosition(pos)
		if !flag {
			log.WithFields(
				log.Fields{
					"npc":   biaoChe.GetName(),
					"pos":   pos.String(),
					"mapId": biaoChe.GetScene().MapId(),
				}).Warn("transportation:找不到路")
		}
	} else {
		scenelogic.NPCEnterScene(n, s, pos)
	}
}
