package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

//怪物被击杀
func monsterKilled(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	n, ok := data.(scene.NPC)
	if !ok {
		return
	}

	biologyGroupIdList := n.GetBiologyTemplate().GetGroupIdList()
	if len(biologyGroupIdList) == 0 {
		return
	}

	// 首杀活动
	huhuFirstDrop(pl, n)

	return nil
}

// 首杀活动
func huhuFirstDrop(pl player.Player, n scene.NPC) {
	s := n.GetScene()
	if s == nil {
		return
	}

	biologyGroupIdList := n.GetBiologyTemplate().GetGroupIdList()
	biologyId := int32(n.GetBiologyTemplate().Id)

	typ := welfaretypes.OpenActivityTypeHuHu
	subType := welfaretypes.OpenActivitySpecialSubTypeFirstDrop
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		// 是否活动怪
		for _, biologyGroupId := range biologyGroupIdList {
			if groupId != biologyGroupId {
				continue
			}

			if welfare.GetWelfareService().IsBossFirstKill(groupId, biologyId) {
				continue
			}
			dropId := welfaretemplate.GetWelfareTemplateService().GetGroupBiologyDropId(groupId, 0)
			scenelogic.CustomDrop(s, n, n.GetPosition(), pl.GetId(), []int32{dropId}, 1)
			welfare.GetWelfareService().AddBossKillRecord(groupId, biologyId)
		}
	}
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterKilled, event.EventListenerFunc(monsterKilled))
}
