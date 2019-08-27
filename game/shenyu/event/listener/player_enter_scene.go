package listener

import (
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/shenyu/pbutil"
	playershenyu "fgame/fgame/game/shenyu/player"
	shenyuscene "fgame/fgame/game/shenyu/scene"
	shenyutypes "fgame/fgame/game/shenyu/types"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeShenYu {
		return
	}

	sd, ok := s.SceneDelegate().(shenyuscene.ShenYuSceneData)
	if !ok {
		return
	}

	shenYuTemp := sd.GetShenYuTemplate()
	sRound := shenYuTemp.RoundType
	shenYuManager := pl.GetPlayerDataManager(playertypes.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)
	shenYuManager.EnterShenYu(sd.GetActivityEndTime(), sRound, shenYuTemp.IsResetkey())

	// 更新排行玩家数据
	keyNum := shenYuManager.GetKeyNum()
	totalValue := int64(keyNum)
	pl.UpdateActivityRankValue(activitytypes.ActivityTypeShenYu, shenyutypes.ShenYuSceneRankTypeKey, totalValue)

	rankMap := s.GetAllRanks()
	scMsg := pbutil.BuildSCShenYuSceneInfoNotice(pl, keyNum, sRound, rankMap)
	pl.SendMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
