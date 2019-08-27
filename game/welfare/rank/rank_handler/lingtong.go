package rank

import (
	rankentity "fgame/fgame/game/rank/entity"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingQi, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingBing, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingBao, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingTi, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingYi, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingYu, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeLingShen, welfare.RankSystemDataHandlerFunc(getLingTongRankData))
}

//灵童
func getLingTongRankData(groupId int32, page int32) (rankList []*rankentity.PlayerOrderData, rankTime int64) {
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	if timeTemp == nil {
		return
	}

	rankSubType, ok := timeTemp.GetOpenSubType().(welfaretypes.OpenActivityRankSubType)
	if !ok {
		return
	}

	rankList, rankTime = rank.GetRankService().GetOrderListByPage(rankSubType.RankType(), ranktypes.RankClassTypeLocalActivity, groupId, page)
	if page == 0 {
		// 因前端七日冲刺显示需要第11名的信息，特殊处理
		nextPageList, _ := rank.GetRankService().GetOrderListByPage(rankSubType.RankType(), ranktypes.RankClassTypeLocalActivity, groupId, page+1)
		if len(nextPageList) > 0 {
			rankList = append(rankList, nextPageList[0])
		}
	}

	return
}
