package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_INFO_TYPE), (*uipb.CSShangguzhilingInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHANGGUZHILING_INFO_TYPE), (*uipb.SCShangguzhilingInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_UPLEVEL_TYPE), (*uipb.CSShangguzhilingUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHANGGUZHILING_UPLEVEL_TYPE), (*uipb.SCShangguzhilingUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_LINGWEN_UPLEVEL_TYPE), (*uipb.CSShangguzhilingLingWenUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHANGGUZHILING_LINGWEN_UPLEVEL_TYPE), (*uipb.SCShangguzhilingLingWenUplevel)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_UPRANK_TYPE), (*uipb.CSShangguzhilingUpRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHANGGUZHILING_UPRANK_TYPE), (*uipb.SCShangguzhilingUpRank)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_LINGLIAN_TYPE), (*uipb.CSShangguzhilingLingLian)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHANGGUZHILING_LINGLIAN_TYPE), (*uipb.SCShangguzhilingLingLian)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHANGGUZHILING_RECEIVE_TYPE), (*uipb.CSShangguzhilingReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHANGGUZHILING_RECEIVE_TYPE), (*uipb.SCShangguzhilingReceive)(nil))
}
