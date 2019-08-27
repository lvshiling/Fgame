package rank

import (
	rankentity "fgame/fgame/game/rank/entity"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterRankSystemDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeShenfa, welfare.RankSystemDataHandlerFunc(getShenFaRankData))
}

func getShenFaRankData(groupId int32, page int32) (rankList []*rankentity.PlayerOrderData, rankTime int64) {
	rankList, rankTime = rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeShenFa, ranktypes.RankClassTypeLocalActivity, groupId, page)
	if page == 0 {
		nextPageList, _ := rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeShenFa, ranktypes.RankClassTypeLocalActivity, groupId, page+1)
		if len(nextPageList) > 0 {
			rankList = append(rankList, nextPageList[0])
		}
	}

	return
}
