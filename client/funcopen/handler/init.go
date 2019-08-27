package handler

import (
	clientcodec "fgame/fgame/client/codec"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func init() {
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUNC_OPEN_LIST_TYPE), (*uipb.SCFuncOpenList)(nil))
	clientcodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUNC_OPEN_UPDATE_LIST_TYPE), (*uipb.SCFuncOpenUpdateList)(nil))
}
