package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
)

func init() {
	dispatch.Register(codec.MessageType(uipb.MessageType_SC_SERVER_LIST_TYPE), dispatch.HandlerFunc(handleServerList))

}
