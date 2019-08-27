package property

import (
	"fgame/fgame/game/player"
	propertycommon "fgame/fgame/game/property/common"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fmt"
)

//属性作用器
type PropertyEffectorFunc func(pt player.Player, prop *propertycommon.SystemPropertySegment)

var (
	//角色属性作用器列表
	playerPropertyEffectorList = []playerpropertytypes.PropertyEffectorType{}
	playerPropertyEffectorMap  = make(map[playerpropertytypes.PropertyEffectorType]PropertyEffectorFunc)
)

//注册玩家属性作用器
func RegisterPlayerPropertyEffector(propertyEffectorType playerpropertytypes.PropertyEffectorType, pef PropertyEffectorFunc) {
	_, exist := playerPropertyEffectorMap[propertyEffectorType]
	if exist {
		panic(fmt.Errorf("repeate register player property effector %s", propertyEffectorType.String()))
	}
	playerPropertyEffectorList = append(playerPropertyEffectorList, propertyEffectorType)
	playerPropertyEffectorMap[propertyEffectorType] = pef
}

//获取角色属性作用器
func GetPlayerPropertyEffector(propertyEffectorType playerpropertytypes.PropertyEffectorType) PropertyEffectorFunc {
	pef, exist := playerPropertyEffectorMap[propertyEffectorType]
	if !exist {
		return nil
	}
	return pef
}

//获取角色属性作用器
func GetAllPlayerPropertyEffectors() []playerpropertytypes.PropertyEffectorType {
	return playerPropertyEffectorList
}
