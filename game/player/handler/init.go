package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_PROPERTY_TYPE), (*uipb.SCPlayerProperty)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_BASIC_INFO_GET_TYPE), (*uipb.CSPlayerBasicInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_BASIC_INFO_GET_TYPE), (*uipb.SCPlayerBasicInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_BASIC_INFO_BATCH_GET_TYPE), (*uipb.CSPlayerBasicInfoBatchGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_BASIC_INFO_BATCH_GET_TYPE), (*uipb.SCPlayerBasicInfoBatchGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_INFO_TYPE), (*uipb.SCPlayerInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_INFO_GET_TYPE), (*uipb.CSPlayerInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_INFO_GET_TYPE), (*uipb.SCPlayerInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_OPEN_VIDEO_TYPE), (*uipb.CSPlayerOpenVedio)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_OPEN_VIDEO_TYPE), (*uipb.SCPlayerOpenVedio)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_SEX_CHANGED_TYPE), (*uipb.SCPlayerSexChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_NAME_CHANGED_TYPE), (*uipb.SCPlayerNameChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_COUNT_DATA_TYPE), (*uipb.CSPlayerCountData)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_COUNT_DATA_TYPE), (*uipb.SCPlayerCountData)(nil))

}
