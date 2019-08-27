package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	playerhunt "fgame/fgame/game/hunt/player"
	hunttypes "fgame/fgame/game/hunt/types"
)

func BuildSCHuntXunBao(rewItemList []*droptemplate.DropItemData, huntType int32, lastHuntTime int64, freeCount int32) *uipb.SCHuntXunBao {
	scMsg := &uipb.SCHuntXunBao{}
	for _, itemData := range rewItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()

		scMsg.DropInfo = append(scMsg.DropInfo, buildDropInfo(itemId, num, level))
	}
	scMsg.HuntType = &huntType
	scMsg.LastFreeHuntTime = &lastHuntTime
	scMsg.HuntCount = &freeCount
	return scMsg
}

func BuildSCHuntInfoNotice(huntInfoMap map[hunttypes.HuntType]*playerhunt.PlayerHuntObject) *uipb.SCHuntInfoNotice {
	scMsg := &uipb.SCHuntInfoNotice{}

	for _, huntObj := range huntInfoMap {
		huntCount := huntObj.GetFreeHuntCount()
		lastTime := huntObj.GetLastHuntTime()
		huntType := int32(huntObj.GetHuntType())

		info := &uipb.HuntFreeInfo{}
		info.HuntType = &huntType
		info.HuntCount = &huntCount
		info.LastHuntTime = &lastTime

		scMsg.InfoList = append(scMsg.InfoList, info)
	}

	return scMsg
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}
