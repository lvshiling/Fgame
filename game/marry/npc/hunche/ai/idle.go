package marry

import (
	hunchenpc "fgame/fgame/game/marry/npc/hunche"
	marrytemplate "fgame/fgame/game/marry/template"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterAction(scenetypes.BiologyScriptTypeWeddingCar, scene.NPCStateInit, scene.NPCActionHandler(idleAction))
}

//行走
func idleAction(n scene.NPC) {
	// 强制转换
	hunChe := n.(*hunchenpc.HunCheNPC)
	moveId := hunChe.GetHunCheObject().GetMoveId()
	moveTemplate := marrytemplate.GetMarryTemplateService().GetMarryMoveTeamplate(moveId)
	nextTemp := moveTemplate.GetNextTemp()
	if nextTemp == nil {
		return
	}

	pos := nextTemp.GetPos()

	mapId := nextTemp.MapId
	s := scene.GetSceneService().GetWorldSceneByMapId(mapId)
	if s == nil {
		return
	}
	if s == n.GetScene() {
		flag := hunChe.SetDestPosition(pos)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": hunChe.GetOwnerId(),
				}).Warn("marry:婚车找不到路")
			return
		}
	} else {
		scenelogic.NPCEnterScene(n, s, pos)
	}
}
