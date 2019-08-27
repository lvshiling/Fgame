package handler

import (
	"fgame/fgame/client/processor"
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
)

func init() {
	processor.Register(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_TYPE), dispatch.HandlerFunc(handleBuff))
}

func handleBuff(s session.Session, msg interface{}) (err error) {
	return
}
