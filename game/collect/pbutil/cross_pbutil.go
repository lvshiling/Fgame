package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSICollectFinish(itemMap map[int32]int32) *crosspb.SICollectFinish {
	siCollectFinish := &crosspb.SICollectFinish{}
	for itemId, num := range itemMap {
		siCollectFinish.ItemList = append(siCollectFinish.ItemList, buildItem(itemId, num))
	}
	return siCollectFinish
}

func buildItem(itemId int32, num int32) *crosspb.ItemInfo {
	itemInfo := &crosspb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}

func BuildSICollectMiZangFinish(npcId int64) *crosspb.SICollectMiZangFinish {
	siCollectMiZangFinish := &crosspb.SICollectMiZangFinish{}
	siCollectMiZangFinish.NpcId = &npcId
	return siCollectMiZangFinish
}
