package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/utils"
	playermingge "fgame/fgame/game/mingge/player"
	minggetypes "fgame/fgame/game/mingge/types"
)

func BuildSCMingGePanGet(mingGePanMap map[minggetypes.MingGeType]map[minggetypes.MingGeAllSubType]*playermingge.PlayerMingGePanObject) *uipb.SCMingGePanGet {
	scMingGePanGet := &uipb.SCMingGePanGet{}
	for mingGeType, mingGeAllSubTypeMap := range mingGePanMap {
		scMingGePanGet.MingGePanList = append(scMingGePanGet.MingGePanList, buildMingGePanInfo(int32(mingGeType), mingGeAllSubTypeMap))
	}
	return scMingGePanGet
}

func buildMingGePanInfo(panType int32, mingGeAllSubTypeMap map[minggetypes.MingGeAllSubType]*playermingge.PlayerMingGePanObject) *uipb.MingGePanInfo {
	mingGePanInfo := &uipb.MingGePanInfo{}
	mingGePanInfo.PanType = &panType
	for subType, obj := range mingGeAllSubTypeMap {
		mingGePanInfo.MingGePanList = append(mingGePanInfo.MingGePanList, buildMingGeInfo(int32(subType), obj))
	}
	return mingGePanInfo
}

func buildMingGeInfo(subType int32, obj *playermingge.PlayerMingGePanObject) *uipb.MingGeInfo {
	mingGeInfo := &uipb.MingGeInfo{}
	mingGeInfo.MingGeType = &subType
	for slot, itemId := range obj.GetMingPanItemMap() {
		mingGeInfo.ItemList = append(mingGeInfo.ItemList, buildMingGeItem(int32(slot), itemId))
	}
	return mingGeInfo
}

func buildMingGeItem(slot int32, itemId int32) *uipb.MingGeItemInfo {
	mingGeItemInfo := &uipb.MingGeItemInfo{}
	mingGeItemInfo.Slot = &slot
	mingGeItemInfo.ItemId = &itemId
	return mingGeItemInfo
}

func BuildSCMingGeMingLiGet(mingLiMap map[minggetypes.MingGongType]map[minggetypes.MingGongAllSubType]*playermingge.PlayerMingLiObject) *uipb.SCMingGeMingLiGet {
	scMingGeMingLiGet := &uipb.SCMingGeMingLiGet{}
	for mingGongType, mingGongAllSubTypeMap := range mingLiMap {
		scMingGeMingLiGet.MingLiList = append(scMingGeMingLiGet.MingLiList, buildMingGeMingInfo(int32(mingGongType), mingGongAllSubTypeMap))
	}
	return scMingGeMingLiGet
}

func buildMingGeMingInfo(mingGongType int32, mingGongAllSubTypeMap map[minggetypes.MingGongAllSubType]*playermingge.PlayerMingLiObject) *uipb.MingGeMingLiInfo {
	mingGeMingLiInfo := &uipb.MingGeMingLiInfo{}
	mingGeMingLiInfo.MingGongType = &mingGongType
	for subType, obj := range mingGongAllSubTypeMap {
		mingGeMingLiInfo.MingLiList = append(mingGeMingLiInfo.MingLiList, buildMingLiInfo(int32(subType), obj))
	}
	return mingGeMingLiInfo
}

func buildMingLiInfo(subType int32, obj *playermingge.PlayerMingLiObject) *uipb.MingLiInfo {
	mingLiInfo := &uipb.MingLiInfo{}
	mingLiInfo.PosTag = &subType
	for _, mingLiObj := range obj.GetMingLiMap() {
		mingLiInfo.PropertyList = append(mingLiInfo.PropertyList, buildMingGeProperty(mingLiObj))
	}
	return mingLiInfo
}

func buildMingGeProperty(mingLiInfo *playermingge.MingLiInfo) *uipb.MingLiPropertyInfo {
	mingLiPropertyInfo := &uipb.MingLiPropertyInfo{}
	slot := int32(mingLiInfo.GetSlot())
	propertyType := int32(mingLiInfo.GetMingGeProperty())
	times := mingLiInfo.GetTimes()
	mingLiPropertyInfo.Slot = &slot
	mingLiPropertyInfo.PropertyType = &propertyType
	mingLiPropertyInfo.Times = &times
	return mingLiPropertyInfo
}

func BuildSCMingGeRefinedGet(refinedMap map[minggetypes.MingGeAllSubType]*playermingge.PlayerMingGeRefinedObject) *uipb.SCMingGeRefinedGet {
	scMingGeRefinedGet := &uipb.SCMingGeRefinedGet{}
	for mingGeType, refinedObj := range refinedMap {
		scMingGeRefinedGet.RefinedList = append(scMingGeRefinedGet.RefinedList, buildMingGeRefinedInfo(int32(mingGeType), refinedObj))
	}
	return scMingGeRefinedGet
}

