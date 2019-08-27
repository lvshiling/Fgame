package effect

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playeryinglingpu "fgame/fgame/game/yinglingpu/player"
	yinglingputemplate "fgame/fgame/game/yinglingpu/template"
	"math"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeYingLingPu, YingLingPuPropertyEffect)
}

//英灵谱
func YingLingPuPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	if !p.IsFuncOpen(funcopentypes.FuncOpenTypeYingLingPu) {
		return
	}

	//套装
	yingLingPuManager := p.GetPlayerDataManager(playertypes.PlayerYingLingPuManagerType).(*playeryinglingpu.PlayerYingLingPuManager)
	ylpSuitMap := yingLingPuManager.GetYingLingPuSuitMap()
	for _, suitTemp := range ylpSuitMap {
		for typ, val := range suitTemp.GetBattlePropertyMap() {
			val += prop.GetBase(typ)
			prop.SetBase(typ, val)
		}
	}

	// 英灵普
	for _, ylp := range yingLingPuManager.GetAllYingLingPu() {

		// levelTemplate := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuLevel(ylp.TuJianId, ylp.Level, ylp.TuJianType)
		// if levelTemplate == nil {
		// 	continue
		// }

		ylpTemplate := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuById(ylp.TuJianId, ylp.TuJianType)
		if ylpTemplate == nil {
			continue
		}

		levelBattlPropertyMap := ylpTemplate.GetLevelSuan().GetLevelBattlePropertyMap(ylp.Level)

		suitTemp, isSuit := ylpSuitMap[int32(ylpTemplate.Id)]
		for typ, val := range levelBattlPropertyMap {
			ratio := int64(common.MAX_RATE)
			if isSuit {
				ratio += suitTemp.GetBattlePropertyPercentMap()[typ]
			}
			val = int64(math.Ceil(float64(val) * float64(ratio) / float64(common.MAX_RATE)))
			val += prop.GetBase(typ)
			prop.SetBase(typ, val)
		}
	}

	// 碎片
	for _, suiPian := range yingLingPuManager.GetAllYingLingPuSuiPian() {
		suiPianTemplate := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuSuiPian(suiPian.TuJianId, suiPian.SuiPianId, suiPian.TuJianType)
		if suiPianTemplate == nil {
			continue
		}

		ylpTemplate := yinglingputemplate.GetYingLingPuTemplateService().GetYingLingPuById(suiPian.TuJianId, suiPian.TuJianType)
		if ylpTemplate == nil {
			continue
		}

		suitTemp, isSuit := ylpSuitMap[int32(ylpTemplate.Id)]
		for typ, val := range suiPianTemplate.GetBattlePropertyMap() {
			ratio := int64(common.MAX_RATE)
			if isSuit {
				ratio += suitTemp.GetBattlePropertyPercentMap()[typ]
			}
			val = int64(math.Ceil(float64(val) * float64(ratio) / float64(common.MAX_RATE)))
			val += prop.GetBase(typ)
			prop.SetBase(typ, val)
		}
	}
}
