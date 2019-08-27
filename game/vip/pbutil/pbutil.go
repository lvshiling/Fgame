package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCVipInfoNotice(level, star int32, chargeNum int64, record map[int32]int32, freeRecord map[int32]int32) *uipb.SCVipInfoNotice {
	scVipInfoNotice := &uipb.SCVipInfoNotice{}
	scVipInfoNotice.VipLevel = &level
	scVipInfoNotice.VipStar = &star
	num := int32(chargeNum)
	scVipInfoNotice.ChargeNum = &num

	for giftLevel, _ := range record {
		scVipInfoNotice.GiftRecordList = append(scVipInfoNotice.GiftRecordList, giftLevel)
	}
	for freeGiftLevel, _ := range freeRecord {
		scVipInfoNotice.FreeGiftRecordList = append(scVipInfoNotice.FreeGiftRecordList, freeGiftLevel)
	}

	return scVipInfoNotice
}

func BuildSCVipGiftBuy(giftLevel, star int32, itemMap map[int32]int32) *uipb.SCVipGiftBuy {
	scVipGiftBuy := &uipb.SCVipGiftBuy{}
	scVipGiftBuy.GiftLevel = &giftLevel
	scVipGiftBuy.GiftStar = &star

	for itemId, num := range itemMap {
		scVipGiftBuy.DropList = append(scVipGiftBuy.DropList, buildDropInfo(itemId, num, 0))
	}

	return scVipGiftBuy
}

func BuildSCReceiveFreeGift(giftLevel, star int32, itemMap map[int32]int32) *uipb.SCReceiveFreeGift {
	scMsg := &uipb.SCReceiveFreeGift{}
	scMsg.GiftLevel = &giftLevel
	scMsg.GiftStar = &star

	for itemId, num := range itemMap {
		scMsg.DropList = append(scMsg.DropList, buildDropInfo(itemId, num, 0))
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
