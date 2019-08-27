package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/welfare"
)

//排行榜结束
func rankEnd(target event.EventTarget, data event.EventData) (err error) {
	groupId, ok := target.(int32)
	if !ok {
		return
	}
	rankType, ok := data.(ranktypes.RankType)
	if !ok {
		return
	}

	// timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	// now := global.GetGame().GetTimeService().Now()
	// openTime := welfare.GetWelfareService().GetServerStartTime() //global.GetGame().GetServerTime()
	// TODO :xzk 循环活动计算结束时间会计算错误，下一天的结束时间
	// endTime, err := timeTemp.GetEndTime(now, openTime)
	// if err != nil {
	// 	return
	// }
	_, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	err = rank.GetRankService().UpdateRankData(ranktypes.RankClassTypeLocalActivity, rankType, groupId)
	if err != nil {
		return
	}

	rankList := rank.GetRankService().GetRankingInfoList(ranktypes.RankClassTypeLocalActivity, rankType, groupId)
	// 发送排行奖励
	welfarelogic.AddRankRewards(rankList, groupId, endTime)

	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeRankEnd, event.EventListenerFunc(rankEnd))
}
