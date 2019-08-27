package center

import (
	centerpb "fgame/fgame/center/pb"
	centertypes "fgame/fgame/center/types"
)

func ConvertFromCrossServerInfo(info *centerpb.CrossServerInfo) *ServerInfo {
	tempInfo := &ServerInfo{}
	tempInfo.serverIp = info.GetServerIp()
	tempInfo.serverPort = info.GetServerPort()
	tempInfo.serverType = centertypes.GameServerType(info.GetServerType())
	return tempInfo
}

func ConvertFromCrossServerInfoList(infoList []*centerpb.CrossServerInfo) []*ServerInfo {
	tempList := make([]*ServerInfo, 0, len(infoList))
	for _, info := range infoList {
		tempList = append(tempList, ConvertFromCrossServerInfo(info))
	}
	return tempList
}
