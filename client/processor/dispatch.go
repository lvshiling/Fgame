package processor

import (
	"fgame/fgame/common/codec"
	"fgame/fgame/common/dispatch"
)

var (
	d = dispatch.NewDispatch()
)

func Register(msgType codec.MessageType, h dispatch.Handler) {
	d.Register(msgType, h)
}

func GetDispatch() dispatch.Dispatch {
	return d
}
