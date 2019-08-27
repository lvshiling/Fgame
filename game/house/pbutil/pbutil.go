package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	"fgame/fgame/game/house/house"
	playerhouse "fgame/fgame/game/house/player"
	propertypbutil "fgame/fgame/game/property/pbutil"
	propertytypes "fgame/fgame/game/property/types"
)

func BuildSCHouseListGet(houseMap map[int32]*playerhouse.PlayerHouseObject, logList []*house.HouseLogObject) *uipb.SCHouseListGet {
	scMsg := &uipb.SCHouseListGet{}
	scMsg.HouseList = buildHouseInfoList(houseMap)
	scMsg.LogList = buildHouseLogList(logList)
	return scMsg
}

func BuildSCHouseActivate(houseType, houseIndex int32, logList []*house.HouseLogObject) *uipb.SCHouseActivate {
	scMsg := &uipb.SCHouseActivate{}
	scMsg.HouseIndex = &houseIndex
	scMsg.HouseType = &houseType
	scMsg.LogList = buildHouseLogList(logList)
	return scMsg
}

func BuildSCHouseUpgrade(houseIndex int32, rewItemsMap map[int32]int32, logList []*house.HouseLogObject) *uipb.SCHouseUpgrade {
	scMsg := &uipb.SCHouseUpgrade{}
	scMsg.HouseIndex = &houseIndex
	scMsg.LogList = buildHouseLogList(logList)
	scMsg.DropInfo = droppbutil.BuildSimpleDropInfoList(rewItemsMap) //通用物品展示
	return scMsg
}

func BuildSCHouseSell(houseIndex int32, rd *propertytypes.RewData, logList []*house.HouseLogObject) *uipb.SCHouseSell {
	scMsg := &uipb.SCHouseSell{}
	scMsg.HouseIndex = &houseIndex
	scMsg.LogList = buildHouseLogList(logList)
	scMsg.RewInfo = propertypbutil.BuildRewProperty(rd) //通用资源展示
	return scMsg
}

func BuildSCHouseReceiveRent(houseIndex int32, rd *propertytypes.RewData) *uipb.SCHouseReceiveRent {
	scMsg := &uipb.SCHouseReceiveRent{}
	scMsg.HouseIndex = &houseIndex
	scMsg.RewInfo = propertypbutil.BuildRewProperty(rd)
	return scMsg
}

func BuildSCHouseRepair(houseIndex int32) *uipb.SCHouseRepair {
	scMsg := &uipb.SCHouseRepair{}
	scMsg.HouseIndex = &houseIndex
	return scMsg
}

func BuildSCHouseLogIncr(logList []*house.HouseLogObject) *uipb.SCHouseLogIncr {
	scMsg := &uipb.SCHouseLogIncr{}
	scMsg.LogList = buildHouseLogList(logList)
	return scMsg
}

func buildHouseLogList(logObjList []*house.HouseLogObject) (logList []*uipb.HouseLog) {
	for _, logObj := range logObjList {
		createTime := logObj.GetCreateTime()
		plName := logObj.GetPlayerName()
		houseIndex := logObj.GetHouseIndex()
		houseLevel := logObj.GetHouseLevel()
		houseType := logObj.GetHouseType()
		operateType := logObj.GetHouseOperateType()

		info := &uipb.HouseLog{}
		info.PlayerName = &plName
		info.CreateTime = &createTime
		info.HouseIndex = &houseIndex
		info.HouseType = &houseType
		info.HouseLevel = &houseLevel
		info.OperateType = &operateType

		logList = append(logList, info)
	}

	return logList
}

func buildHouseInfoList(houseMap map[int32]*playerhouse.PlayerHouseObject) (infoList []*uipb.HouseInfo) {
	for _, house := range houseMap {
		info := &uipb.HouseInfo{}
		houseIndex := house.GetHouseIndex()
		isBroken := house.GetIsBroken()
		houseType := int32(house.GetHouseType())
		uplevelTimes := house.GetDayTimes()
		level := house.GetHouseLevel()
		maxLevel := house.GetHouseMaxLevel()
		isRent := house.GetIsRent()

		info.HouseIndex = &houseIndex
		info.IsBroken = &isBroken
		info.HouseType = &houseType
		info.Level = &level
		info.MaxLevel = &maxLevel
		info.UplevelTimes = &uplevelTimes
		info.IsRent = &isRent

		infoList = append(infoList, info)
	}

	return infoList
}
