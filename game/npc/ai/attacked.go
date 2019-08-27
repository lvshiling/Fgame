package ai

import (
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterDefaultAction(scene.NPCStateAttacked, scene.NPCActionHandler(attackedAction))
}

//受击动作
func attackedAction(n scene.NPC) {

	now := global.GetGame().GetTimeService().Now()
	elapse := now - n.GetSkilledTime()
	if elapse < n.GetSkilledStopTime() {
		attackedMove(n)
		return
	}

	n.Trace()
}

//被动移动
func attackedMove(n scene.NPC) {
	//获取时间
	elapse := 100
	//获取移动距离
	speed := n.GetAttackedMoveSpeed()
	//移动的距离
	distance := (float64(elapse) / float64(common.SECOND)) * speed
	totalDistance := coreutils.Distance(n.GetPosition(), n.GetDestPosition())
	if totalDistance <= common.MIN_DISTANCE_ERROR {
		return
	}
	if distance > totalDistance {
		scenelogic.Move(n, n.GetDestPosition(), n.GetAngle(), speed, types.MoveTypeHit, true, false)
		return
	}
	t := distance / totalDistance
	targetPos := coreutils.Lerp(n.GetPosition(), n.GetDestPosition(), t)
	//获取位置
	targetPos.Y = n.GetScene().MapTemplate().GetMap().GetHeight(targetPos.X, targetPos.Z)
	scenelogic.Move(n, targetPos, n.GetAngle(), speed, types.MoveTypeHit, true, false)
}
