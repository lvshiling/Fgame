package handler

import (
	"fgame/fgame/client/processor"
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_SC_FUNC_OPEN_LIST_TYPE), dispatch.HandlerFunc(handleFuncOpenList))
}

func handleFuncOpenList(s session.Session, msg interface{}) (err error) {
	return
}
