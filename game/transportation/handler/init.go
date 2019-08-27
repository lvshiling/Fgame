package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PERSONAL_TRANSPORTATION_TYPE), (*uipb.CSPersonalTransportation)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PERSONAL_TRANSPORTATION_TYPE), (*uipb.SCPersonalTransportation)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_TRANSPORTATION_TYPE), (*uipb.CSAllianceTransportation)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_TRANSPORTATION_TYPE), (*uipb.SCAllianceTransportation)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DISTRESS_SIGNAL_TYPE), (*uipb.CSDistressSignal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DISTRESS_SIGNAL_TYPE), (*uipb.SCDistressSignal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DISTRESS_SIGNAL_BROADCAST_TYPE), (*uipb.SCDistressSignalBroadcast)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_AGREE_DISTRESS_SIGNAL_TYPE), (*uipb.CSAgreeDistressSignal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_AGREE_DISTRESS_SIGNAL_TYPE), (*uipb.SCAgreeDistressSignal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_TRANSPORT_INFO_TYPE), (*uipb.CSPlayerTransportInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_TRANSPORT_INFO_TYPE), (*uipb.SCPlayerTransportInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRANSPORT_BRIEF_INFO_NOTICE_TYPE), (*uipb.SCTransportBriefInfoNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_TRANSPORTATION_PROTECT_TYPE), (*uipb.CSTransportationProtect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRANSPORTATION_PROTECT_TYPE), (*uipb.SCTransportationProtect)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_TRANSPORTATION_PROTECT_NOTICE_TYPE), (*uipb.SCTransportationProtectNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RECEIVE_TRANSPORT_REW_TYPE), (*uipb.CSReceiveTransportRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RECEIVE_TRANSPORT_REW_TYPE), (*uipb.SCReceiveTransportRew)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ROB_SUCCESS_NOTICE_TYPE), (*uipb.SCRobSuccessNotice)(nil))
}
