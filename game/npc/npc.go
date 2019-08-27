package npc

import (
	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/types"
)

import (
	_ "fgame/fgame/game/npc/ai"
	_ "fgame/fgame/game/npc/event/listener"
	_ "fgame/fgame/game/npc/use"
)

func init() {
	scene.RegisterNPC(types.BiologyScriptTypeMonster, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeBattleNPC, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeRobber, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeBiaoChe, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeFourGodSpecial, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeFourGodCollect, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeRelivePoint, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeXianMengNPC, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeBuildingMonster, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeSoulBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeFourGodBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeWorldBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeWedBanquet, scene.NPCFactoryFunc(npc.CreateNPC))
	// scene.RegisterNPC(types.BiologyScriptTypeWeddingCar, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeOneArenaGuardian, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeArenaExpTree, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeArenaShengShou, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeArenaTreasure, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossWorldBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeBossCallTicket, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossBigBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossSmallBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeTowerMonster, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeTowerBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeMyBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeVIPMyBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossLianYuBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeGodSiegeBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeUnrealBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeOutlandBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCangJingGeBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeAllianceShengTan, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeAllianceBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeShenYuBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeLongGongBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeQiYuDao, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeZhenXiBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeDingShiBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeArenaBossHuWei, scene.NPCFactoryFunc(npc.CreateNPC))
}
