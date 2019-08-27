package grpc_pbutil

import (
	tulongpb "fgame/fgame/cross/tulong/pb"
	"fgame/fgame/cross/tulong/tulong"
)

func BuildTuLongRankInfoList(dataList []*tulong.TuLongRankObject) (rankList []*tulongpb.TuLongRankInfo) {
	rankList = make([]*tulongpb.TuLongRankInfo, 0, len(dataList))
	for _, data := range dataList {
		rankList = append(rankList, buildRank(data))
	}
	return rankList
}

func buildRank(data *tulong.TuLongRankObject) *tulongpb.TuLongRankInfo {
	rankInfo := &tulongpb.TuLongRankInfo{}
	rankInfo.ServerId = data.ServerId
	rankInfo.AllianceId = data.AllianceId
	rankInfo.AllianceName = data.AllianceName
	rankInfo.KillNum = data.KillNum
	return rankInfo
}
