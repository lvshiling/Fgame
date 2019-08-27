package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	activitytypes "fgame/fgame/game/activity/types"
)

func BuildSCActivityAttend(activityType activitytypes.ActivityType, isSuccess bool) *uipb.SCActivityAttend {
	scActivityAttend := &uipb.SCActivityAttend{}
	typ := int32(activityType)
	scActivityAttend.ActiveId = &typ
	scActivityAttend.IsSuccess = &isSuccess

	return scActivityAttend
}

func BuildSCActivityCollectInfoNotice(activtyType int32, countMap map[int32]int32) *uipb.SCActivityCollectInfoNotice {
	scMsg := &uipb.SCActivityCollectInfoNotice{}
	scMsg.ActiveId = &activtyType
	scMsg.InfoList = buildActivityCollectInfo(countMap)
	return scMsg
}

func buildActivityCollectInfo(countMap map[int32]int32) (infoList []*uipb.ActivityCollectInfo) {
	for key, val := range countMap {
		npcId := key
		count := val
		info := &uipb.ActivityCollectInfo{}
		info.NpcId = &npcId
		info.Count = &count
		infoList = append(infoList, info)
	}
	return infoList
}
