package listener

import (
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
	shenyueventtypes "fgame/fgame/game/shenyu/event/types"
	pbuitl "fgame/fgame/game/shenyu/pbutil"
	playershenyu "fgame/fgame/game/shenyu/player"
	shenyutypes "fgame/fgame/game/shenyu/types"
)

//神域玩家钥匙改变
func shenYuKeyChange(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	if pl.GetScene().MapTemplate().GetMapType() != scenetypes.SceneTypeShenYu {
		return
	}

	shenYuManager := pl.GetPlayerDataManager(types.PlayerShenYuDataManagerType).(*playershenyu.PlayerShenYuDataManager)
	keyNum := shenYuManager.GetKeyNum()
	scMsg := pbuitl.BuildSCShenYuKeyNumChangeNotice(keyNum)
	pl.SendMsg(scMsg)
	pl.SetShenYuKey(keyNum)

	//更新排行榜玩家数据
	totalValue := int64(keyNum)
	pl.UpdateActivityRankValue(activitytypes.ActivityTypeShenYu, shenyutypes.ShenYuSceneRankTypeKey, totalValue)
	return
}

func init() {
	gameevent.AddEventListener(shenyueventtypes.EventTypeShenYuKeyChange, event.EventListenerFunc(shenYuKeyChange))
}
