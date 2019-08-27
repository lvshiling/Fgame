package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	shenmo "fgame/fgame/game/shenmo/shenmo"
)

func BuildSCShenMoGetReward(rankTime int64) *uipb.SCShenMoGetReward {
	scShenMoGetReward := &uipb.SCShenMoGetReward{}
	scShenMoGetReward.RankTime = &rankTime
	return scShenMoGetReward
}

func BuildSCShenMoMyRank(isThis bool, pos int32, rankTime int64) *uipb.SCShenMoMyRank {
	scShenMoMyRank := &uipb.SCShenMoMyRank{}
	scShenMoMyRank.IsThis = &isThis
	scShenMoMyRank.Pos = &pos
	scShenMoMyRank.RankTime = &rankTime
	return scShenMoMyRank
}

func BuildSCShenMoRankGet(isThis bool, page int32, rankTime int64, dataList []*shenmo.ShenMoRankData) *uipb.SCShenMoRankGet {
	scShenMoRankGet := &uipb.SCShenMoRankGet{}
	scShenMoRankGet.IsThis = &isThis
	scShenMoRankGet.Page = &page
	scShenMoRankGet.RankTime = &rankTime

	for _, data := range dataList {
		scShenMoRankGet.RankList = append(scShenMoRankGet.RankList, buildShenMoRank(data))
	}
	return scShenMoRankGet
}

func BuildSCPlayerGongXunChanged(gongXunNum int32) *uipb.SCPlayerGongXunChanged {
	scPlayerGongXunChanged := &uipb.SCPlayerGongXunChanged{}
	scPlayerGongXunChanged.GongXunNum = &gongXunNum
	return scPlayerGongXunChanged
}

func buildShenMoRank(data *shenmo.ShenMoRankData) *uipb.ShenMoRank {
	shenMoRank := &uipb.ShenMoRank{}
	serverId := data.GetServerId()
	allianceId := data.GetAllianceId()
	allianceName := data.GetAllianceName()
	jiFenNum := data.GetJiFenNum()
	shenMoRank.ServerId = &serverId
	shenMoRank.AllianceId = &allianceId
	shenMoRank.AllianceName = &allianceName
	shenMoRank.JiFenNum = &jiFenNum
	return shenMoRank
}

func BuildSCShenMoLineUp(beforeNum int32) *uipb.SCShenMoLineUp {
	scShenMoLineUp := &uipb.SCShenMoLineUp{}
	scShenMoLineUp.BeforeNum = &beforeNum
	return scShenMoLineUp
}

func BuildSCShenMoCancleLineUp() *uipb.SCShenMoCancleLineUp {
	scShenMoCancleLineUp := &uipb.SCShenMoCancleLineUp{}
	return scShenMoCancleLineUp
}

func BuildSCShenMoLineUpSuccess() *uipb.SCShenMoLineUpSuccess {
	scShenMoLineUpSuccess := &uipb.SCShenMoLineUpSuccess{}
	return scShenMoLineUpSuccess
}

func BuildSCShenMoFinishToLineUp() *uipb.SCShenMoFinishToLineUp {
	scShenMoFinishToLineUp := &uipb.SCShenMoFinishToLineUp{}
	return scShenMoFinishToLineUp
}

func BuildSCShenMoGongXunNumChanged(gongXunNum int32) *uipb.SCShenMoSceneDataChanged {
	scShenMoSceneDataChanged := &uipb.SCShenMoSceneDataChanged{}
	scShenMoSceneDataChanged.GongXun = &gongXunNum
	return scShenMoSceneDataChanged
}

func BuildSCShenMoKillNumChanged(killNum int32) *uipb.SCShenMoSceneDataChanged {
	scShenMoSceneDataChanged := &uipb.SCShenMoSceneDataChanged{}
	scShenMoSceneDataChanged.KillNum = &killNum
	return scShenMoSceneDataChanged
}

func BuildSCShenMoSceneEnd() *uipb.SCShenMoSceneEnd {
	scShenMoSceneEnd := &uipb.SCShenMoSceneEnd{}
	return scShenMoSceneEnd
}

func BuildSCShenMoBioBroadcast(npc scene.NPC) *uipb.SCShenMoBioBroadcast {
	shenMoBioBroadcast := &uipb.SCShenMoBioBroadcast{}
	shenMoBioBroadcast.Bio = scenepbutil.BuildGeneralCollectInfo(npc)
	return shenMoBioBroadcast
}

func BuildSCShenMoSceneInfo(gongXunNum int32, killNum int32, jiFenNum int32, daQiList []scene.NPC) *uipb.SCShenMoSceneInfo {
	scShenMoGet := &uipb.SCShenMoSceneInfo{}
	scShenMoGet.JiFenNum = &jiFenNum
	scShenMoGet.GongXun = &gongXunNum
	scShenMoGet.KillNum = &killNum
	scShenMoGet.DaQiList = scenepbutil.BuildGeneralCollectInfoListByList(daQiList)
	return scShenMoGet
}

func BuildSCShenMoJiFenNumChanged(jiFenNum int32) *uipb.SCShenMoSceneDataChanged {
	scShenMoSceneDataChanged := &uipb.SCShenMoSceneDataChanged{}
	scShenMoSceneDataChanged.JiFenNum = &jiFenNum
	return scShenMoSceneDataChanged
}
