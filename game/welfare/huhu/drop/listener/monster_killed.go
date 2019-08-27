package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	huhudroptypes "fgame/fgame/game/welfare/huhu/drop/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
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

	// 击杀活动
	huhuDrop(pl, n)

	return nil
}

// 击杀活动
func huhuDrop(pl player.Player, n scene.NPC) {
	s := n.GetScene()
	if s == nil {
		return
	}
	biologyGroupIdList := n.GetBiologyTemplate().GetGroupIdList()
	typ := welfaretypes.OpenActivityTypeHuHu
	subType := welfaretypes.OpenActivitySpecialSubTypeDrop
	//掉落活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	dropTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range dropTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		// 是否活动怪
		for _, biologyGroupId := range biologyGroupIdList {
			if groupId != biologyGroupId {
				continue
			}
			welfareManager.RefreshActivityData(typ, subType)

			obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
			info := obj.GetActivityData().(*huhudroptypes.HuHuInfo)
			dropId := welfaretemplate.GetWelfareTemplateService().GetGroupBiologyDropId(groupId, info.CurDayDropNum)
			scenelogic.CustomDrop(s, n, n.GetPosition(), pl.GetId(), []int32{dropId}, 1)
		}
	}
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterKilled, event.EventListenerFunc(monsterKilled))
}
