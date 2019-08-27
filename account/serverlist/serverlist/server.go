package serverlist

import (
	centerpb "fgame/fgame/center/pb"
)

type ServerInfo struct {
	Id     int32
	Name   string
	Ip     string
	Tag    int32
	Status int32
}

func convertFromServerInfo(info *centerpb.ServerInfo) *ServerInfo {
	serverInfo := &ServerInfo{}
	serverInfo.Id = info.Id
	serverInfo.Name = info.Name
	serverInfo.Ip = info.Ip
	serverInfo.Tag = info.Tag
	serverInfo.Status = info.Status
	return serverInfo
}

func convertFromServerInfoList(infoList []*centerpb.ServerInfo) []*ServerInfo {
	serverInfoList := make([]*ServerInfo, 0, len(infoList))
	for _, info := range infoList {
		serverInfoList = append(serverInfoList, convertFromServerInfo(info))
	}
	return serverInfoList
}
