package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func BuildSCShengTanSceneInfo(p scene.Player, rankMap map[scenetypes.SceneRankType]*scene.SceneRank, bossHp, bossMaxHp int64, jiuNiangNum int32, jiuNiangPercent int32) *uipb.SCShengTanSceneInfo {
	scShengTanSceneInfo := &uipb.SCShengTanSceneInfo{}
	scShengTanSceneInfo.BossHp = &bossHp
	scShengTanSceneInfo.BossMaxHp = &bossMaxHp
	scShengTanSceneInfo.JiuNiangNum = &jiuNiangNum
	scShengTanSceneInfo.JiuNiangExpPercent = &jiuNiangPercent
	for _, r := range rankMap {
		scShengTanSceneInfo.RankInfoList = append(scShengTanSceneInfo.RankInfoList, scenepbutil.BuildSceneRankInfo(p, r))
	}

	return scShengTanSceneInfo
}

func BuildSCShengTanSceneBossHpChanged(bossHp int64) *uipb.SCShengTanSceneBossHpChanged {
	scShengTanSceneBossHpChanged := &uipb.SCShengTanSceneBossHpChanged{}
	scShengTanSceneBossHpChanged.BossHp = &bossHp
	return scShengTanSceneBossHpChanged
}

func BuildSCShengTanSceneJiuNiangChanged(jiuNiang int32, jiuNiangPercent int32) *uipb.SCShengTanSceneJiuNiangChanged {
	scShengTanSceneJiuNiangChanged := &uipb.SCShengTanSceneJiuNiangChanged{}
	scShengTanSceneJiuNiangChanged.JiuNiangNum = &jiuNiang
	scShengTanSceneJiuNiangChanged.JiuNiangExpPercent = &jiuNiangPercent
	return scShengTanSceneJiuNiangChanged
}

var (
	scShengTanSceneEnd = &uipb.SCShengTanSceneEnd{}
)

func BuildSCShengTanSceneEnd() *uipb.SCShengTanSceneEnd {

	return scShengTanSceneEnd
}
