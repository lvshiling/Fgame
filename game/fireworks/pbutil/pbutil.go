package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCFireworksBroadcast(itemId int32, num int32) *uipb.SCFireWorksBroadcast {
	scFireWorksBroadcast := &uipb.SCFireWorksBroadcast{}
	scFireWorksBroadcast.ItemId = &itemId
	scFireWorksBroadcast.Num = &num
	return scFireWorksBroadcast
}
