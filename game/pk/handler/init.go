package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"

	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PK_STATE_SWITCH_TYPE), (*uipb.CSPkStateSwitch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PK_STATE_SWITCH_TYPE), (*uipb.SCPkStateSwitch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PK_VALUE_CHANGED_TYPE), (*uipb.SCPKValueChanged)(nil))
}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_PK_STATE_SWITCH_TYPE))
}
