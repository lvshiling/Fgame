package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	lianyuscene "fgame/fgame/game/lianyu/scene"
	"fgame/fgame/game/scene/scene"
)

func BuildSCLianYuLineUp(beforeNum int32) *uipb.SCLianYuLineUp {
	scLianYuLineUp := &uipb.SCLianYuLineUp{}
	scLianYuLineUp.BeforeNum = &beforeNum
	return scLianYuLineUp
}

func BuildSCLianYuCancleLineUp() *uipb.SCLianYuCancleLineUp {
	scLianYuCancleLineUp := &uipb.SCLianYuCancleLineUp{}
	return scLianYuCancleLineUp
}

func BuildSCLianYuLineUpSuccess() *uipb.SCLianYuLineUpSuccess {
	scLianYuLineUpSuccess := &uipb.SCLianYuLineUpSuccess{}
	return scLianYuLineUpSuccess
}

func BuildSCLianYuFinishToLineUp() *uipb.SCLianYuFinishToLineUp {
	scLianYuFinishToLineUp := &uipb.SCLianYuFinishToLineUp{}
	return scLianYuFinishToLineUp
}

// -----------------------------------------------

func BuildSCLianYuLineUpChanged(beforeNum int32) *uipb.SCLianYuLineUp {
	scLianYuLineUp := &uipb.SCLianYuLineUp{}
	scLianYuLineUp.BeforeNum = &beforeNum
	return scLianYuLineUp
}

func BuildSCLianYuGet(bossNpc scene.NPC, bossStatus int32, rankList []*lianyuscene.LianYuRank, shaQiNum int32) *uipb.SCLianYuGet {
	scLianYuGet := &uipb.SCLianYuGet{}
	scLianYuGet.BossStatus = &bossStatus
	for index, shaQiRank := range rankList {
		scLianYuGet.RankList = append(scLianYuGet.RankList, buildShaQiRank(int32(index+1), shaQiRank))
	}
	if bossNpc != nil {
		pos := bossNpc.GetPosition()
		scLianYuGet.Pos = commonpbutil.BuildPos(pos)
	}
	scLianYuGet.ShaQiNum = &shaQiNum
	return scLianYuGet
}

func BuildSCLianYuBossDead(bossStatus int32) *uipb.SCLianYuBossStatus {
	scLianYuBossStatus := &uipb.SCLianYuBossStatus{}
	scLianYuBossStatus.BossStatus = &bossStatus
	return scLianYuBossStatus
}

func BuildSCLianYuBossRefresh(bossStatus int32, pos types.Position) *uipb.SCLianYuBossStatus {
	scLianYuBossStatus := &uipb.SCLianYuBossStatus{}
	scLianYuBossStatus.BossStatus = &bossStatus
	scLianYuBossStatus.Pos = commonpbutil.BuildPos(pos)
	return scLianYuBossStatus
}

func BuildSCLianYuResult(itemMap map[int32]int32) *uipb.SCLianYuResult {
	scLianYuResult := &uipb.SCLianYuResult{}
	for itemId, num := range itemMap {
		scLianYuResult.ItemList = append(scLianYuResult.ItemList, buildItem(itemId, num))
	}
	return scLianYuResult
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}

func BuildSCLianYuRankChanged(rankList []*lianyuscene.LianYuRank) *uipb.SCLianYuRankChanged {
	scLianYuRankChanged := &uipb.SCLianYuRankChanged{}
	for index, shaQiRank := range rankList {
		scLianYuRankChanged.RankList = append(scLianYuRankChanged.RankList, buildShaQiRank(int32(index+1), shaQiRank))
	}
	return scLianYuRankChanged
}

func BuildSCLianYuShaQiChanged(shaQiNum int32) *uipb.SCLianYuShaQiChanged {
	scMsg := &uipb.SCLianYuShaQiChanged{}
	scMsg.ShaQiNum = &shaQiNum
	return scMsg
}

func buildShaQiRank(pos int32, shaQiRank *lianyuscene.LianYuRank) *uipb.LianYuRank {
	lianYuRank := &uipb.LianYuRank{}
	serverId := shaQiRank.GetServerId()
	name := shaQiRank.GetName()
	shaQi := shaQiRank.GetShaQi()
	lianYuRank.Pos = &pos
	lianYuRank.ServerId = &serverId
	lianYuRank.Name = &name
	lianYuRank.Num = &shaQi
	return lianYuRank
}
