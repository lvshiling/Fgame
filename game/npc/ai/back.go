package ai

import (
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterDefaultAction(scene.NPCStateBack, scene.NPCActionHandler(backAction))
}

//返回
func backAction(n scene.NPC) {
	//判断是否正在移动
	if n.IsMove() {
		return
	}

	if coreutils.DistanceSquare(n.GetPosition(), n.GetBornPosition()) > 0 {
		pos := n.GetBornPosition()
		flag := n.SetDestPosition(pos)
		if !flag {
			log.WithFields(
				log.Fields{
					"npc":   n.GetName(),
					"pos":   pos.String(),
					"mapId": n.GetScene().MapId(),
				}).Warn("npc:返回找不到路")
			return
		}
		return
	}
	n.Idle()
	// //回到出生地
	// flag := scenelogic.MoveTo(n, n.GetBornPosition(), 300)
	// if !flag {
	// 	return
	// }
	// //TODO 优化
	// if coreutils.DistanceSquare(n.GetPosition(), n.GetBornPosition()) > 0 {
	// 	return
	// }
	// n.Idle()
}
