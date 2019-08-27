package center

import (
	centerpb "fgame/fgame/center/pb"
	"fgame/fgame/center/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"time"
)

func ConvertCrossServerInfo(info *ServerInfo) *centerpb.CrossServerInfo {
	crossServerInfo := &centerpb.CrossServerInfo{}
	crossServerInfo.ServerIp = info.GetIp()
	crossServerInfo.ServerPort = info.GetPort()
	crossServerInfo.ServerType = int32(info.GetServerType())
	return crossServerInfo
}

func buildServerInfo(info *ServerInfo, mergeInfo *ServerInfo) *centerpb.ServerInfo {

	serverInfo := &centerpb.ServerInfo{}
	serverInfo.Id = info.GetServerId()
	serverInfo.Name = info.GetName()
	serverInfo.Tag = int32(info.GetServerTag())
	if mergeInfo != nil {
		serverInfo.Ip = fmt.Sprintf("%s:%s", mergeInfo.GetServerIp(), mergeInfo.GetServerPort())
		if mergeInfo.GetStatus() == types.GameServerStatusMaintain {
			serverInfo.Status = int32(types.ServerStatusMaintained)
		} else {
			serverInfo.Status = int32(mergeInfo.GetServerStatus())
		}

		serverInfo.RemoteIp = fmt.Sprintf("%s:%s", mergeInfo.GetRemoteIp(), mergeInfo.GetRemotePort())

	} else {
		serverInfo.Ip = fmt.Sprintf("%s:%s", info.GetServerIp(), info.GetServerPort())
		if info.GetStatus() == types.GameServerStatusMaintain {
			serverInfo.Status = int32(types.ServerStatusMaintained)
		} else {
			serverInfo.Status = int32(info.GetServerStatus())
		}

		serverInfo.RemoteIp = fmt.Sprintf("%s:%s", info.GetRemoteIp(), info.GetRemotePort())
	}

	return serverInfo
}

func buildServerInfoList(gm int32, infoList []*ServerInfo, mergeInfoList []*ServerInfo) []*centerpb.ServerInfo {
	now := timeutils.TimeToMillisecond(time.Now())
	serverInfoList := make([]*centerpb.ServerInfo, 0, len(infoList))
	for i, info := range infoList {
		//gm忽略
		if gm == 0 {
			mergeInfo := mergeInfoList[i]
			if mergeInfo == nil {
				//首服特殊处理
				if info.serverId != 1 && info.serverId != 100 {
					//不提前预告
					tempPreShowTime := preShowTime
					if info.preShow == 0 {
						tempPreShowTime = 0
					}
					if (info.startTime - now) > tempPreShowTime {
						continue
					}
				}
			} else {
				//首服特殊处理
				if mergeInfo.serverId != 1 && info.serverId != 100 {
					//不提前预告
					tempPreShowTime := preShowTime
					if mergeInfo.preShow == 0 {
						tempPreShowTime = 0
					}
					if (mergeInfo.startTime - now) > tempPreShowTime {
						continue
					}
				}
			}
		}
		serverInfoList = append(serverInfoList, buildServerInfo(info, mergeInfoList[i]))
	}
	return serverInfoList
}
