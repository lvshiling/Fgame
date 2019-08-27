package handler

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/logic"
	"fgame/fgame/chatproxy/proxy"
)

func init() {
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeWenJian, proxy.ChatProxyHandlerFunc(wenJianProxy))
}

func wenJianProxy(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, sdkUserId string, userId int64, playerName string, targetPlayerId int64, targetPlayerName string, chatType string, chatTime int64, body string) {
	sdkType = logintypes.SDKTypeTianShen
	logic.YouQuProxy(sdkType, deviceType, serverId, sdkUserId, userId, playerName, targetPlayerId, targetPlayerName, chatType, chatTime, body)
}
