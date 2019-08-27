package player

// import (
// 	propertycommon "fgame/fgame/game/property/common"
// 	propertynpctypes "fgame/fgame/game/property/npc/types"
// 	"fgame/fgame/game/scene/scene"
// )

// //属性作用器
// type PropertyEffectorFunc func(pt scene.NPC, prop *propertycommon.BattlePropertySegment)

// var (
// 	//npc属性作用器列表
// 	npcPropertyEffectorList = []propertynpctypes.NPCPropertyEffectorType{}
// 	npcPropertyEffectorMap  = make(map[propertynpctypes.NPCPropertyEffectorType]PropertyEffectorFunc)
// )

// //注册npc属性作用器
// func RegisterNPCPropertyEffector(propertyEffectorType propertynpctypes.NPCPropertyEffectorType, pef PropertyEffectorFunc) {
// 	_, exist := npcPropertyEffectorMap[propertyEffectorType]
// 	if exist {
// 		panic("repeate register player property effector")
// 	}
// 	npcPropertyEffectorList = append(npcPropertyEffectorList, propertyEffectorType)
// 	npcPropertyEffectorMap[propertyEffectorType] = pef
// }

// //获取角色属性作用器
// func getNPCPropertyEffector(propertyEffectorType propertynpctypes.NPCPropertyEffectorType) PropertyEffectorFunc {
// 	pef, exist := npcPropertyEffectorMap[propertyEffectorType]
// 	if !exist {
// 		return nil
// 	}
// 	return pef
// }
