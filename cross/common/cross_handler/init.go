package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_HEARTBEAT_TYPE), (*crosspb.SIHeartBeat)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_HEARTBEAT_TYPE), (*crosspb.ISHeartBeat)(nil))
}
