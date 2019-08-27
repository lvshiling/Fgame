package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	propertypbutil "fgame/fgame/game/property/pbutil"
	propertytypes "fgame/fgame/game/property/types"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	xianzuncardtypes "fgame/fgame/game/xianzuncard/types"
)

func BuildSCXianZunCardInfo(xianZunCardObjMap map[xianzuncardtypes.XianZunCardType]*playerxianzuncard.PlayerXianZunCardObject) *uipb.SCXianZunCardInfo {
	scXianZunCardInfo := &uipb.SCXianZunCardInfo{}
	for typ, xianZunCardObj := range xianZunCardObjMap {
		if !xianZunCardObj.IsActivite() {
			continue
		}
		xianZunCardInfo := &uipb.XianZunCardInfo{}
		activiteTime := xianZunCardObj.GetActiviteTime()
		isReceive := xianZunCardObj.GetIsReceive()
		t := int32(typ)
		xianZunCardInfo.Type = &t
		xianZunCardInfo.ActiviteTime = &activiteTime
		xianZunCardInfo.IsReceive = &isReceive
		scXianZunCardInfo.XianZunCardInfo = append(scXianZunCardInfo.XianZunCardInfo, xianZunCardInfo)
	}
	return scXianZunCardInfo
}

func BuildSCXianZunCardBuy(typ int32, rd *propertytypes.RewData, receiveItem map[int32]int32) *uipb.SCXianZunCardBuy {
	scXianZunCardBuy := &uipb.SCXianZunCardBuy{}
	scXianZunCardBuy.Type = &typ
	scXianZunCardBuy.RewInfo = propertypbutil.BuildRewProperty(rd)
	for itemId, num := range receiveItem {
		scXianZunCardBuy.DropInfo = append(scXianZunCardBuy.DropInfo, buildDropInfo(itemId, num, 0))
	}

	return scXianZunCardBuy
}

func BuildSCReceiveXianZunCardReward(typ int32, rd *propertytypes.RewData, receiveItem map[int32]int32) *uipb.SCReceiveXianZunCardReward {
	scReceiveXianZunCardReward := &uipb.SCReceiveXianZunCardReward{}
	scReceiveXianZunCardReward.Type = &typ
	scReceiveXianZunCardReward.RewInfo = propertypbutil.BuildRewProperty(rd)
	for itemId, num := range receiveItem {
		scReceiveXianZunCardReward.DropInfo = append(scReceiveXianZunCardReward.DropInfo, buildDropInfo(itemId, num, 0))
	}

	return scReceiveXianZunCardReward
}

func BuildSCXianZunCardNotice(typ int32) *uipb.SCXianZunCardNotice {
	scXianZunCardNotice := &uipb.SCXianZunCardNotice{}
	scXianZunCardNotice.Type = &typ
	return scXianZunCardNotice
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level
	return dropInfo
}
