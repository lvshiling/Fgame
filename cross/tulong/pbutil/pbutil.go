package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCTuLongBossStatus(status int32, biaoShi int32) *uipb.SCTuLongBossStatus {
	tuLongBossStatus := &uipb.SCTuLongBossStatus{}
	tuLongBossStatus.Status = &status
	tuLongBossStatus.BiaoShi = &biaoShi
	return tuLongBossStatus
}

func BuildSCTuLongAllianceBiaoShi(biaoShi int32) *uipb.SCTuLongAllianceBiaoShi {
	tuLongAllianceBiaoShi := &uipb.SCTuLongAllianceBiaoShi{}
	tuLongAllianceBiaoShi.BiaoShi = &biaoShi
	return tuLongAllianceBiaoShi
}

func BuildSCTuLongResult(killNum int32, itemMap map[int32]int32) *uipb.SCTuLongResult {
	tuLongResult := &uipb.SCTuLongResult{}
	tuLongResult.Num = &killNum

	for itemId, num := range itemMap {
		tuLongResult.ItemList = append(tuLongResult.ItemList, buildItem(itemId, num))
	}
	return tuLongResult
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}

func BuildSCTuLongCollect(npcId int64) *uipb.SCTuLongCollect {
	tuLongCollect := &uipb.SCTuLongCollect{}
	tuLongCollect.NpcId = &npcId
	return tuLongCollect
}

func BuildSCTuLongCollectStop(npcId int64) *uipb.SCTuLongCollectStop {
	tuLongCollectStop := &uipb.SCTuLongCollectStop{}
	tuLongCollectStop.NpcId = &npcId
	return tuLongCollectStop
}

func BuildSCTuLongCollectFinish(npcId int64) *uipb.SCTuLongCollectFinish {
	tuLongCollectFinish := &uipb.SCTuLongCollectFinish{}
	tuLongCollectFinish.NpcId = &npcId
	return tuLongCollectFinish
}
