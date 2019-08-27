package listener

import (
	"fgame/fgame/core/event"
	bufflogic "fgame/fgame/game/buff/logic"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
)

//移动
func battleObjectAttacked(target event.EventTarget, data event.EventData) (err error) {
	// attackedEventData := e.EventData().(*sceneeventtypes.BattleObjectAttackedData)
	bo := target.(scene.BattleObject)
	attackObject := data.(scene.BattleObject)
	//获得伤害触发
	bufflogic.TouchBuffByAction(bo, scenetypes.BuffTouchTypeHurted)
	bufflogic.TouchBuffByAction(bo, scenetypes.BuffTouchTypeHurtedOther)
	//移除被攻击打断
	bufflogic.RemoveBuffByAction(bo, scenetypes.BuffRemoveTypeAttacked)

	// 添加buff的buff
	for _, buffObj := range bo.GetBuffs() {
		attackBuffId := buffObj.GetBuffId()
		attackBuffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(attackBuffId)
		if attackBuffTemplate.GetTouchType() != scenetypes.BuffTouchTypeHurtedOther {
			continue
		}
		//判断是否触发状态
		for subBuffId, subBuffRage := range attackBuffTemplate.GetSubBuffMap() {
			subRateHit := mathutils.RandomHit(common.MAX_RATE, int(subBuffRage))
			if subRateHit {
				attackId := int64(0)
				if attackObject.GetSceneObjectType() == scenetypes.BiologyTypePlayer {
					attackId = attackObject.GetId()
				}
				//添加buff
				scenelogic.AddBuff(attackObject, subBuffId, attackId, int64(subBuffRage))
			}
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttacked, event.EventListenerFunc(battleObjectAttacked))
}
