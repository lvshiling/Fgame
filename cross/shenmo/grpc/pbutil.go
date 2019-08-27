package grpc_pbutil

import (
	shenmopb "fgame/fgame/cross/shenmo/pb"
	"fgame/fgame/cross/shenmo/shenmo"
)

func BuildShenMoRankInfoList(thisDataList []*shenmo.ShenMoRankObject, thisRankTime int64,
	lastDataList []*shenmo.ShenMoRankObject, lastRankTime int64) (thisRankData *shenmopb.ShenMoRanInfoData, lastRankData *shenmopb.ShenMoRanInfoData) {

	thisRankData = &shenmopb.ShenMoRanInfoData{}
	thisRankData.RankInfoList = make([]*shenmopb.ShenMoRankInfo, 0, len(thisDataList))
	thisRankData.RankTime = thisRankTime
	for _, data := range thisDataList {
		if data.JiFenNum == 0 {
			continue
		}
		thisRankData.RankInfoList = append(thisRankData.RankInfoList, buildRank(data, true))
	}

	lastRankData = &shenmopb.ShenMoRanInfoData{}
	lastRankData.RankInfoList = make([]*shenmopb.ShenMoRankInfo, 0, len(lastDataList))
	lastRankData.RankTime = lastRankTime
	for _, data := range lastDataList {
		if data.LastJiFenNum == 0 {
			continue
		}
		lastRankData.RankInfoList = append(lastRankData.RankInfoList, buildRank(data, false))
	}
	return thisRankData, lastRankData
}

func buildRank(data *shenmo.ShenMoRankObject, isThis bool) *shenmopb.ShenMoRankInfo {
	rankInfo := &shenmopb.ShenMoRankInfo{}
	rankInfo.ServerId = data.ServerId
	rankInfo.AllianceId = data.AllianceId
	rankInfo.AllianceName = data.AllianceName
	jiFenNum := int32(0)
	if isThis {
		jiFenNum = data.JiFenNum
	} else {
		jiFenNum = data.LastJiFenNum
	}
	rankInfo.JiFenNum = jiFenNum

	return rankInfo
}