func buildMingGeRefinedInfo(mingGeType int32, obj *playermingge.PlayerMingGeRefinedObject) *uipb.MingGeRefinedInfo {
	mingGeRefinedInfo := &uipb.MingGeRefinedInfo{}
	mingGeRefinedInfo.Type = &mingGeType
	number := obj.GetNumber()
	star := obj.GetStar()
	refinedNum := obj.GetRefinedNum()
	refinedPro := obj.GetRefinedPro()
	mingGeRefinedInfo.Number = &number
	mingGeRefinedInfo.Star = &star
	mingGeRefinedInfo.RefinedNum = &refinedNum
	mingGeRefinedInfo.RefinedPro = &refinedPro
	return mingGeRefinedInfo
}

func BuildSCMingGeMosaic(mingPanType int32, mingGeType int32, slot int32, itemId int32) *uipb.SCMingGePanMosaic {
	scMingGePanMosaic := &uipb.SCMingGePanMosaic{}
	scMingGePanMosaic.PanType = &mingPanType
	scMingGePanMosaic.MingGeType = &mingGeType
	scMingGePanMosaic.Slot = &slot
	scMingGePanMosaic.ItemId = &itemId
	return scMingGePanMosaic
}

func BuildSCMingGeUnload(panType int32, mingGeSubType int32, slot int32) *uipb.SCMingGePanUnload {
	scMingGePanUnload := &uipb.SCMingGePanUnload{}
	scMingGePanUnload.PanType = &panType
	scMingGePanUnload.MingGeType = &mingGeSubType
	scMingGePanUnload.Slot = &slot
	return scMingGePanUnload
}

func BuildSCMingGeSynthesis(sucess bool, itemId int32, num int32) *uipb.SCMingGeSynthesis {
	scMingGeSynthesis := &uipb.SCMingGeSynthesis{}
	scMingGeSynthesis.Sucess = &sucess
	scMingGeSynthesis.ItemId = &itemId
	scMingGeSynthesis.ItemNum = &num
	return scMingGeSynthesis
}

func BuildSCMingGeRefined(refinedMap map[minggetypes.MingGeAllSubType]*playermingge.PlayerMingGeRefinedObject,
	mingGeAllSubTypeMap map[minggetypes.MingGeAllSubType]bool) *uipb.SCMingGeRefined {

	scMingGeRefined := &uipb.SCMingGeRefined{}
	for mingGeType, refinedObj := range refinedMap {
		_, ok := mingGeAllSubTypeMap[mingGeType]
		if !ok {
			continue
		}
		scMingGeRefined.RefinedList = append(scMingGeRefined.RefinedList, buildMingGeRefinedInfo(int32(mingGeType), refinedObj))
	}
	return scMingGeRefined
}

func BuildSCMingGeMingGongActivate(mingLiMap map[minggetypes.MingGongType]map[minggetypes.MingGongAllSubType]*playermingge.PlayerMingLiObject,
	mingGongTypeMap map[minggetypes.MingGongType]bool) *uipb.SCMingGeMingGongActive {

	scMingGeMingGongActive := &uipb.SCMingGeMingGongActive{}
	for mingGongType, mingGongAllSubTypeMap := range mingLiMap {
		_, ok := mingGongTypeMap[mingGongType]
		if !ok {
			continue
		}
		scMingGeMingGongActive.MingGeMingLiList = append(scMingGeMingGongActive.MingGeMingLiList, buildMingGeMingInfo(int32(mingGongType), mingGongAllSubTypeMap))
	}
	return scMingGeMingGongActive
}

func BuildSCMingGeMingLiBaptize(mingGongType int32, posTag int32, obj *playermingge.PlayerMingLiObject, slotList []int32) *uipb.SCMingGeMingLiBaptize {
	scMingGeMingLiBaptize := &uipb.SCMingGeMingLiBaptize{}
	scMingGeMingLiBaptize.MingGongType = &mingGongType
	scMingGeMingLiBaptize.PosTag = &posTag

	for _, mingLiInfo := range obj.GetMingLiMap() {
		slotInt32 := int32(mingLiInfo.GetSlot())
		flag := utils.ContainInt32(slotList, slotInt32)
		if !flag {
			continue
		}
		scMingGeMingLiBaptize.PropertyList = append(scMingGeMingLiBaptize.PropertyList, buildMingGeProperty(mingLiInfo))
	}
	return scMingGeMingLiBaptize
}
