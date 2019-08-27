package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BODYSHIELD_GET_TYPE), (*uipb.CSBodyShieldGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BODYSHIELD_GET_TYPE), (*uipb.SCBodyShieldGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BODYSHIELD_JJDAN_TYPE), (*uipb.CSBodyShieldJJDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BODYSHIELD_JJDAN_TYPE), (*uipb.SCBodyShieldJJDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BODYSHIELD_ADVANCED_TYPE), (*uipb.CSBodyShieldAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BODYSHIELD_ADVANCED_TYPE), (*uipb.SCBodyShieldAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHIELD_GET_TYPE), (*uipb.CSShieldGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHIELD_GET_TYPE), (*uipb.SCShieldGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHIELD_ADVANCED_TYPE), (*uipb.CSShieldAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHIELD_ADVANCED_TYPE), (*uipb.SCShieldAdvanced)(nil))
}
