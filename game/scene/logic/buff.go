package logic

import (
	buffcommon "fgame/fgame/game/buff/common"
	bufftemplate "fgame/fgame/game/buff/template"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/pkg/mathutils"
)

func AddBuff(bo scene.BattleObject, buffId int32, ownerId int64, rate int64) {
	AddBuffWithTianFuList(bo, buffId, ownerId, rate, nil)
}

func AddBuffs(bo scene.BattleObject, buffId int32, ownerId int64, times int32, rate int64) {
	AddBuffsWithTianFuList(bo, buffId, ownerId, times, rate, nil)
}

func AddBuffWithTianFuList(bo scene.BattleObject, buffId int32, ownerId int64, rate int64, tianFuList []int32) {
	AddBuffsWithTianFuList(bo, buffId, ownerId, 1, rate, tianFuList)
}

func AddBuffsWithTianFuList(bo scene.BattleObject, buffId int32, ownerId int64, times int32, rate int64, tianFuList []int32) {
	buffTemplate := bufftemplate.GetBuffTemplateService().GetBuff(buffId)
	//buff不存在
	if buffTemplate == nil {
		return
	}
	touch := true
	if len(buffTemplate.GetParentBuffList()) != 0 {
		touch = false
		//判断前置buff
		for _, parentBuffId := range buffTemplate.GetParentBuffList() {
			if bo.GetBuff(parentBuffId) != nil {
				touch = true
				break
			}
		}
	}
	if !touch {
		return
	}
	//被动光效
	if buffTemplate.GetBuffType() == scenetypes.BuffTypePassive {
		//广播
		bo.AddBuff(buffId, ownerId, times, tianFuList)
		return
	}

	//buff
	if buffTemplate.GetBuffType() == scenetypes.BuffTypeCommon || buffTemplate.GetBuffType() == scenetypes.BuffTypeTitleDingZhi {
		if bo.GetSceneObjectSetType().Immune()&buffTemplate.ImmuneType != 0 {
			//TODO: 免疫
			scObjectDamage := pbutil.BuildSCObjectDamage(bo, scenetypes.DamageTypeMianYi, 0, 0, ownerId)
			BroadcastNeighborIncludeSelf(bo, scObjectDamage)
			return
		}

		//判断抗性
		buffKangXing := buffTemplate.GetBuffKangXing()
		flag, pType := buffKangXing.PropertyType()
		if flag {
			kangXing := bo.GetBattleProperty(pType)
			rate -= kangXing
			hit := false
			if rate > 0 {
				hit = mathutils.RandomHit(common.MAX_RATE, int(rate))
			}
			if !hit {
				//TODO 免疫
				scObjectDamage := pbutil.BuildSCObjectDamage(bo, scenetypes.DamageTypeMianYi, 0, 0, ownerId)
				BroadcastNeighborIncludeSelf(bo, scObjectDamage)
				return
			}
		}
		bo.AddBuff(buffId, ownerId, times, tianFuList)
		return
	}
	return
}

func RemoveBuff(bo scene.BattleObject, buffId int32) {
	bo.RemoveBuff(buffId)
}

func UpdateBuff(bo scene.BattleObject, buffObject buffcommon.BuffObject) {
	bo.UpdateBuff(buffObject)
}
