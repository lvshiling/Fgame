package listener

import (
	"fgame/fgame/core/event"

	activitytypes "fgame/fgame/game/activity/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/shengtan/pbutil"
	shengtanscene "fgame/fgame/game/shengtan/scene"
)

//玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeAllianceShengTan {
		return
	}

	sd, ok := s.SceneDelegate().(shengtanscene.ShengTanSceneData)
	if !ok {
		return
	}
	endTime := s.GetEndTime()
	pl.EnterActivity(activitytypes.ActivityTypeAllianceShengTan, endTime)
	rankMap := s.GetAllRanks()

	curHp, maxHp := sd.GetBossHp()
	jiuNiangNum, jiuNiangPercent := sd.GetJiuNiangNum()
	scShengTanSceneInfo := pbutil.BuildSCShengTanSceneInfo(pl, rankMap, curHp, maxHp, jiuNiangNum, jiuNiangPercent)
	pl.SendMsg(scShengTanSceneInfo)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
