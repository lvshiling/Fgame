package effect

import (
	"fgame/fgame/game/common/common"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeGoldequip, GoldequipPropertyEffect)
}

//元神金装属性
func GoldequipPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	equipBag := goldequipManager.GetGoldEquipBag()

	totalHp := int64(0)
	totalAttack := int64(0)
	totalDefence := int64(0)
	totalHpPercent := int64(0)
	totalAttackPercent := int64(0)
	totalDefencePercent := int64(0)
	totalForce := int64(0)

	mubiaoMap := make(map[goldequiptypes.TaoZhuangMuBiaoType]int32)
	mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeGodCasting] = common.MAX_LEVEL
	for _, goldequip := range equipBag.GetAll() {
		if goldequip.IsEmpty() {
			mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeGodCasting] = 0
			continue
		}
		//装备属性
		itemId := int(goldequip.GetItemId())
		level := goldequip.GetLevel()
		equipBasicAttrPercent := float64(0)
		propertyData := goldequip.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)
		newStLevel := goldequip.GetNewStLevel()
		posType := goldequip.GetSlotId()
		goldEquipTemplate := item.GetItemService().GetItem(itemId).GetGoldEquipTemplate()

		mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeUpstar] += newStLevel
		mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeLevel] += level
		mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeOpen] += propertyData.OpenLightLevel
		godCastingEquipLevel := goldEquipTemplate.GetGodCastingEquipLevel()
		if godCastingEquipLevel < mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeGodCasting] {
			mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeGodCasting] = godCastingEquipLevel
		}

		//开光基础属性百分比
		if goldEquipTemplate.GoldeuipOpenlightId != 0 {
			openLightTemp := goldEquipTemplate.GetOpenLightTemplate(propertyData.OpenLightLevel)
			equipBasicAttrPercent = float64(openLightTemp.AttrPercent)
		}
		//附加属性
		for attrIndex, attrId := range propertyData.AttrList {
			attrTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetFuJiaAttrTemplate(attrId)
			if attrTemplate == nil {
				continue
			}

			showAttrIndex := int32(attrIndex + 1)
			upstarCondition := goldEquipTemplate.GetActivateCondition(showAttrIndex)
			if newStLevel < upstarCondition {
				continue
			}

			equipBasicAttrPercent += float64(attrTemplate.EquipPercent)

			totalHp += int64(attrTemplate.Hp)
			totalAttack += int64(attrTemplate.Attack)
			totalDefence += int64(attrTemplate.Defence)
		}

		totalHp += int64(math.Ceil(float64(goldEquipTemplate.Hp) * (float64(common.MAX_RATE) + equipBasicAttrPercent) / float64(common.MAX_RATE)))
		totalAttack += int64(math.Ceil(float64(goldEquipTemplate.Attack) * (float64(common.MAX_RATE) + equipBasicAttrPercent) / float64(common.MAX_RATE)))
		totalDefence += int64(math.Ceil(float64(goldEquipTemplate.Defence) * (float64(common.MAX_RATE) + equipBasicAttrPercent) / float64(common.MAX_RATE)))
		totalHpPercent += int64(goldEquipTemplate.HpPercent)
		totalAttackPercent += int64(goldEquipTemplate.AttPercent)
		totalDefencePercent += int64(goldEquipTemplate.DefPercent)

		//强化升星
		// if goldEquipTemplate.GoldeuipUpstarId > 0 {
		// 	upstarTemplate := goldEquipTemplate.GetUpstarTemplate(propertyData.UpstarLevel)
		// 	totalHp += int64(upstarTemplate.AddHp)
		// 	totalAttack += int64(upstarTemplate.AddAttack)
		// 	totalDefence += int64(upstarTemplate.AddDef)
		// }

		//新部位强化
		strengthenBuWeiTemp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(posType, newStLevel)
		if newStLevel > 0 && strengthenBuWeiTemp != nil {
			totalHp += int64(strengthenBuWeiTemp.AddHp)
			totalAttack += int64(strengthenBuWeiTemp.AddAttack)
			totalDefence += int64(strengthenBuWeiTemp.AddDef)
		}

		//强化属性
		if goldEquipTemplate.GoldequipStrenId > 0 {
			goldEquipStrengthenTemplate := goldEquipTemplate.GetStrengthenTemplate(level)
			totalHp += int64(goldEquipStrengthenTemplate.Hp)
			totalAttack += int64(goldEquipStrengthenTemplate.Attack)
			totalDefence += int64(goldEquipStrengthenTemplate.Defence)
			totalHpPercent += int64(goldEquipStrengthenTemplate.HpPercent)
			totalAttackPercent += int64(goldEquipStrengthenTemplate.AttPercent)
			totalDefencePercent += int64(goldEquipStrengthenTemplate.DefPercent)
		}

		//宝石属性
		for _, itemId := range goldequip.GemInfo {
			//宝石属性
			gemItemTemplate := item.GetItemService().GetItem(int(itemId))
			if gemItemTemplate == nil {
				continue
			}
			//装备自身属性
			attrTemplate := gemItemTemplate.GetGemAttrTemplate()
			for typ, val := range attrTemplate.GetAllBattleProperty() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
			mubiaoMap[goldequiptypes.TaoZhuangMuBiaoTypeGem] += gemItemTemplate.TypeFlag2
		}

		//铸灵属性
		if goldEquipTemplate.IsGodCastingEquip() {
			for i := goldequiptypes.MinSpiritType; i <= goldequiptypes.MaxSpiritType; i++ {
				spiritTemp := goldequiptemplate.GetGoldEquipTemplateService().GetCastingSpiritTemplate(posType, i)
				if spiritTemp == nil {
					continue
				}
				spiritInfo := goldequip.GetCastingSpiritInfo(i)
				spiritLevelTemp := spiritTemp.GetLevelTemplate(spiritInfo.Level)
				if spiritLevelTemp == nil {
					continue
				}
				totalHp += int64(spiritLevelTemp.Hp)
				totalAttack += int64(spiritLevelTemp.Attack)
				totalDefence += int64(spiritLevelTemp.Defence)
			}
			for i := goldequiptypes.MinForgeSoulType; i <= goldequiptypes.MaxForgeSoulType; i++ {
				soulTemp := goldequiptemplate.GetGoldEquipTemplateService().GetForgeSoulTemplate(posType, i)
				if soulTemp == nil {
					continue
				}
				soulInfo := goldequip.GetForgeSoulInfo(i)
				soulLevelTemp := soulTemp.GetLevelTemplate(soulInfo.Level)
				if soulLevelTemp == nil {
					continue
				}
				totalForce += int64(soulLevelTemp.AddPower)
			}
		}
	}

	// 套装属性
	for mubiaoType, level := range mubiaoMap {
		mubiaoTaoZhuangTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetMuBiaoTaoZhuangTemplate(mubiaoType, level)
		if mubiaoTaoZhuangTemplate != nil {
			for typ, val := range mubiaoTaoZhuangTemplate.GetBattlePropertyMap() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

	curHp := prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
	curHp += totalHp
	curAttack := prop.GetBase(propertytypes.BattlePropertyTypeAttack)
	curAttack += totalAttack
	curDefend := prop.GetBase(propertytypes.BattlePropertyTypeDefend)
	curDefend += totalDefence
	curForce := prop.GetBase(propertytypes.BattlePropertyTypeForce)
	curForce += totalForce
	prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, curHp)
	prop.SetBase(propertytypes.BattlePropertyTypeAttack, curAttack)
	prop.SetBase(propertytypes.BattlePropertyTypeDefend, curDefend)
	prop.SetBase(propertytypes.BattlePropertyTypeForce, curForce)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeMaxHP, totalHpPercent)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeAttack, totalAttackPercent)
	prop.SetGlobalPercent(propertytypes.BattlePropertyTypeDefend, totalDefencePercent)

	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeGoldEquipSuit, prop)
}
