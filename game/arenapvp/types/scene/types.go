package scene

import (
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/scene/scene"
)

type ArenapvpSceneRankType int32

func (t ArenapvpSceneRankType) GetRankType() int32 {
	return int32(t)
}

const (
	ArenapvpSceneRankTypePoint ArenapvpSceneRankType = iota //pvp海选积分排行榜
)

func CreateArenapvpSceneRankType(rankType int32) scene.ActivityRankType {
	return ArenapvpSceneRankType(rankType)
}

func init() {
	scene.RegistActivityRankTypeFactory(activitytypes.ActivityTypeArenapvp, scene.ActivityRankTypeFactoryFunc(CreateArenapvpSceneRankType))
}
