package handler

import (
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PING_TYPE), (*scenepb.CSPing)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PING_TYPE), (*scenepb.SCPing)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ENTER_SCENE_TYPE), (*scenepb.CSEnterScene)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ENTER_SCENE_TYPE), (*scenepb.SCEnterScene)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_ENTER_SCOPE_TYPE), (*scenepb.SCObjectEnterScope)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_EXIT_SCOPE_TYPE), (*scenepb.SCObjectExitScope)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_OBJECT_MOVE_TYPE), (*scenepb.CSObjectMove)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_MOVE_TYPE), (*scenepb.SCObjectMove)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_OBJECT_ATTACK_TYPE), (*scenepb.CSObjectAttack)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_ATTACK_TYPE), (*scenepb.SCObjectAttack)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_DAMAGE_TYPE), (*scenepb.SCObjectDamage)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_TYPE), (*scenepb.SCObjectBuff)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BATTLE_TYPE), (*scenepb.SCObjectBattle)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PLAYER_RELIVE_TYPE), (*scenepb.CSPlayerRelive)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PLAYER_RELIVE_TYPE), (*scenepb.SCPlayerRelive)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PLAYER_DATA_CHANGED_TYPE), (*scenepb.SCPlayerDataChanged)(nil))
	//buff
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_REMOVE_TYPE), (*scenepb.SCObjectBuffRemove)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ITEM_GET_TYPE), (*scenepb.CSItemGet)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ITEM_GET_TYPE), (*scenepb.SCItemGet)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ITEM_OWNER_CHANGED_TYPE), (*scenepb.SCItemOwnerChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_EXIT_SCENE_TYPE), (*scenepb.SCExitScene)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_FIXED_POSITION_TYPE), (*scenepb.SCObjectFixPosition)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_MONSTER_CAMP_CHANGED_TYPE), (*scenepb.SCMonsterCampChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_REENTER_SCENE_TYPE), (*scenepb.CSReenterScene)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_SCENE_HEARTBEAT_TYPE), (*scenepb.SCSceneHeartBeat)(nil))

	//副本退出
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FUBEN_EXIT_TYPE), (*uipb.CSFuBenExit)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FUBEN_EXIT_TYPE), (*uipb.SCFuBenExit)(nil))
	//跳转npc
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GO_TO_NPC_TYPE), (*uipb.CSGoToNPC)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GO_TO_NPC_TYPE), (*uipb.SCGoToNPC)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_KILL_TYPE), (*uipb.SCPlayerKill)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_KILLED_TYPE), (*uipb.SCPlayerKilled)(nil))

	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ATTACKED_TYPE), (*uipb.SCPlayerAttacked)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_PLAYER_DATA_CHANGED_TYPE), (*uipb.SCScenePlayerDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_PLAYER_SKILL_USE_TYPE), (*uipb.SCScenePlayerSkillUse)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BUFF_LIST_TYPE), (*uipb.CSBuffList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BUFF_LIST_TYPE), (*uipb.SCBuffList)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_RELIVE_TYPE), (*uipb.SCPlayerRelive)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_EXIT_PVP_TYPE), (*uipb.SCPlayerExitPVP)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_ENTER_PVP_TYPE), (*uipb.SCPlayerEnterPVP)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SCENE_RANK_CHANGED_TYPE), (*uipb.SCSceneRankChanged)(nil))

	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ENTER_PORTAL_TYPE), (*scenepb.CSEnterPortal)(nil))
	crosscodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ENTER_PORTAL_TYPE), (*scenepb.SCEnterPortal)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BUFF_SEARCH_TYPE), (*uipb.CSBuffSearch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BUFF_SEARCH_TYPE), (*uipb.SCBuffSearch)(nil))
}
