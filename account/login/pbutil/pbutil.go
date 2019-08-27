package pbutil

import (
	"fgame/fgame/account/player/player"
	"fgame/fgame/account/serverlist/serverlist"
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildPlayerInfo(info *player.PlayerInfo) *uipb.ServerPlayerInfo {
	playerInfo := &uipb.ServerPlayerInfo{}
	userId := info.UserId
	serverId := info.ServerId
	playerId := info.PlayerId
	role := info.Role
	sex := info.Sex
	level := info.Level
	zhuanShu := info.ZhuanShu
	playerInfo.UserId = &userId
	playerInfo.ServerId = &serverId
	playerInfo.PlayerId = &playerId
	playerInfo.Role = &role
	playerInfo.Sex = &sex
	playerInfo.Level = &level
	playerInfo.ZhuanShu = &zhuanShu
	return playerInfo
}

func BuildPlayerInfoList(infoList []*player.PlayerInfo) []*uipb.ServerPlayerInfo {
	playerInfoList := make([]*uipb.ServerPlayerInfo, 0, len(infoList))
	for _, info := range infoList {
		playerInfoList = append(playerInfoList, BuildPlayerInfo(info))
	}
	return playerInfoList
}
func BuildServerInfo(info *serverlist.ServerInfo) *uipb.ServerInfo {
	serverInfo := &uipb.ServerInfo{}
	id := info.Id
	name := info.Name
	ip := info.Ip
	tag := info.Tag
	status := info.Status

	serverInfo.Id = &id
	serverInfo.Name = &name
	serverInfo.Ip = &ip
	serverInfo.Tag = &tag
	serverInfo.Status = &status
	return serverInfo
}

func BuildServerInfoList(infoList []*serverlist.ServerInfo) []*uipb.ServerInfo {
	serverInfoList := make([]*uipb.ServerInfo, 0, len(infoList))
	for _, info := range infoList {
		serverInfoList = append(serverInfoList, BuildServerInfo(info))
	}
	return serverInfoList
}

func BuildSCAccountLogin(userId int64, token string, expiredTime int64, playerInfoList []*player.PlayerInfo, serverInfoList []*serverlist.ServerInfo, notice string, platformUserId string) *uipb.SCAccountLogin {
	scAccountLogin := &uipb.SCAccountLogin{}
	scAccountLogin.UserId = &userId
	scAccountLogin.Token = &token
	scAccountLogin.ExpiredTime = &expiredTime
	scAccountLogin.ServerPlayerInfoList = BuildPlayerInfoList(playerInfoList)
	scAccountLogin.ServerList = BuildServerInfoList(serverInfoList)
	scAccountLogin.Notice = &notice
	scAccountLogin.PlatformUserId = &platformUserId
	return scAccountLogin
}
