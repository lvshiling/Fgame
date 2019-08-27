package npc

import (
	_ "fgame/fgame/game/npc/ai"
	"fgame/fgame/game/npc/npc"
	_ "fgame/fgame/game/npc/use"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterNPC(types.BiologyScriptTypeCrossBigEgg, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossSmallEgg, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossBigBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossSmallBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeCrossLianYuBoss, scene.NPCFactoryFunc(npc.CreateNPC))
	scene.RegisterNPC(types.BiologyScriptTypeGodSiegeBoss, scene.NPCFactoryFunc(npc.CreateNPC))
}
