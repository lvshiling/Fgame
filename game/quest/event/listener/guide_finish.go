package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	guidereplicaeventtypes "fgame/fgame/game/guidereplica/event/types"
	guidereplicatypes "fgame/fgame/game/guidereplica/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	scenetypes "fgame/fgame/game/scene/types"
)

func guideSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	success, ok := data.(bool)
	if !ok {
		return
	}

	if !success {
		return
	}

	//猫狗奖励
	catDogFinish(pl, success)
	// 救援奖励
	rescureFinish(pl, success)
	return
}

func catDogFinish(pl player.Player, success bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeGuideReplicaCatDog {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeGuideReplica, int32(guidereplicatypes.GuideReplicaTypeCatDog), 1)
	return
}

func rescureFinish(pl player.Player, success bool) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeGuideReplicaRescue {
		return
	}

	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeGuideReplica, int32(guidereplicatypes.GuideReplicaTypeRescure), 1)
	return
}

func init() {
	gameevent.AddEventListener(guidereplicaeventtypes.EventTypeGuideFinish, event.EventListenerFunc(guideSceneFinish))
}
