package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家增长魅力值
func playerRankAddCharmNum(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	addNum, ok := data.(int32)
	if !ok {
		return
	}

	//增长魅力值排行活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	rankTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeCharm)
	for _, timeTemp := range rankTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		welfareManager.AddActivityAddNumVal(groupId, addNum)
	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerCharmAdd, event.EventListenerFunc(playerRankAddCharmNum))
}
