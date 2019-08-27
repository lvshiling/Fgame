package effect

import (
	"fgame/fgame/game/dragon/dragon"
	playerdragon "fgame/fgame/game/dragon/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertyproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

func init() {
	playerpropertyproperty.RegisterPlayerPropertyEffector(playerpropertytypes.PlayerPropertyEffectorTypeDragon, DragonPropertyEffect)
}

//神龙现世作用器
func DragonPropertyEffect(p player.Player, prop *propertycommon.SystemPropertySegment) {
	manager := p.GetPlayerDataManager(types.PlayerDragonDataManagerType).(*playerdragon.PlayerDragonDataManager)
	dragonInfo := manager.GetDragon()
	dragonTemplate := dragon.GetDragonService().GetDragonTemplate(dragonInfo.StageId)
	if dragonTemplate == nil {
		return
	}
	//当前阶数属性
	for itemId, num := range dragonInfo.ItemInfoMap {
		if num <= 0 {
			continue
		}
		attrTemplate := item.GetItemService().GetItem(int(itemId)).GetDragonAttrTemplate()
		if attrTemplate != nil {
			for typ, val := range attrTemplate.GetAllBattleProperty() {
				total := prop.GetBase(typ)
				total += val * int64(num)
				prop.SetBase(typ, total)
			}
		}
	}
	//统计1-StageId阶数 属性
	for i := int32(1); i < int32(dragonInfo.StageId); i++ {
		dragonTemplate := dragon.GetDragonService().GetDragonTemplate(i)
		if dragonTemplate == nil {
			continue
		}
		if dragonTemplate != nil {
			for typ, val := range dragonTemplate.GetBattlePropertyMap() {
				total := prop.GetBase(typ)
				total += val
				prop.SetBase(typ, total)
			}
		}
	}

}
