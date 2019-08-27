package property

import (
	lingtongplayertypes "fgame/fgame/game/lingtong/player/types"
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	"fmt"
)

//属性作用器
type PropertyEffectorFunc func(pt player.Player, prop *propertycommon.SystemPropertySegment)

var (
	//角色属性作用器列表
	lingTongPropertyEffectorList = []lingtongplayertypes.PropertyEffectorType{}
	lingTongPropertyEffectorMap  = make(map[lingtongplayertypes.PropertyEffectorType]PropertyEffectorFunc)
)

//注册玩家属性作用器
func RegisterLingTongPropertyEffector(propertyEffectorType lingtongplayertypes.PropertyEffectorType, pef PropertyEffectorFunc) {
	_, exist := lingTongPropertyEffectorMap[propertyEffectorType]
	if exist {
		panic(fmt.Errorf("repeate register lingTong property effector %s", propertyEffectorType.String()))
	}
	lingTongPropertyEffectorList = append(lingTongPropertyEffectorList, propertyEffectorType)
	lingTongPropertyEffectorMap[propertyEffectorType] = pef
}

//获取角色属性作用器
func GetLingTongPropertyEffector(propertyEffectorType lingtongplayertypes.PropertyEffectorType) PropertyEffectorFunc {
	pef, exist := lingTongPropertyEffectorMap[propertyEffectorType]
	if !exist {
		return nil
	}
	return pef
}

//获取角色属性作用器
func GetAllLingTongPropertyEffectors() []lingtongplayertypes.PropertyEffectorType {
	return lingTongPropertyEffectorList
}
