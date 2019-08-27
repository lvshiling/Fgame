package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	chattypes "fgame/fgame/game/chat/types"
	hongbao "fgame/fgame/game/hongbao/hongbao"
)

func buildItem(award *hongbao.AwardInfo) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemId := award.ItemId
	num := award.ItemCnt
	level := award.Level

	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	itemInfo.Level = &level
	return itemInfo
}

func buildSnatcher(snatcher *hongbao.SnatcherInfo) *uipb.SnatcherInfo {
	snatcherInfo := &uipb.SnatcherInfo{}
	name := snatcher.Name
	role := int32(snatcher.Role)
	sex := int32(snatcher.Sex)
	level := snatcher.Level
	playerId := snatcher.PlayerId

	snatcherInfo.Name = &name
	snatcherInfo.Role = &role
	snatcherInfo.Sex = &sex
	snatcherInfo.Level = &level
	snatcherInfo.PlayerId = &playerId
	return snatcherInfo
}

func BuildSCHongBaoGet(hongbaoObj *hongbao.HongBaoObject, endTime int64, channelType chattypes.ChannelType) *uipb.SCHongbaoGet {
	scMsg := &uipb.SCHongbaoGet{}

	awardList := hongbaoObj.GetAwardList()
	snatchLog := hongbaoObj.GetSnatchLog()
	hongBaoType := int32(hongbaoObj.GetHongBaoType())
	hongBaoId := hongbaoObj.GetDBId()
	sendId := hongbaoObj.GetSendId()
	countMax := int32(len(awardList))
	for logCount := 0; logCount < len(snatchLog); logCount++ {
		scMsg.ItemList = append(scMsg.ItemList, buildItem(awardList[logCount]))
		scMsg.SnatcherList = append(scMsg.SnatcherList, buildSnatcher(snatchLog[logCount]))
	}
	channelTypeInt := int32(channelType)

	scMsg.EndTime = &endTime
	scMsg.HongBaoType = &hongBaoType
	scMsg.HongBaoId = &hongBaoId
	scMsg.SendId = &sendId
	scMsg.CountMax = &countMax
	scMsg.Channel = &channelTypeInt
	return scMsg
}

func BuildSCHongBaoSend(id int64) *uipb.SCHongbaoSend {
	scMsg := &uipb.SCHongbaoSend{}
	scMsg.HongBaoId = &id
	return scMsg
}

func BuildSCHongBaoSnatch(hongbaoObj *hongbao.HongBaoObject, endTime int64, result int32, snatchCount int32) *uipb.SCHongbaoSnatch {
	scMsg := &uipb.SCHongbaoSnatch{}

	awardList := hongbaoObj.GetAwardList()
	snatchLog := hongbaoObj.GetSnatchLog()
	hongBaoType := int32(hongbaoObj.GetHongBaoType())
	hongBaoId := hongbaoObj.GetDBId()
	sendId := hongbaoObj.GetSendId()
	countMax := int32(len(awardList))
	for logCount := 0; logCount < len(snatchLog); logCount++ {
		scMsg.ItemList = append(scMsg.ItemList, buildItem(awardList[logCount]))
		scMsg.SnatcherList = append(scMsg.SnatcherList, buildSnatcher(snatchLog[logCount]))
	}

	scMsg.EndTime = &endTime
	scMsg.HongBaoType = &hongBaoType
	scMsg.Result = &result
	scMsg.HongBaoId = &hongBaoId
	scMsg.SendId = &sendId
	scMsg.SnatchCount = &snatchCount
	scMsg.CountMax = &countMax
	return scMsg
}

func BuildSCHongBaoSnatchGet(count int32) *uipb.SCHongbaoSnatchGet {
	scMsg := &uipb.SCHongbaoSnatchGet{}
	scMsg.SnatchCount = &count
	return scMsg
}
