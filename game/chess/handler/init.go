package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHESS_INFO_GET_TYPE), (*uipb.CSChessInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHESS_INFO_GET_TYPE), (*uipb.SCChessInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHESS_ATTEND_TYPE), (*uipb.CSChessAttend)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHESS_ATTEND_TYPE), (*uipb.SCChessAttend)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHESS_ATTEND_BATCH_TYPE), (*uipb.CSChessAttendBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHESS_ATTEND_BATCH_TYPE), (*uipb.SCChessAttendBatch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHESS_CHANGED_TYPE), (*uipb.CSChessChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHESS_CHANGED_TYPE), (*uipb.SCChessChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_CHESS_LOG_INCR_TYPE), (*uipb.CSChessLogIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_CHESS_LOG_INCR_TYPE), (*uipb.SCChessLogIncr)(nil))
}
