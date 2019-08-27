package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PING_TYPE), (*scenepb.CSPing)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PING_TYPE), (*scenepb.SCPing)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ENTER_SCENE_TYPE), (*scenepb.CSEnterScene)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ENTER_SCENE_TYPE), (*scenepb.SCEnterScene)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_ENTER_SCOPE_TYPE), (*scenepb.SCObjectEnterScope)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_EXIT_SCOPE_TYPE), (*scenepb.SCObjectExitScope)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_OBJECT_MOVE_TYPE), (*scenepb.CSObjectMove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_MOVE_TYPE), (*scenepb.SCObjectMove)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_OBJECT_ATTACK_TYPE), (*scenepb.CSObjectAttack)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_ATTACK_TYPE), (*scenepb.SCObjectAttack)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_DAMAGE_TYPE), (*scenepb.SCObjectDamage)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_TYPE), (*scenepb.SCObjectBuff)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BATTLE_TYPE), (*scenepb.SCObjectBattle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PLAYER_RELIVE_TYPE), (*scenepb.CSPlayerRelive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PLAYER_RELIVE_TYPE), (*scenepb.SCPlayerRelive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PLAYER_DATA_CHANGED_TYPE), (*scenepb.SCPlayerDataChanged)(nil))
	//buff
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_REMOVE_TYPE), (*scenepb.SCObjectBuffRemove)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ITEM_GET_TYPE), (*scenepb.CSItemGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ITEM_GET_TYPE), (*scenepb.SCItemGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ITEM_OWNER_CHANGED_TYPE), (*scenepb.SCItemOwnerChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_EXIT_SCENE_TYPE), (*scenepb.SCExitScene)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_FIXED_POSITION_TYPE), (*scenepb.SCObjectFixPosition)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_MONSTER_CAMP_CHANGED_TYPE), (*scenepb.SCMonsterCampChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_WORLD_ENTER_SCENE_TYPE), (*scenepb.CSWorldEnterScene)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ENTER_PORTAL_TYPE), (*scenepb.CSEnterPortal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ENTER_PORTAL_TYPE), (*scenepb.SCEnterPortal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_REENTER_SCENE_TYPE), (*scenepb.CSReenterScene)(nil))

	//副本退出
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FUBEN_EXIT_TYPE), (*uipb.CSFuBenExit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUBEN_EXIT_TYPE), (*uipb.SCFuBenExit)(nil))
	//跳转npc
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GO_TO_NPC_TYPE), (*uipb.CSGoToNPC)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GO_TO_NPC_TYPE), (*uipb.SCGoToNPC)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_KILLED_TYPE), (*uipb.SCPlayerKilled)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ATTACKED_TYPE), (*uipb.SCPlayerAttacked)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_SCENE_HEARTBEAT_TYPE), (*scenepb.CSSceneHeartBeat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_SCENE_HEARTBEAT_TYPE), (*scenepb.SCSceneHeartBeat)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_PLAYER_DATA_CHANGED_TYPE), (*uipb.SCScenePlayerDataChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PET_ATTACK_TYPE), (*scenepb.CSPetAttack)(nil))
	gamecodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PET_ATTACK_TYPE), (*scenepb.SCPetAttack)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_PLAYER_SKILL_USE_TYPE), (*uipb.SCScenePlayerSkillUse)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_KILL_TYPE), (*uipb.SCPlayerKill)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PLAYER_FEIXIE_TRANSFER_TYPE), (*uipb.CSPlayerFeiXieTransfer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_FEIXIE_TRANSFER_TYPE), (*uipb.SCPlayerFeiXieTransfer)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BUFF_LIST_TYPE), (*uipb.CSBuffList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BUFF_LIST_TYPE), (*uipb.SCBuffList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_RELIVE_TYPE), (*uipb.SCPlayerRelive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_EXIT_PVP_TYPE), (*uipb.SCPlayerExitPVP)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ENTER_PVP_TYPE), (*uipb.SCPlayerEnterPVP)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_RANK_CHANGED_TYPE), (*uipb.SCSceneRankChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BUFF_SEARCH_TYPE), (*uipb.CSBuffSearch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BUFF_SEARCH_TYPE), (*uipb.SCBuffSearch)(nil))
}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_PING_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_PING_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_ENTER_SCENE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_ENTER_SCENE_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_ENTER_SCOPE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_EXIT_SCOPE_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_OBJECT_MOVE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_MOVE_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_OBJECT_ATTACK_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_ATTACK_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_DAMAGE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_BATTLE_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_PLAYER_RELIVE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_PLAYER_RELIVE_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_PLAYER_DATA_CHANGED_TYPE))
	//buff
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_REMOVE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_ITEM_GET_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_ITEM_GET_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_ITEM_OWNER_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_EXIT_SCENE_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_OBJECT_FIXED_POSITION_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_MONSTER_CAMP_CHANGED_TYPE))
	// processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_WORLD_ENTER_SCENE_TYPE))

	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_ENTER_PORTAL_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_SC_ENTER_PORTAL_TYPE))
	processor.RegisterProxy(codec.MessageType(scenepb.MessageType_CS_REENTER_SCENE_TYPE))

	//副本退出
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_FUBEN_EXIT_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_FUBEN_EXIT_TYPE))
	//跳转npc
	// processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_GO_TO_NPC_TYPE))
	// processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_GO_TO_NPC_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_PLAYER_KILLED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_BUFF_LIST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_BUFF_LIST_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_BUFF_SEARCH_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_BUFF_SEARCH_TYPE))
}
