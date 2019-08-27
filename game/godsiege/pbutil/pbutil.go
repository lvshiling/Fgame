package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildSCGodSiegeLineUp(godType int32, beforeNum int32) *uipb.SCGodSiegeLineUp {
	scGodSiegeLineUp := &uipb.SCGodSiegeLineUp{}
	scGodSiegeLineUp.GodType = &godType
	scGodSiegeLineUp.BeforeNum = &beforeNum
	return scGodSiegeLineUp
}

func BuildSCGodSiegeCancleLineUp(godType int32) *uipb.SCGodSiegeCancleLineUp {
	scGodSiegeCancleLineUp := &uipb.SCGodSiegeCancleLineUp{}
	scGodSiegeCancleLineUp.GodType = &godType
	return scGodSiegeCancleLineUp
}

func BuildSCGodSiegeLineUpSuccess(godType int32) *uipb.SCGodSiegeLineUpSuccess {
	scGodSiegeLineUpSuccess := &uipb.SCGodSiegeLineUpSuccess{}
	scGodSiegeLineUpSuccess.GodType = &godType
	return scGodSiegeLineUpSuccess
}

func BuildSCGodSiegeFinishToLineUp(godType int32) *uipb.SCGodSiegeFinishToLineUp {
	scGodSiegeFinishToLineUp := &uipb.SCGodSiegeFinishToLineUp{}
	scGodSiegeFinishToLineUp.GodType = &godType
	return scGodSiegeFinishToLineUp
}

//
func BuildSCGodSiegeGet(pl scene.Player, godType int32, bossNpc scene.NPC, bossStatus int32, itemMap map[int32]int32, collectList map[int64]scene.NPC) *uipb.SCGodSiegeGet {
	scGodSiegeGet := &uipb.SCGodSiegeGet{}
	scGodSiegeGet.BossStatus = &bossStatus
	scGodSiegeGet.GodType = &godType
	scGodSiegeGet.CollectNpcInfo = scenepbutil.BuildGeneralCollectInfoList(collectList)
	if bossNpc != nil {
		pos := bossNpc.GetPosition()
		scGodSiegeGet.Pos = commonpbutil.BuildPos(pos)
	}
	if godsiegetypes.GodSiegeType(godType) == godsiegetypes.GodSiegeTypeDenseWat {
		num := pl.GetDenseWatNum()
		scGodSiegeGet.CollectInfo = buildCollectInfo(num, itemMap)
	}
	return scGodSiegeGet
}

func BuildSCGodSiegeCollectNpcChanged(npc scene.NPC) *uipb.SCGodSiegeCollectNpcChanged {
	scMsg := &uipb.SCGodSiegeCollectNpcChanged{}
	scMsg.CollectNpcInfo = scenepbutil.BuildGeneralCollectInfo(npc)
	return scMsg
}

func BuildSCGodSiegeBossDead(godType int32, bossStatus int32) *uipb.SCGodSiegeBossStatus {
	scGodSiegeBossStatus := &uipb.SCGodSiegeBossStatus{}
	scGodSiegeBossStatus.BossStatus = &bossStatus
	scGodSiegeBossStatus.GodType = &godType
	return scGodSiegeBossStatus
}

func BuildSCGodSiegeBossRefresh(godType int32, bossStatus int32, pos types.Position) *uipb.SCGodSiegeBossStatus {
	scGodSiegeBossStatus := &uipb.SCGodSiegeBossStatus{}
	scGodSiegeBossStatus.GodType = &godType
	scGodSiegeBossStatus.BossStatus = &bossStatus
	scGodSiegeBossStatus.Pos = commonpbutil.BuildPos(pos)
	return scGodSiegeBossStatus
}

func BuildSCGodSiegeResult(godType int32, itemMap map[int32]int32) *uipb.SCGodSiegeResult {
	scGodSiegeResult := &uipb.SCGodSiegeResult{}
	scGodSiegeResult.GodType = &godType
	for itemId, num := range itemMap {
		scGodSiegeResult.ItemList = append(scGodSiegeResult.ItemList, buildItem(itemId, num))
	}
	return scGodSiegeResult
}

func BuildSCGodSiegeCollectChanged(pl scene.Player, godType int32, itemMap map[int32]int32) *uipb.SCGodSiegeCollectChanged {
	scGodSiegeCollectChanged := &uipb.SCGodSiegeCollectChanged{}
	scGodSiegeCollectChanged.GodType = &godType
	num := pl.GetDenseWatNum()
	scGodSiegeCollectChanged.CollectInfo = buildCollectInfo(num, itemMap)
	return scGodSiegeCollectChanged
}

func buildCollectInfo(num int32, itemMap map[int32]int32) *uipb.CollectInfo {
	collectInfo := &uipb.CollectInfo{}
	collectInfo.Num = &num
	for itemId, num := range itemMap {
		collectInfo.ItemList = append(collectInfo.ItemList, buildItem(itemId, num))
	}
	return collectInfo
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}
