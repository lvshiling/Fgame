package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droppbutil "fgame/fgame/game/drop/pbutil"
	"fgame/fgame/game/player"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func BuildSCShenYuSceneInfoNotice(pl player.Player, keyNum, round int32, rankMap map[scenetypes.SceneRankType]*scene.SceneRank) *uipb.SCShenYuSceneInfoNotice {
	scMsg := &uipb.SCShenYuSceneInfoNotice{}
	scMsg.KeyNum = &keyNum
	scMsg.Round = &round
	for _, r := range rankMap {
		scMsg.RankInfoList = append(scMsg.RankInfoList, scenepbutil.BuildSceneRankInfo(pl, r))
	}
	return scMsg
}

func BuildSCShenYuKeyNumChangeNotice(keyNum int32) *uipb.SCShenYuKeyNumChangeNotice {
	scMsg := &uipb.SCShenYuKeyNumChangeNotice{}
	scMsg.KeyNum = &keyNum
	return scMsg
}

func BuildSCShenYuFinishRew(round, ranking int32, itemMap map[int32]int32) *uipb.SCShenYuFinishRew {
	scMsg := &uipb.SCShenYuFinishRew{}
	scMsg.Round = &round
	scMsg.Ranking = &ranking
	scMsg.DropList = droppbutil.BuildSimpleDropInfoList(itemMap)
	return scMsg
}
 