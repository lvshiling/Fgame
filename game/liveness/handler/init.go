package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LIVENESS_GET_TYPE), (*uipb.CSLivenessGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIVENESS_GET_TYPE), (*uipb.SCLivenessGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIVENESS_NUM_CHANGED_TYPE), (*uipb.SCLivenessNumChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_LIVENESS_OPEN_TYPE), (*uipb.CSLivenessOpen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_LIVENESS_OPEN_TYPE), (*uipb.SCLivenessOpen)(nil))
}
