package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/tulong/tulong"
)

func BuildSCRankList(dataList []*tulong.TuLongRankData, pos int32) *uipb.SCTuLongRank {
	tuLongRank := &uipb.SCTuLongRank{}
	tuLongRank.Pos = &pos
	for index, dataObj := range dataList {
		pos := int32(index + 1)
		tuLongRank.RankList = append(tuLongRank.RankList, buildRank(dataObj, pos))
	}
	return tuLongRank
}

func BuildSCTuLongStart() *uipb.SCTuLongStart {
	tuLongStart := &uipb.SCTuLongStart{}
	return tuLongStart
}

func buildRank(data *tulong.TuLongRankData, pos int32) *uipb.TuLongRank {
	tuLongRank := &uipb.TuLongRank{}
	tuLongRank.Pos = &pos

	serverId := data.GetServerId()
	allianceName := data.GetAllianceName()
	num := data.GetKillNum()

	tuLongRank.ServerId = &serverId
	tuLongRank.Name = &allianceName
	tuLongRank.Num = &num
	return tuLongRank
}
