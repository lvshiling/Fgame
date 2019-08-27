package handler

import (
	monitor "fgame/fgame/gm/gamegm/monitor"
	messagetypepb "fgame/fgame/gm/gamegm/monitor/chatmonitor/pb/messagetype"
)

//初始化分发器
func InitDispatcher(d *monitor.Dispatcher) {
	d.Register(int32(messagetypepb.ChatMonitorMessageType_CGChatMinitorType), monitor.MessageHandlerFunc(HandleUserServer))
}
