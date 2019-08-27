package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	playergem "fgame/fgame/game/gem/player"
)

func BuildSCGemMineGet(mine *playergem.PlayerMiningObject) *uipb.SCGemMineGet {
	gemMineGet := &uipb.SCGemMineGet{}
	stone := mine.Stone
	gemMineGet.Stone = &stone
	gemMineGet.MineInfo = buildMine(mine)
	return gemMineGet
}

func BuildSCGemMineActive(mine *playergem.PlayerMiningObject) *uipb.SCGemMineActive {
	mineMineActive := &uipb.SCGemMineActive{}
	mineMineActive.MineInfo = buildMine(mine)
	return mineMineActive
}

func BuildSCGemMineReceive(stone int64, lastTime int64) *uipb.SCGemMineReceive {
	gemMineReceive := &uipb.SCGemMineReceive{}
	gemMineReceive.Stone = &stone
	gemMineReceive.LastTime = &lastTime
	return gemMineReceive
}

func BuildSCGemGamble(dropItemList []*droptemplate.DropItemData, stone int64) *uipb.SCGemGamble {
	gemGamble := &uipb.SCGemGamble{}
	gemGamble.Stone = &stone
	for _, itemData := range dropItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()
		gemGamble.DropList = append(gemGamble.DropList, buildGamble(itemId, num, level))
	}

	return gemGamble
}

func buildMine(mine *playergem.PlayerMiningObject) *uipb.MineInfo {
	mineInfo := &uipb.MineInfo{}
	storage := mine.Storage
	lastTime := mine.LastTime
	level := mine.Level

	mineInfo.Storage = &storage
	mineInfo.Level = &level
	mineInfo.LastTime = &lastTime
	return mineInfo
}

func buildGamble(itemId int32, num int32, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level
	return dropInfo
}
