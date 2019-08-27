package logic

import (
	"fgame/fgame/cross/player/player"
	"fgame/fgame/game/chat/pbutil"
	chattypes "fgame/fgame/game/chat/types"
)

//系统全服广播
func SystemBroadcast(msgType chattypes.MsgType, content []byte) {
	chatRecv := pbutil.BuildSCChatRecv(int64(0), chattypes.ChannelTypeSystem, int64(0), msgType, content)
	player.GetOnlinePlayerManager().BroadcastMsg(chatRecv)

}
