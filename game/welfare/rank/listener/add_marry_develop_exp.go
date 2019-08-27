package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家增长表白经验
func playerRankAddMarryDevelopExp(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	addNum, ok := data.(int32)
	if !ok {
		return
	}

	//增长表白经验排行活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	rankTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeMarryDevelop)
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
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryDevelopExpAdd, event.EventListenerFunc(playerRankAddMarryDevelopExp))
}
