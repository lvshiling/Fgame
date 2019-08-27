package proxy

import (
	logintypes "fgame/fgame/account/login/types"
	logservermodel "fgame/fgame/logserver/model"
	logserverpb "fgame/fgame/logserver/pb"
	"fgame/fgame/logserver/pbutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func handleLog(msg *logserverpb.LogMessage) (err error) {
	m, err := pbutil.ConvertFromLogMessage(msg)
	if err != nil {
		log.WithFields(
			log.Fields{
				"err": err,
			}).Error("handler:玩家日志转换,失败")
		return
	}
	chatContent, ok := m.(*logservermodel.ChatContent)
	if !ok {
		log.WithFields(
			log.Fields{}).Warn("handler:聊天内容无效")
		return
	}

	sdkType := logintypes.SDKType(chatContent.SdkType)
	if !sdkType.Valid() {
		log.WithFields(
			log.Fields{
				"sdk": chatContent.SdkType,
			}).Warn("handler:sdk,无效")
		return
	}
	h := GetChatProxyHandler(sdkType)
	if h == nil {
		log.WithFields(
			log.Fields{}).Warn("handler:sdk处理器不存在")
		return
	}
	deviceType := logintypes.DevicePlatformType(chatContent.DeviceType)
	sdkUserId := chatContent.SdkUserId
	h.ProxyChat(
		sdkType,
		deviceType,
		chatContent.ServerId,
		sdkUserId,
		chatContent.UserId,
		chatContent.Name,
		chatContent.RecvId,
		chatContent.RecvName,
		GetChatType(chatContent.Channel),
		chatContent.LogTime,
		string(chatContent.Content))
	return nil
}

var (
	chatTypeMap = map[int32]string{
		0: "世界",
		1: "帮派",
		2: "队伍",
		3: "系统",
		4: "私聊",
		5: "答题",
	}
)

func GetChatType(typ int32) string {
	return chatTypeMap[typ]
}

type ChatProxyHandler interface {
	ProxyChat(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, sdkUserId string, userId int64, playerName string, targetPlayerId int64, targetPlayerName string, chatType string, chatTime int64, body string)
}

type ChatProxyHandlerFunc func(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, sdkUserId string, userId int64, playerName string, targetPlayerId int64, targetPlayerName string, chatType string, chatTime int64, body string)

func (f ChatProxyHandlerFunc) ProxyChat(sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, serverId int32, sdkUserId string, userId int64, playerName string, targetPlayerId int64, targetPlayerName string, chatType string, chatTime int64, body string) {
	f(sdkType, deviceType, serverId, sdkUserId, userId, playerName, targetPlayerId, targetPlayerName, chatType, chatTime, body)
}

var (
	proxyHandlerMap = make(map[logintypes.SDKType]ChatProxyHandler)
)

func RegisterChatProxyHandler(sdkType logintypes.SDKType, h ChatProxyHandler) {
	_, ok := proxyHandlerMap[sdkType]
	if ok {
		panic(fmt.Errorf("重复注册%s处理器", sdkType.String()))
	}
	proxyHandlerMap[sdkType] = h
}

func GetChatProxyHandler(sdkType logintypes.SDKType) ChatProxyHandler {
	h, ok := proxyHandlerMap[sdkType]
	if !ok {
		return nil
	}
	return h
}
