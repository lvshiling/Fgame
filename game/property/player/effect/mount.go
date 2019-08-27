package effect

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/mount/mount"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	skilltypes "fgame/fgame/game/skill/types"
	playerxianzuncard "fgame/fgame/game/xianzuncard/player"
	xianzuncardtemplate "fgame/fgame/game/xianzuncard/template"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeMount, MountPropertyEffect)
}

//坐骑作用器
func MountPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeMount) {
		return
	}
	xianZunManager := p.GetPlayerDataManager(types.PlayerXianZunCardManagerType).(*playerxianzuncard.PlayerXianZunCardDataManager)
	mountManager := p.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	advancedId := mountInfo.AdvanceId
	mountTemplate := mount.GetMountService().GetMountNumber(int32(advancedId))
	if mountTemplate == nil {
		return
	}

	hp := int64(0)
	attack := int64(0)
	defence := int64(0)
	//培养丹食丹等级
	culLevel := mountInfo.CulLevel
	caoLiaoTemplate := mount.GetMountService().GetMountCaoLiaoTemplate(culLevel)
	if caoLiaoTemplate != nil {
		hp += int64(caoLiaoTemplate.Hp)
		attack += int64(caoLiaoTemplate.Attack)
		defence += int64(caoLiaoTemplate.Defence)
	}
	//幻化丹食丹等级
	unrealLevel := mountInfo.UnrealLevel
	huanHuaTemplate := mount.GetMountService().GetMountHuanHuaTemplate(unrealLevel)
	if huanHuaTemplate != nil {
		hp += int64(huanHuaTemplate.Hp) + prop.GetGlobal(propertytypes.BattlePropertyTypeMaxHP)
		prop.SetGlobal(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack += int64(huanHuaTemplate.Attack) + prop.GetGlobal(propertytypes.BattlePropertyTypeAttack)
		prop.SetGlobal(propertytypes.BattlePropertyTypeAttack, attack)
		defence += int64(huanHuaTemplate.Defence) + prop.GetGlobal(propertytypes.BattlePropertyTypeDefend)
		prop.SetGlobal(propertytypes.BattlePropertyTypeDefend, defence)
	}

	// 计算加成的万分比
	xianZunMap := xianZunManager.GetXianZunCardObjectMap()
	for typ, obj := range xianZunMap {
		if !obj.IsActivite() {
			continue
		}

		xianZunTemp := xianzuncardtemplate.GetXianZunCardTemplateService().GetXianZunCardTemplate(typ)
		if xianZunTemp == nil {
			continue
		}

		hp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP) + int64(xianZunTemp.MountAttrAddPercent)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, hp)
		attack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack) + int64(xianZunTemp.MountAttrAddPercent)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, attack)
		defence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend) + int64(xianZunTemp.MountAttrAddPercent)
		prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, defence)
	}

	//坐骑属性
	if mountTemplate.GetBattleAttrTemplate() != nil {
		for typ, val := range mountTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
			//坐骑不算移动速度
			if typ == propertytypes.BattlePropertyTypeMoveSpeed {
				continue
			}
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

	//非进阶坐骑
	mountOtherMap := mountManager.GetMountOtherMap()
	for _, mountTypeOtherMap := range mountOtherMap {
		for mountOtherId, mo := range mountTypeOtherMap {
			mountTemplate := mount.GetMountService().GetMount(int(mountOtherId))
			//非进阶坐骑属性
			if mountTemplate.GetBattleAttrTemplate() == nil {
				continue
			}

			skinHp := int64(0)
			skinAttack := int64(0)
			skinDefence := int64(0)
			if mountTemplate.MountUpstarBeginId != 0 && mo.Level != 0 {
				mountUpstarTemplate := mountTemplate.GetMountUpstarByLevel(mo.Level)
				skinHp = int64(mountUpstarTemplate.Hp)
				skinAttack = int64(mountUpstarTemplate.Attack)
				skinDefence = int64(mountUpstarTemplate.Defence)

				//坐骑基础全属性万分比
				if mountUpstarTemplate.MountPercent != 0 {
					oldHp := prop.GetBasePercent(propertytypes.BattlePropertyTypeMaxHP)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeMaxHP, int64(mountUpstarTemplate.MountPercent)+oldHp)
					oldAttack := prop.GetBasePercent(propertytypes.BattlePropertyTypeAttack)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeAttack, int64(mountUpstarTemplate.MountPercent)+oldAttack)
					oldDefence := prop.GetBasePercent(propertytypes.BattlePropertyTypeDefend)
					prop.SetBasePercent(propertytypes.BattlePropertyTypeDefend, int64(mountUpstarTemplate.MountPercent)+oldDefence)
				}
			}

			for typ, val := range mountTemplate.GetBattleAttrTemplate().GetAllBattleProperty() {
				//坐骑没骑的时候 没有移动速度
				if typ == propertytypes.BattlePropertyTypeMoveSpeed {
					continue
				}
				switch typ {
				case propertytypes.BattlePropertyTypeMaxHP:
					{
						val += skinHp
						break
					}
				case propertytypes.BattlePropertyTypeAttack:
					{
						val += skinAttack
						break
					}
				case propertytypes.BattlePropertyTypeDefend:
					{
						val += skinDefence
						break
					}
				}
				total := prop.GetGlobal(typ)
				total += val
				prop.SetGlobal(typ, total)
			}
		}
	}
	//附加系统属性
	additionSysPropertyEffect(p, additionsystypes.AdditionSysTypeMountEquip, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeRole, prop)
	skillByModulePropertyEffect(p, skilltypes.SkillFirstTypeMountEquip, prop)
}
