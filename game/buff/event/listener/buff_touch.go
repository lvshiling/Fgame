package listener

import (
	scenepb "fgame/fgame/common/codec/pb/scene"
	"fgame/fgame/core/event"
	"fgame/fgame/game/buff/buff"
	buffcommon "fgame/fgame/game/buff/common"
	buffeventtypes "fgame/fgame/game/buff/event/types"
	bufflogic "fgame/fgame/game/buff/logic"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	propertytypes "fgame/fgame/game/property/types"
	scenelogic "fgame/fgame/game/scene/logic"
	scenepbutil "fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
	"math"
)

//buff触发
func buffTouch(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	buffObject := data.(buffcommon.BuffObject)

	buffId := buffObject.GetBuffId()
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	hp := bo.GetHP()
	maxHp := bo.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	// maxTp := bo.GetBattleProperty(propertytypes.BattlePropertyTypeMaxTP)
	lifeAdd := int64(buffTemplate.LifeAdd) + int64(math.Ceil(float64(buffTemplate.LifePercentMax)/float64(common.MAX_RATE)*float64(maxHp))) + int64(math.Ceil(float64(buffTemplate.LifePercent)/float64(common.MAX_RATE)*float64(hp)))

	if !bo.IsDead() {
		if buffTemplate.TypeLimit&scenetypes.BattleLimitTypeMount.Mask() != 0 {
			p, ok := bo.(scene.Player)
			if ok {
				if !p.IsMountHidden() {
					p.MountHidden(true)
				}
			}
		}
		//释放子技能
		if buffTemplate.GetSkillTemplate() != nil {
			if !bo.IsSkillInCd(buffTemplate.GetSkillTemplate().TypeId) {
				scenelogic.Attack(bo, bo.GetPosition(), bo.GetAngle(), buffTemplate.GetSkillTemplate(), false)
			}
		}
		// tpAdd := int64(buffTemplate.TpAdd) + int64(math.Ceil(float64(buffTemplate.LifePercent)/float64(common.MAX_RATE)*float64(maxTp)))
		if buffTemplate.StackType&int32(scenetypes.BuffStackTypeEffect) != 0 {
			lifeAdd *= int64(buffObject.GetCulTime())
		}
		//发送伤害数据
		if lifeAdd != 0 {
			flag, damageType := buffTemplate.GetBuffPiaoZi().DamageType()
			var scObjectDamage *scenepb.SCObjectDamage
			if lifeAdd > 0 {
				if !flag {
					//以防万一 策划没配
					damageType = scenetypes.DamageTypeRecovery
				}
				bo.AddHP(lifeAdd)
				scObjectDamage = scenepbutil.BuildSCObjectDamage(bo, damageType, lifeAdd, 0, buffObject.GetOwnerId())
				scenelogic.BroadcastNeighborIncludeSelf(bo, scObjectDamage)
			} else {
				if !flag {
					//以防万一 策划没配
					damageType = scenetypes.DamageTypeAttack
				}
				// bo.CostHP(-lifeAdd, buffObject.GetOwnerId())
				// scObjectDamage = pbutil.BuildSCObjectDamage(bo, damageType, -lifeAdd, 0, buffObject.GetOwnerId())
				scenelogic.CostHP(bo, -lifeAdd, 0, buffObject.GetOwnerId(), damageType)
			}

		}

	}
	//经验
	if buffTemplate.GetExp != 0 || buffTemplate.GetExpPoint != 0 {
		d := buff.CreateBuffExpEventData(int32(buffTemplate.TemplateId()), buffTemplate.GetExp, buffTemplate.GetExpPoint)
		gameevent.Emit(buffeventtypes.EventTypeBuffAddExp, bo, d)
	}
	//同步属性
	if buffTemplate.GetTouchType() == scenetypes.BuffTouchTypeImmediate {
		bufflogic.UpdateBattleProperty(bo)
	}

	//判断是否触发状态
	if buffTemplate.GetTouchType() == scenetypes.BuffTouchTypeObjectDamageSelf {
		for subBuffId, subBuffRage := range buffTemplate.GetSubBuffMap() {
			hit := mathutils.RandomHit(common.MAX_RATE, int(subBuffRage))
			if hit {
				scenelogic.AddBuff(bo, subBuffId, buffObject.GetOwnerId(), int64(subBuffRage))
			}
		}
	}

	//判断是否是变形
	if buffTemplate.ModelId != 0 {
		switch tbo := bo.(type) {
		case scene.Player:
			tbo.SetModel(buffTemplate.ModelId)
			break
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(buffeventtypes.EventTypeBuffTouch, event.EventListenerFunc(buffTouch))
}
