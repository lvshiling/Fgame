package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	longgongtypes "fgame/fgame/game/longgong/types"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func BuildSCLonggongGet(p scene.Player, rankMap map[scenetypes.SceneRankType]*scene.SceneRank, bossHp int64, bossStatus longgongtypes.HeiLongStatusType, collectCount int32, pearlCount int32, cn scene.NPC) *uipb.SCLonggongGet {
	scMsg := &uipb.SCLonggongGet{}
	bossStatusInt := int32(bossStatus)

	scMsg.BossHp = &bossHp
	scMsg.CollectCount = &collectCount
	scMsg.PearlCount = &pearlCount
	scMsg.BossStatus = &bossStatusInt
	for _, r := range rankMap {
		scMsg.RankInfoList = append(scMsg.RankInfoList, scenepbutil.BuildSceneRankInfo(p, r))
	}
	if cn != nil {
		cnUid := cn.GetId()
		scMsg.Uid = &cnUid
	}

	return scMsg
}

func BuildSCLonggongPlayerValChange(collectCount int32) *uipb.SCLonggongPlayerValChange {
	scMsg := &uipb.SCLonggongPlayerValChange{}
	scMsg.CollectCount = &collectCount
	return scMsg
}

func BuildSCLonggongSceneValBroadcast(pearlCount int32, bossStatus longgongtypes.HeiLongStatusType, bossHp int64) *uipb.SCLonggongSceneValBroadcast {
	scMsg := &uipb.SCLonggongSceneValBroadcast{}
	bossStatusInt := int32(bossStatus)

	scMsg.PearlCount = &pearlCount
	scMsg.BossStatus = &bossStatusInt
	scMsg.BossHp = &bossHp
	return scMsg
}

func BuildSCLonggongScenePearlCountBroadcast(pearlCount int32) *uipb.SCLonggongSceneValBroadcast {
	scMsg := &uipb.SCLonggongSceneValBroadcast{}
	scMsg.PearlCount = &pearlCount
	return scMsg
}

func BuildSCLonggongSceneBossStatusBroadcast(bossStatus longgongtypes.HeiLongStatusType) *uipb.SCLonggongSceneValBroadcast {
	scMsg := &uipb.SCLonggongSceneValBroadcast{}
	bossStatusInt := int32(bossStatus)
	scMsg.BossStatus = &bossStatusInt
	return scMsg
}

func BuildSCLonggongSceneBossDieBroadcast(bossStatus longgongtypes.HeiLongStatusType, cnUid int64) *uipb.SCLonggongSceneValBroadcast {
	scMsg := &uipb.SCLonggongSceneValBroadcast{}
	bossStatusInt := int32(bossStatus)
	scMsg.BossStatus = &bossStatusInt
	scMsg.Uid = &cnUid
	return scMsg
}

func BuildSCLonggongSceneBossHpBroadcast(bossHp int64) *uipb.SCLonggongSceneValBroadcast {
	scMsg := &uipb.SCLonggongSceneValBroadcast{}
	scMsg.BossHp = &bossHp
	return scMsg
}

func BuildSCLonggongResult() *uipb.SCLonggongResult {
	scMsg := &uipb.SCLonggongResult{}
	return scMsg
}
