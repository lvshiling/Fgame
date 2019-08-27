package effect

import (
	"fgame/fgame/game/dan/dan"
	playerdan "fgame/fgame/game/dan/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeDan, DanPropertyEffect)
}

//丹药作用器
func DanPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	danManager := p.GetPlayerDataManager(types.PlayerDanDataManagerType).(*playerdan.PlayerDanDataManager)
	danInfo := danManager.GetDanInfo()
	danTemplate := dan.GetDanService().GetEatDan(int(danInfo.LevelId))
	//食丹系统默认不开启 levelId=0
	if danTemplate == nil {
		return
	}
	//当前levelId级 食丹属性
	for itemId, num := range danInfo.DanInfoMap {
		if num <= 0 {
			continue
		}
		attrTemplate := item.GetItemService().GetItem(itemId).GetDanAttrTemplate()
		for typ, val := range attrTemplate.GetAllBattleProperty() {
			total := prop.GetBase(typ)
			total += val * int64(num)
			prop.SetBase(typ, total)
		}
	}
	//统计1-levelId级 食丹属性
	for i := 1; i < int(danInfo.LevelId); i++ {
		danTemplate := dan.GetDanService().GetEatDan(i)
		for typ, val := range danTemplate.GetBattlePropertyMap() {
			total := prop.GetBase(typ)
			total += val
			prop.SetBase(typ, total)
		}
	}

}
