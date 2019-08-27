package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	longgonglogic "fgame/fgame/game/longgong/logic"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//黑龙boss被击杀 刷采集物
func monsterKilled(target event.EventTarget, data event.EventData) (err error) {
	//场景过滤
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	n := data.(scene.NPC)
	bioType := n.GetBiologyTemplate().GetBiologyScriptType()
	if bioType != scenetypes.BiologyScriptTypeLongGongBoss {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeLongGong {
		return
	}

	sd := s.SceneDelegate()
	longgongSd, ok := sd.(longgonglogic.LongGongSceneData)
	if !ok {
		return
	}

	longgongSd.DealHeiLongBossDead()

	//龙宫排行奖励
	longgonglogic.LongGongRankRewards(longgongSd)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterKilled, event.EventListenerFunc(monsterKilled))
}
