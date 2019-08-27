package handler

import (
	clientcodec "fgame/fgame/client/codec"
	"fgame/fgame/common/codec"
	scenepb "fgame/fgame/common/codec/pb/scene"
)

func init() {
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PING_TYPE), (*scenepb.CSPing)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PING_TYPE), (*scenepb.SCPing)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ENTER_SCENE_TYPE), (*scenepb.CSEnterScene)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ENTER_SCENE_TYPE), (*scenepb.SCEnterScene)(nil))

	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_ENTER_SCOPE_TYPE), (*scenepb.SCObjectEnterScope)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_EXIT_SCOPE_TYPE), (*scenepb.SCObjectExitScope)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_OBJECT_MOVE_TYPE), (*scenepb.CSObjectMove)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_MOVE_TYPE), (*scenepb.SCObjectMove)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_OBJECT_ATTACK_TYPE), (*scenepb.CSObjectAttack)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_ATTACK_TYPE), (*scenepb.SCObjectAttack)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_DAMAGE_TYPE), (*scenepb.SCObjectDamage)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_TYPE), (*scenepb.SCObjectBuff)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BATTLE_TYPE), (*scenepb.SCObjectBattle)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PLAYER_DATA_CHANGED_TYPE), (*scenepb.SCPlayerDataChanged)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_PLAYER_RELIVE_TYPE), (*scenepb.CSPlayerRelive)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_PLAYER_RELIVE_TYPE), (*scenepb.SCPlayerRelive)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_BUFF_REMOVE_TYPE), (*scenepb.SCObjectBuffRemove)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ITEM_GET_TYPE), (*scenepb.CSItemGet)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ITEM_GET_TYPE), (*scenepb.SCItemGet)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_EXIT_SCENE_TYPE), (*scenepb.SCExitScene)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_OBJECT_FIXED_POSITION_TYPE), (*scenepb.SCObjectFixPosition)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_MONSTER_CAMP_CHANGED_TYPE), (*scenepb.SCMonsterCampChanged)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ITEM_OWNER_CHANGED_TYPE), (*scenepb.SCItemOwnerChanged)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_WORLD_ENTER_SCENE_TYPE), (*scenepb.CSWorldEnterScene)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_ENTER_PORTAL_TYPE), (*scenepb.CSEnterPortal)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_SC_ENTER_PORTAL_TYPE), (*scenepb.SCEnterPortal)(nil))
	clientcodec.RegisterMsg(codec.MessageType(scenepb.MessageType_CS_REENTER_SCENE_TYPE), (*scenepb.CSReenterScene)(nil))
}
