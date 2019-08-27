package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	reddottypes "fgame/fgame/game/reddot/types"
)

func BuildSCActivityNoticeOnTimeAll(groupIdList []int32) *uipb.SCActivityNotice {
	scMsg := &uipb.SCActivityNotice{}
	scMsg.InfoList = buildReddotInfoLoad(groupIdList)

	return scMsg
}

func BuildSCActivityNoticeLoad(groupIdList []int32) *uipb.SCActivityNotice {
	scMsg := &uipb.SCActivityNotice{}
	scMsg.InfoList = buildReddotInfoLoad(groupIdList)

	return scMsg
}

func BuildSCActivityNoticeChanged(redList []*reddottypes.RedDotInfo) *uipb.SCActivityNotice {
	scMsg := &uipb.SCActivityNotice{}
	scMsg.InfoList = buildReddotInfoChanged(redList)

	return scMsg
}

func buildReddotInfoChanged(redList []*reddottypes.RedDotInfo) (infoList []*uipb.RedDotInfo) {
	for _, redInfo := range redList {
		groupId := redInfo.GroupId
		isRed := redInfo.IsReddot

		info := &uipb.RedDotInfo{}
		info.GroupId = &groupId
		info.IsReddot = &isRed

		infoList = append(infoList, info)
	}
	return infoList
}

func buildReddotInfoLoad(groupIdList []int32) (infoList []*uipb.RedDotInfo) {
	for _, group := range groupIdList {
		isRed := true
		groupId := group

		info := &uipb.RedDotInfo{}
		info.GroupId = &groupId
		info.IsReddot = &isRed

		infoList = append(infoList, info)
	}
	return infoList
}

func BuildSCActivityDataNotice(groupId int32, recordList []int32) *uipb.SCActivityDataNotice {
	scMsg := &uipb.SCActivityDataNotice{}
	scMsg.GroupId = &groupId
	scMsg.Record = buildRecord(recordList)

	return scMsg
}

func buildRecord(recordList []int32) *uipb.ReceiveRecord {
	info := &uipb.ReceiveRecord{}
	info.RecordList = recordList
	return info
}
