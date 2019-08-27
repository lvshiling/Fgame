package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitemplate "fgame/fgame/game/shenqi/template"
	shenqitypes "fgame/fgame/game/shenqi/types"
	skilltypes "fgame/fgame/game/skill/types"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShenQi, shenQiPropertyEffect)
}

//神器作用器
func shenQiPropertyEffect(pl player.Player, prop *propertycommon.SystemPropertySegment) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQi) {
		return
	}
	shenQiManager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	for typ := shenqitypes.MinShenQiType; typ <= shenqitypes.MaxShenQiType; typ++ {
		//碎片等级
		debeisHp, debeisAttack, debeisDefence := debeisPropertyEffect(typ, shenQiManager)
		//器灵
		qiLingHp, qiLingAttack, qiLingDefence := qiLingPropertyEffect(typ, shenQiManager)
		//汇总单个神器
		totalHp := debeisHp + qiLingHp + prop.GetBase(propertytypes.BattlePropertyTypeMaxHP)
		totalAttack := qiLingAttack + debeisAttack + prop.GetBase(propertytypes.BattlePropertyTypeAttack)
		totalDefence := qiLingDefence + debeisDefence + prop.GetBase(propertytypes.BattlePropertyTypeDefend)
		prop.SetBase(propertytypes.BattlePropertyTypeMaxHP, totalHp)
		prop.SetBase(propertytypes.BattlePropertyTypeAttack, totalAttack)
		prop.SetBase(propertytypes.BattlePropertyTypeDefend, totalDefence)
	}

	skillByModulePropertyEffect(pl, skilltypes.SkillFirstTypeShenQi, prop)
}

//碎片+淬炼
func debeisPropertyEffect(typ shenqitypes.ShenQiType, shenQiManager *playershenqi.PlayerShenQiDataManager) (hp, attack, defence int64) {
	debeisMap := shenQiManager.GetShenQiDebrisMapByShenQi(typ)
	for _, obj := range debeisMap {
		temp := shenqitemplate.GetShenQiTemplateService().GetShenQiDebrisUpByArg(obj.ShenQiType, obj.SlotId, obj.Level)
		if temp == nil {
			continue
		}
		hp += int64(temp.Hp)
		attack += int64(temp.Attack)
		defence += int64(temp.Defence)
	}

	//淬炼等级
	smeltMinLevel := shenQiManager.GetShenQiSmeltMinLevelByShenQi(typ)
	temp := shenqitemplate.GetShenQiTemplateService().GetShenQiSmeltByArg(typ, smeltMinLevel)
	if temp != nil {
		hp += int64(math.Ceil(float64(hp) * float64(temp.Percent) / float64(common.MAX_RATE)))
		attack += int64(math.Ceil(float64(attack) * float64(temp.Percent) / float64(common.MAX_RATE)))
		defence += int64(math.Ceil(float64(defence) * float64(temp.Percent) / float64(common.MAX_RATE)))
	}
	//元素之魂等级
	smeltMap := shenQiManager.GetShenQiSmeltMapByShenQi(typ)
	for _, obj := range smeltMap {
		upTemp := shenqitemplate.GetShenQiTemplateService().GetShenQiSmeltUpByArg(obj.ShenQiType, obj.SlotId, obj.Level)
		if upTemp == nil {
			continue
		}
		hp += int64(upTemp.Hp)
		attack += int64(upTemp.Attack)
		defence += int64(upTemp.Defence)
	}
	return
}

//器灵+注灵
func qiLingPropertyEffect(typ shenqitypes.ShenQiType, shenQiManager *playershenqi.PlayerShenQiDataManager) (hp, attack, defence int64) {
	qiLingMap := shenQiManager.GetShenQiQiLingMapByShenQi(typ)
	for _, objM := range qiLingMap {
		for _, obj := range objM {
			if obj.IsEmpty() {
				continue
			}
			itemId := int(obj.ItemId)
			qiLingTemp := item.GetItemService().GetItem(itemId).GetShenQiQiLingTemplate()
			if qiLingTemp == nil {
				continue
			}
			qiLingHp := int64(qiLingTemp.Hp)
			qiLingAttack := int64(qiLingTemp.Attack)
			qiLingDefence := int64(qiLingTemp.Defence)

			temp := shenqitemplate.GetShenQiTemplateService().GetShenQiZhuLingByArg(obj.ShenQiType, obj.QiLingType, obj.SlotId, obj.Level)
			if temp != nil {
				percentHp := int64(math.Ceil(float64(qiLingHp) * float64(temp.Percent) / float64(common.MAX_RATE)))
				percentAttack := int64(math.Ceil(float64(qiLingAttack) * float64(temp.Percent) / float64(common.MAX_RATE)))
				percentDefence := int64(math.Ceil(float64(qiLingDefence) * float64(temp.Percent) / float64(common.MAX_RATE)))

				qiLingHp += percentHp + int64(temp.Hp)
				qiLingAttack += percentAttack + int64(temp.Attack)
				qiLingDefence += percentDefence + int64(temp.Defence)
			}

			hp += qiLingHp
			attack += qiLingAttack
			defence += qiLingDefence
		}
	}
	//套装
	taoZhuangId := shenQiManager.GetShenQiQiLingTaoZhuangIdByShenQi(typ)
	taoZhuangTemp := shenqitemplate.GetShenQiTemplateService().GetShenQiTaoZhuangById(taoZhuangId)
	if taoZhuangTemp == nil {
		return
	}
	hp += int64(taoZhuangTemp.Hp)
	attack += int64(taoZhuangTemp.Attack)
	defence += int64(taoZhuangTemp.Defence)
	return
}
