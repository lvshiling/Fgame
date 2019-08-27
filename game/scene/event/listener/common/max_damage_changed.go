package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//同步玩家
func maxDamageChanged(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	s := bo.GetScene()
	if s == nil {
		return
	}
	n, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	if n.GetBiologyTemplate().GetDropJudgeType() != scenetypes.DropJudgeTypeMaxHurt && n.GetBiologyTemplate().GetDropJudgeType() != scenetypes.DropJudgeTypeMaxHurtOrTeam {
		return
	}
	buffId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropOwnerBuffId))

	eventData := data.(*battleeventtypes.BattleObjectMaxDamageChangedEventData)
	originAttackId := eventData.GetOriginAttackId()
	originScenObject := s.GetSceneObject(originAttackId)
	originObject, ok := originScenObject.(scene.BattleObject)
	if ok {
		scenelogic.RemoveBuff(originObject, buffId)
	}
	currentAttackId := eventData.GetCurrentAttackId()
	currentSceneObject := s.GetSceneObject(currentAttackId)
	curretObject, ok := currentSceneObject.(scene.BattleObject)
	if ok {
		scenelogic.AddBuff(curretObject, buffId, 0, common.MAX_RATE)
	}
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectMaxDamageChanged, event.EventListenerFunc(maxDamageChanged))
}
