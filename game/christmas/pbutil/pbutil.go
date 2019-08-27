package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCChristmasCollectRefreshBroadcast(groupId int32) *uipb.SCChristmasCollectRefreshBroadcast {
	scMsg := &uipb.SCChristmasCollectRefreshBroadcast{}
	scMsg.GroupId = &groupId
	return scMsg
}
func BuildSCChristmasCollectNumNotice(num int32) *uipb.SCChristmasCollectNumNotice {
	scMsg := &uipb.SCChristmasCollectNumNotice{}
	scMsg.CollectNum = &num
	return scMsg
}
