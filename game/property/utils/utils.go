package proopertyutils

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	propertytypes "fgame/fgame/game/property/types"
	"math"
)

//重新计算战力
func CulculateAllForce(propertyMap map[propertytypes.BattlePropertyType]int64) int64 {
	force := float64(0)
	forceTemplate := constant.GetConstantService().GetForceTemplate()
	for k, v := range forceTemplate.GetAllForceProperty() {
		val := propertyMap[k]
		tempForce := float64(val) / common.MAX_RATE * float64(v)
		force += tempForce
	}
	power := int64(math.Floor(force)) + propertyMap[propertytypes.BattlePropertyTypeForce]
	return power
}
