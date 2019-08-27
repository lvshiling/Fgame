package pbuitl

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerdragon "fgame/fgame/game/dragon/player"
)

func BuildSCDragonGet(dragonObj *playerdragon.PlayerDragonObject) *uipb.SCDragonGet {
	dragonGet := &uipb.SCDragonGet{}
	stageId := dragonObj.StageId

	dragonGet.StageId = &stageId
	for itemId, num := range dragonObj.ItemInfoMap {
		dragonGet.ItemList = append(dragonGet.ItemList, buildItem(itemId, num))
	}

	return dragonGet
}

func BuildSCDragonFeed(dragonObj *playerdragon.PlayerDragonObject) *uipb.SCDragonFeed {
	dragonFeed := &uipb.SCDragonFeed{}
	stageId := dragonObj.StageId
	status := false
	if dragonObj.Status == 1 {
		status = true
	}
	dragonFeed.StageId = &stageId
	dragonFeed.Status = &status
	for itemId, num := range dragonObj.ItemInfoMap {
		dragonFeed.ItemList = append(dragonFeed.ItemList, buildItem(itemId, num))
	}
	return dragonFeed
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}
