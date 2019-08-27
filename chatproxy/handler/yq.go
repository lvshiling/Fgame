package handler

import (
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/chatproxy/logic"
	"fgame/fgame/chatproxy/proxy"
)

func init() {
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeTianXing, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeTianJi, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeYaoJing, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeXiaKeXing, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeSanJie, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeYouMeng, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeJiuMeng, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeLongYu, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeXianFan, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeMengHuan, proxy.ChatProxyHandlerFunc(sanJieProxy))
	proxy.RegisterChatProxyHandler(logintypes.SDKTypeXiaoYao, proxy.ChatProxyHandlerFunc(sanJieProxy))

}

func sanJieProxy(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, sdkUserId string, userId int64, playerName string, targetPlayerId int64, targetPlayerName string, chatType string, chatTime int64, body string) {
	sdkType = logintypes.SDKTypeSanJie
	logic.YouQuProxy(sdkType, deviceType, serverId, sdkUserId, userId, playerName, targetPlayerId, targetPlayerName, chatType, chatTime, body)
}
