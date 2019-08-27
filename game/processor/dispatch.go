package processor

import (
	"fgame/fgame/common/codec"
	"fgame/fgame/common/dispatch"
	"fmt"
)

var (
	d      = dispatch.NewDispatch()
	crossD = dispatch.NewDispatch()
)

func RegisterCross(msgType codec.MessageType, h dispatch.Handler) {
	crossD.Register(msgType, h)
}

func GetCrossDispatch() dispatch.Dispatch {
	return crossD
}

func Register(msgType codec.MessageType, h dispatch.Handler) {
	d.Register(msgType, h)
}

func GetDispatch() dispatch.Dispatch {
	return d
}

var (
	proxyMap = make(map[codec.MessageType]bool)
)

func RegisterProxy(msgType codec.MessageType) {
	_, ok := proxyMap[msgType]
	if ok {
		panic(fmt.Errorf("重复注册代理消息%d", msgType))
	}
	proxyMap[msgType] = true
}

func IsProxy(msgType codec.MessageType) bool {
	return proxyMap[msgType]
}
