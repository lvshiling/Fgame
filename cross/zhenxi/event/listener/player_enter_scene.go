package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/shareboss/shareboss"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

//进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeZhenXi {
		return
	}
	bossType := worldbosstypes.BossTypeZhenXi
	reliveTime := pl.GetBossReliveTime(bossType)
	//推送场景boss信息
	bossList := shareboss.GetShareBossService().GetShareBossListGroupByMap(bossType, s.MapId())
	scMsg := pbutil.BuildSCWorldBossListInfoNotice(bossList, bossType, reliveTime)
	pl.SendMsg(scMsg)

	// zhenXiDataManager := pl.GetPlayerDataManager(playertypes.PlayerZhenXiDataManagerType).(*playerzhenxi.PlayerZhenXiDataManager)
	// reliveTime := zhenXiDataManager.GetReliveTime()
	// scPlayerZhenXiBossInfo := zhenxipbutil.BuildSCPlayerZhenXiBossInfo(pl, reliveTime)
	// pl.SendMsg(scPlayerZhenXiBossInfo)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
