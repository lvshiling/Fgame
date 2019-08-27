package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	itemtypes "fgame/fgame/game/item/types"
	playerlucky "fgame/fgame/game/lucky/player"
)

func BuildSCLuckyInfoChanged(expire int64, itemId int32) *uipb.SCLuckyInfoChanged {
	scLuckyInfoChanged := &uipb.SCLuckyInfoChanged{}

	scLuckyInfoChanged.Info = buildLuckyInfo(expire, itemId)
	return scLuckyInfoChanged
}

func BuildSCLuckyInfoNotice(objMap map[itemtypes.ItemType]map[itemtypes.ItemSubType]*playerlucky.PlayerLuckyObject) *uipb.SCLuckyInfoNotice {
	scLuckyInfoNotice := &uipb.SCLuckyInfoNotice{}
	for _, subMap := range objMap {
		for _, obj := range subMap {
			expire := obj.GetExpireTime()
			itemId := obj.GetItemId()
			scLuckyInfoNotice.InfoList = append(scLuckyInfoNotice.InfoList, buildLuckyInfo(expire, itemId))
		}
	}
	return scLuckyInfoNotice
}

func buildLuckyInfo(expire int64, itemId int32) *uipb.LuckyInfo {
	luckInfo := &uipb.LuckyInfo{}

	luckInfo.ExpireTime = &expire
	luckInfo.ItemId = &itemId

	return luckInfo
}

// expire := obj.GetExpireTime()
// typ := int32(typ)
// subType := subType.SubType()
