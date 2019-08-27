package effect

import (
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playerring "fgame/fgame/game/ring/player"
	ringtemplate "fgame/fgame/game/ring/template"
	ringtypes "fgame/fgame/game/ring/types"
	skilltypes "fgame/fgame/game/skill/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeRing, RingPropertyEffect)
}

// 特戒作用器
func RingPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeRing) {
		return
	}

	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringObjMap := ringManager.GetPlayerRingObjectMap()

	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	hpPercent := int64(0)
	attackPercent := int64(0)
	defencePercent := int64(0)

	for _, ringObj := range ringObjMap {
		itemId := ringObj.GetItemId()
		data := ringObj.GetPropertyData()
		ringData := data.(*ringtypes.RingPropertyData)
		// 计算升阶增加属性
		advanceTemp := ringtemplate.GetRingTemplateService().GetRingAdvanceTemplate(itemId, ringData.Advance)
		if advanceTemp != nil {
			hp += advanceTemp.Hp
			attack += advanceTemp.Attack
			defence += advanceTemp.Defence

			// 计算加成的基础万分比
			hpPercent += int64(advanceTemp.HpPercent)
			attackPercent += int64(advanceTemp.AttackPercent)
			defencePercent += int64(advanceTemp.DefPercent)
		}

		// 计算强化增加属性
		strengthenTemp := ringtemplate.GetRingTemplateService().GetRingStrengthenTemplate(itemId, ringData.StrengthLevel)
		if strengthenTemp != nil {
			hp += strengthenTemp.Hp
			attack += strengthenTemp.Attack
			defence += strengthenTemp.Defence
		}

		// 计算净灵增加属性
		jingLingTemp := ringtemplate.GetRingTemplateService().GetRingJingLingTemplate(itemId, ringData.JingLingLevel)
		if jingLingTemp != nil {
			hp += jingLingTemp.Hp
			attack += jingLingTemp.Attack
			defence += jingLingTemp.Defence
		}

		// 计算融合增加属性
		ringTemp := ringtemplate.GetRingTemplateService().GetRingTemplate(itemId)
		if ringTemp != nil {
			hp += ringTemp.Hp
			attack += ringTemp.Attack
			defence += ringTemp.Defence
		}
	}

	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, hp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, attack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, defence)

	oldHp := prop.GetGlobalPercent(propertytypes.BattlePropertyTypeMaxHP)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeMaxHP, hpPercent+oldHp)
	oldAttack := prop.GetGlobalPercent(propertytypes.BattlePropertyTypeAttack)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeAttack, attackPercent+oldAttack)
	oldDefence := prop.GetGlobalPercent(propertytypes.BattlePropertyTypeDefend)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeDefend, defencePercent+oldDefence)

	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeRing, prop)
}
