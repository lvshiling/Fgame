package tulong

import (
	tulongpb "fgame/fgame/cross/tulong/pb"
)

type TuLongRankData struct {
	serverId     int32
	allianceId   int64
	allianceName string
	killNum      int32
}

func (t *TuLongRankData) GetServerId() int32 {
	return t.serverId
}

func (t *TuLongRankData) GetAllianceId() int64 {
	return t.allianceId
}

func (t *TuLongRankData) GetAllianceName() string {
	return t.allianceName
}

func (t *TuLongRankData) GetKillNum() int32 {
	return t.killNum
}

func convertFromRankInfo(rankInfo *tulongpb.TuLongRankInfo) *TuLongRankData {
	rankData := &TuLongRankData{}
	rankData.serverId = rankInfo.ServerId
	rankData.allianceId = rankInfo.AllianceId
	rankData.allianceName = rankInfo.AllianceName
	rankData.killNum = rankInfo.KillNum
	return rankData
}

func convertFromRankInfoList(rankInfoList []*tulongpb.TuLongRankInfo) (dataList []*TuLongRankData) {
	dataList = make([]*TuLongRankData, 0, len(rankInfoList))
	for _, rankInfo := range rankInfoList {
		dataList = append(dataList, convertFromRankInfo(rankInfo))
	}
	return dataList
}
