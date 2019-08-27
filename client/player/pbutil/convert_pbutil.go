package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	propertytypes "fgame/fgame/game/property/types"
)

func ConvertFromBattleProperty(propertyList []*uipb.Property) map[int32]int64 {
	ps := make(map[int32]int64)
	for _, property := range propertyList {
		typ := propertytypes.BattlePropertyType(property.GetKey())
		if !typ.IsValid() {
			//TODO 警告
			continue
		}
		val := property.GetValue()
		ps[property.GetKey()] = val
	}
	return ps
}

//基本属性
func ConvertFromBaseProperty(propertyList []*uipb.Property) map[int32]int64 {
	ps := make(map[int32]int64)
	for _, property := range propertyList {
		typ := propertytypes.BasePropertyType(property.GetKey())
		if !typ.IsValid() {
			//TODO 警告
			continue
		}
		val := property.GetValue()
		ps[property.GetKey()] = val
	}
	return ps
}
