package handler

import (
	messagetypepb "fgame/fgame/gm/gamegm/monitor/messagetype/pb"

	monitor "fgame/fgame/gm/gamegm/monitor"
)

//初始化分发器
func InitDispatcher(d *monitor.Dispatcher) {
	d.Register(int32(messagetypepb.QiPaiMessageType_CGLoginType), monitor.MessageHandlerFunc(HandleLogin))
	d.Register(int32(messagetypepb.QiPaiMessageType_CGPingType), monitor.MessageHandlerFunc(HandlePing))
}
