package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	rankeventtypes "fgame/fgame/game/rank/event/types"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/game/welfare/pbutil"
	welfaretemplate "fgame/fgame/game/welfare/template"
)

const (
	noticeRankingLimit = 3 //下降通知名次限制
)

//排行榜改变
func rankChanged(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*rankeventtypes.RankChangedEventData)
	if !ok {
		return
	}
	rankClassType := eventData.GetRankClassType()
	if rankClassType != ranktypes.RankClassTypeLocalActivity {
		return
	}

	//rankType := eventData.GetRankType()
	groupId := eventData.GetGroupId()
	oldList := eventData.GetOldRankList()
	newList := eventData.GetNewRankList()

	var oldTopThreeList []*ranktypes.RankingInfo
	if len(oldList) <= noticeRankingLimit {
		oldTopThreeList = oldList
	} else {
		oldTopThreeList = oldList[:noticeRankingLimit]
	}

	//排行榜前三下降通知
	for _, oldInfo := range oldTopThreeList {
		oldPl := player.GetOnlinePlayerManager().GetPlayerById(oldInfo.GetPlayerId())
		if oldPl == nil {
			continue
		}
		oldRanking := oldInfo.GetRanking()
		newRanking := int32(0)
		for _, newInfo := range newList {
			if newInfo.GetPlayerId() != oldInfo.GetPlayerId() {
				continue
			}

			newRanking = newInfo.GetRanking()
		}

		// 排名不变或上升
		if oldRanking >= newRanking {
			continue
		}

		timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
		if timeTemp == nil {
			continue
		}
		scMsg := pbutil.BuildSCOpenActivityRankingDropDownNotice(timeTemp.Name, oldRanking, newRanking, groupId)
		oldPl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(rankeventtypes.RankEventTypeRankChanged, event.EventListenerFunc(rankChanged))
}
