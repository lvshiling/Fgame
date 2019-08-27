package player

import (
	centerpb "fgame/fgame/center/pb"
)

func convertFromPlayerInfo(info *centerpb.PlayerServerInfo) *PlayerInfo {
	playerInfo := &PlayerInfo{}
	playerInfo.UserId = info.UserId
	playerInfo.PlayerId = info.PlayerId
	playerInfo.ServerId = info.ServerId
	playerInfo.Level = info.Level
	playerInfo.Role = info.Role
	playerInfo.ZhuanShu = info.ZhuanShu
	playerInfo.Sex = info.Sex
	return playerInfo
}

func convertFromPlayerInfoList(infoList []*centerpb.PlayerServerInfo) []*PlayerInfo {
	playerInfoList := make([]*PlayerInfo, 0, len(infoList))
	for _, info := range infoList {
		playerInfoList = append(playerInfoList, convertFromPlayerInfo(info))
	}
	return playerInfoList
}
