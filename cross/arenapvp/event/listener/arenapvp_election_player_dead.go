package listener

import (
	"fgame/fgame/core/event"
	activitytypes "fgame/fgame/game/activity/types"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	arenapvpscenetypes "fgame/fgame/game/arenapvp/types/scene"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"math"
)

// 玩家死亡
func haiXuanPlayerDead(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(scene.Player)
	if !ok {
		return
	}

	attackId, ok := data.(int64)
	if !ok {
		return
	}

	s := spl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvpHaiXuan {
		return
	}

	pvpConstantTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpConstantTemplate()
	splVal := spl.GetActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint)
	dropNum := int64(math.Ceil(float64(splVal) * float64(pvpConstantTemp.JifenBekillPercent) / float64(common.MAX_RATE)))
	spl.UpdateActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint, splVal-dropNum)

	//击杀获得积分
	attackPl := s.GetPlayer(attackId)
	if attackPl == nil {
		return
	}
	attackVal := attackPl.GetActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint)
	attackPl.UpdateActivityRankValue(activitytypes.ActivityTypeArenapvp, arenapvpscenetypes.ArenapvpSceneRankTypePoint, attackVal+dropNum)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(haiXuanPlayerDead))
}
