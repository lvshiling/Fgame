package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	playershangguzhiling "fgame/fgame/game/shangguzhiling/player"
	shangguzhilingtemplate "fgame/fgame/game/shangguzhiling/template"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeShangGuZhiLing, ShangGuZhiLingEffect)
}

func ShangGuZhiLingEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeShangGuZhiLing) {
		return
	}
	lingShouManager := p.GetPlayerDataManager(playertypes.PlayerShangguzhilingDataManagerType).(*playershangguzhiling.PlayerShangguzhilingDataManager)
	lingshouTempService := shangguzhilingtemplate.GetShangguzhilingTemplateService()
	lingShouBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
	//未创建的不计算
	for _, obj := range lingShouManager.GetCurrentLingShouObjectList() {
		chushiBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
		uplevelBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
		lingWenBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
		upRankBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
		upRankBasePercentProperty := make(map[propertytypes.BattlePropertyType]int32)
		lingLianBaseProperty := make(map[propertytypes.BattlePropertyType]int64)
		lingLianBasePercentProperty := make(map[propertytypes.BattlePropertyType]int32)

		//灵兽基础属性
		baseTemp := lingshouTempService.GetLingShouTemplate(obj.GetLingShouType())
		if baseTemp == nil {
			//基础模板不存在就不需要进行其他的了
			continue
		}
		chushiBaseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(baseTemp.Hp)
		chushiBaseProperty[propertytypes.BattlePropertyTypeAttack] += int64(baseTemp.Attack)
		chushiBaseProperty[propertytypes.BattlePropertyTypeDefend] += int64(baseTemp.Defence)

		//升级属性
		uplevelTemp := baseTemp.GetLevelTemp(obj.GetLevel())
		if uplevelTemp != nil {
			uplevelBaseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(uplevelTemp.Hp)
			uplevelBaseProperty[propertytypes.BattlePropertyTypeAttack] += int64(uplevelTemp.Attack)
			uplevelBaseProperty[propertytypes.BattlePropertyTypeDefend] += int64(uplevelTemp.Defence)
		}

		//灵纹属性
		for lingwenType := shangguzhilingtypes.MinLingwenType; lingwenType <= shangguzhilingtypes.MaxLingwenType; lingwenType++ {
			lingwenTemp := lingshouTempService.GetLingWenTemplate(obj.GetLingShouType(), lingwenType)
			if lingwenTemp == nil {
				continue
			}
			info := obj.GetLingWenInfo(lingwenType)
			if info == nil {
				continue
			}
			levelTemp := lingwenTemp.GetLevelTemp(info.Level)
			if levelTemp == nil {
				return
			}
			lingWenBaseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(levelTemp.Hp)
			lingWenBaseProperty[propertytypes.BattlePropertyTypeAttack] += int64(levelTemp.Attack)
			lingWenBaseProperty[propertytypes.BattlePropertyTypeDefend] += int64(levelTemp.Defence)
		}

		//进阶属性
		uprankTemp := baseTemp.GetRankTemp(obj.GetUprankLevel())
		if uprankTemp != nil {
			upRankBaseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(uprankTemp.Hp)
			upRankBaseProperty[propertytypes.BattlePropertyTypeAttack] += int64(uprankTemp.Attack)
			upRankBaseProperty[propertytypes.BattlePropertyTypeDefend] += int64(uprankTemp.Defence)
			//这个万分比是给灵纹属性加成的
			upRankBasePercentProperty[propertytypes.BattlePropertyTypeMaxHP] += uprankTemp.Percent
			upRankBasePercentProperty[propertytypes.BattlePropertyTypeAttack] += uprankTemp.Percent
			upRankBasePercentProperty[propertytypes.BattlePropertyTypeDefend] += uprankTemp.Percent
		}

		//灵炼属性--基本随机池属性
		totalStar := int32(0)
		for pos := shangguzhilingtypes.MinLinglianPosType; pos <= shangguzhilingtypes.MaxLinglianPosType; pos++ {
			poolTemp := obj.GetLingLianPoolTemplate(pos)
			if poolTemp == nil {
				continue
			}
			lingLianBaseProperty[propertytypes.BattlePropertyTypeMaxHP] += int64(poolTemp.Hp)
			lingLianBaseProperty[propertytypes.BattlePropertyTypeAttack] += int64(poolTemp.Attack)
			lingLianBaseProperty[propertytypes.BattlePropertyTypeDefend] += int64(poolTemp.Defence)
			totalStar += poolTemp.Star
		}
		//灵炼属性--星级套装属性
		taozhuangList := lingshouTempService.GetLingLianTaoZhuangListTempLate()
		for i := 0; i < len(taozhuangList); i++ {
			taozhuangTemp := taozhuangList[i]
			//降序序列第一个符合条件的
			if totalStar >= taozhuangTemp.NeedLevel {
				//这个是给初始和升级属性加成的
				lingLianBasePercentProperty[propertytypes.BattlePropertyTypeMaxHP] += taozhuangTemp.Percent
				lingLianBasePercentProperty[propertytypes.BattlePropertyTypeAttack] += taozhuangTemp.Percent
				lingLianBasePercentProperty[propertytypes.BattlePropertyTypeDefend] += taozhuangTemp.Percent
				break
			}
		}

		//----------------------------------分割线-----------------------------------

		totalHp := int64(0)
		totalAttack := int64(0)
		totalDefence := int64(0)
		var val int64
		var percent int64

		//初始和升级
		val = chushiBaseProperty[propertytypes.BattlePropertyTypeMaxHP] + uplevelBaseProperty[propertytypes.BattlePropertyTypeMaxHP]
		percent = int64(lingLianBasePercentProperty[propertytypes.BattlePropertyTypeMaxHP])
		totalHp += int64(math.Ceil(float64(val) * float64(int64(common.MAX_RATE)+percent) / float64(common.MAX_RATE)))

		val = chushiBaseProperty[propertytypes.BattlePropertyTypeAttack] + uplevelBaseProperty[propertytypes.BattlePropertyTypeAttack]
		percent = int64(lingLianBasePercentProperty[propertytypes.BattlePropertyTypeAttack])
		totalAttack += int64(math.Ceil(float64(val) * float64(int64(common.MAX_RATE)+percent) / float64(common.MAX_RATE)))

		val = chushiBaseProperty[propertytypes.BattlePropertyTypeDefend] + uplevelBaseProperty[propertytypes.BattlePropertyTypeDefend]
		percent = int64(lingLianBasePercentProperty[propertytypes.BattlePropertyTypeDefend])
		totalDefence += int64(math.Ceil(float64(val) * float64(int64(common.MAX_RATE)+percent) / float64(common.MAX_RATE)))

		//灵纹
		val = lingWenBaseProperty[propertytypes.BattlePropertyTypeMaxHP]
		percent = int64(upRankBasePercentProperty[propertytypes.BattlePropertyTypeMaxHP])
		totalHp += int64(math.Ceil(float64(val) * float64(int64(common.MAX_RATE)+percent) / float64(common.MAX_RATE)))

		val = lingWenBaseProperty[propertytypes.BattlePropertyTypeAttack]
		percent = int64(upRankBasePercentProperty[propertytypes.BattlePropertyTypeAttack])
		totalAttack += int64(math.Ceil(float64(val) * float64(int64(common.MAX_RATE)+percent) / float64(common.MAX_RATE)))

		val = lingWenBaseProperty[propertytypes.BattlePropertyTypeDefend]
		percent = int64(upRankBasePercentProperty[propertytypes.BattlePropertyTypeDefend])
		totalDefence += int64(math.Ceil(float64(val) * float64(int64(common.MAX_RATE)+percent) / float64(common.MAX_RATE)))

		//进阶
		totalHp += upRankBaseProperty[propertytypes.BattlePropertyTypeMaxHP]
		totalAttack += upRankBaseProperty[propertytypes.BattlePropertyTypeAttack]
		totalDefence += upRankBaseProperty[propertytypes.BattlePropertyTypeDefend]

		//灵炼
		totalHp += lingLianBaseProperty[propertytypes.BattlePropertyTypeMaxHP]
		totalAttack += lingLianBaseProperty[propertytypes.BattlePropertyTypeAttack]
		totalDefence += lingLianBaseProperty[propertytypes.BattlePropertyTypeDefend]

		//--------------------------------------------分割线-----------------------------------------
		lingShouBaseProperty[propertytypes.BattlePropertyTypeMaxHP] += totalHp
		lingShouBaseProperty[propertytypes.BattlePropertyTypeAttack] += totalAttack
		lingShouBaseProperty[propertytypes.BattlePropertyTypeDefend] += totalDefence
	}

	for typ, val := range lingShouBaseProperty {
		newVal := prop.GetGlobal(typ) + val
		prop.SetGlobal(typ, newVal)
	}
	return
}
