package ai

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	scene.RegisterDefaultAction(scene.NPCStateDead, scene.NPCActionHandler(deadAction))
}

//死亡动作
func deadAction(n scene.NPC) {
	if !n.GetBiologyTemplate().CanReborn() {
		return
	}
	now := global.GetGame().GetTimeService().Now()

	if n.GetBiologyTemplate().GetRebornType() == scenetypes.BiologyRebornTypeCall {
		n.Reborn(n.GetBornPosition())
		return
	}
	switch n.GetBiologyTemplate().GetRebornType() {
	case scenetypes.BiologyRebornTypeCall:
		n.Reborn(n.GetBornPosition())
		return
	case scenetypes.BiologyRebornTypeSecond:
		elapse := now - n.GetDeadTime()
		if elapse < n.GetBiologyTemplate().GetRebornTime(now) {
			return
		}
		n.Reborn(n.GetBornPosition())
		return
	case scenetypes.BiologyRebornTypeTime:
		day, _ := timeutils.DiffDay(now, n.GetDeadTime())
		//跨2天了
		if day >= 2 {
			n.Reborn(n.GetBornPosition())
			return
		}
		rebornTime := n.GetBiologyTemplate().GetRebornTime(now)
		if rebornTime < n.GetDeadTime() {
			return
		}
		//超过重生时间
		if now >= n.GetBiologyTemplate().GetRebornTime(now) {
			n.Reborn(n.GetBornPosition())
			return
		}

		if day == 1 {
			lastRebornTime := n.GetBiologyTemplate().GetRebornTime(n.GetDeadTime())
			// 重生后死亡的
			if lastRebornTime < n.GetDeadTime() {
				return
			}
			n.Reborn(n.GetBornPosition())
			return
		}

	}

	return

}
