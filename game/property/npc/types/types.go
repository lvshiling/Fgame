package types

//属性作用器效果
// type NPCPropertyEffectorType uint32

// const (
// 	NPCPropertyEffectorTypeInit NPCPropertyEffectorType = iota
// 	NPCPropertyEffectorTypeSkill
// 	NPCPropertyEffectorTypeBuff
// 	NPCPropertyEffectorTypeBuffPercent
// )

// var (
// 	npcPropertyEffectorStringMap = map[NPCPropertyEffectorType]string{
// 		NPCPropertyEffectorTypeInit:        "初始化",
// 		NPCPropertyEffectorTypeSkill:       "技能",
// 		NPCPropertyEffectorTypeBuff:        "buff",
// 		NPCPropertyEffectorTypeBuffPercent: "buff万分比",
// 	}
// )

// func (t NPCPropertyEffectorType) String() string {
// 	return npcPropertyEffectorStringMap[t]
// }

// func (t NPCPropertyEffectorType) EffectorType() uint32 {
// 	return uint32(t)
// }

// func (pet NPCPropertyEffectorType) Mask() uint32 {
// 	return 1 << uint(pet)
// }

// var (
// 	propertyEffectoryTypeSystemMap = map[NPCPropertyEffectorType]bool{
// 		NPCPropertyEffectorTypeInit:        true,
// 		NPCPropertyEffectorTypeSkill:       true,
// 		NPCPropertyEffectorTypeBuff:        false,
// 		NPCPropertyEffectorTypeBuffPercent: false,
// 	}
// )

// //是否是系统的
// func (pet NPCPropertyEffectorType) IsSystem() bool {
// 	return propertyEffectoryTypeSystemMap[pet]
// }

// var (
// 	propertyEffectoryTypePercentMap = map[NPCPropertyEffectorType]bool{
// 		NPCPropertyEffectorTypeInit:        false,
// 		NPCPropertyEffectorTypeSkill:       false,
// 		NPCPropertyEffectorTypeBuff:        false,
// 		NPCPropertyEffectorTypeBuffPercent: true,
// 	}
// )

// //是否是系统的
// func (pet NPCPropertyEffectorType) IsPercent() bool {
// 	return propertyEffectoryTypePercentMap[pet]
// }

// //所有作用器
// const (
// 	PropertyEffectorTypeMaskAll = 1<<32 - 1
// )
