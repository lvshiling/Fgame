package rank

import (
	rankentity "fgame/fgame/game/rank/entity"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterRankPropertyDataHandler(welfaretypes.OpenActivityTypeRank, welfaretypes.OpenActivityRankSubTypeGoldEquipForce, welfare.RankPropertyDataHandlerFunc(getGoldEquipForceRankData))
}

func getGoldEquipForceRankData(groupId int32, page int32) (rankList []*rankentity.PlayerPropertyData, rankTime int64) {
	rankList, rankTime = rank.GetRankService().GetPropertyListByPage(ranktypes.RankTypeGoldEquipForce, ranktypes.RankClassTypeLocalActivity, groupId, page)
	if page == 0 {
		nextPageList, _ := rank.GetRankService().GetPropertyListByPage(ranktypes.RankTypeGoldEquipForce, ranktypes.RankClassTypeLocalActivity, groupId, page+1)
		if len(nextPageList) > 0 {
			rankList = append(rankList, nextPageList[0])
		}
	}

	return
}
