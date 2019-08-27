package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家消耗元宝
func playerCostRank(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int64)
	if !ok {
		return
	}

	//消费排行活动
	welfareManager := p.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	rankTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeCost)
	for _, timeTemp := range rankTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		curGold := welfareManager.AddOpenActivityCostRecord(groupId, goldNum)
		ss := pbutil.BuildSCOpenActivityRankCostNotice(curGold, groupId)
		p.SendMsg(ss)
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCost, event.EventListenerFunc(playerCostRank))
}
