package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_FORCE_GET_TYPE), (*uipb.CSRankForceGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_FORCE_GET_TYPE), (*uipb.SCRankForceGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_MOUNT_GET_TYPE), (*uipb.CSRankMountGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_MOUNT_GET_TYPE), (*uipb.SCRankMountGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_WING_GET_TYPE), (*uipb.CSRankWingGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_WING_GET_TYPE), (*uipb.SCRankWingGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_BODYSHIELD_GET_TYPE), (*uipb.CSRankBodyShieldGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_BODYSHIELD_GET_TYPE), (*uipb.SCRankBodyShieldGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_WEAPON_GET_TYPE), (*uipb.CSRankWeaponGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_WEAPON_GET_TYPE), (*uipb.SCRankWeaponGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_GANG_GET_TYPE), (*uipb.CSRankGangGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_GANG_GET_TYPE), (*uipb.SCRankGangGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_MY_GET_TYPE), (*uipb.CSRankMyGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_MY_GET_TYPE), (*uipb.SCRankMyGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_SHENFA_GET_TYPE), (*uipb.CSRankShenFaGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_SHENFA_GET_TYPE), (*uipb.SCRankShenFaGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_LINGYU_GET_TYPE), (*uipb.CSRankLingYuGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_LINGYU_GET_TYPE), (*uipb.SCRankLingYuGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_FEATHER_GET_TYPE), (*uipb.CSRankFeatherGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_FEATHER_GET_TYPE), (*uipb.SCRankFeatherGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_SHIELD_GET_TYPE), (*uipb.CSRankShieldGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_SHIELD_GET_TYPE), (*uipb.SCRankShieldGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_ANQI_GET_TYPE), (*uipb.CSRankAnQiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_ANQI_GET_TYPE), (*uipb.SCRankAnQiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_FABAO_GET_TYPE), (*uipb.CSRankFaBaoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_FABAO_GET_TYPE), (*uipb.SCRankFaBaoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_XIANTI_GET_TYPE), (*uipb.CSRankXianTiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_XIANTI_GET_TYPE), (*uipb.SCRankXianTiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_SHIHUNFAN_GET_TYPE), (*uipb.CSRankShiHunFanGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_SHIHUNFAN_GET_TYPE), (*uipb.SCRankShiHunFanGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_TIANMOTI_GET_TYPE), (*uipb.CSRankTianMoTiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_TIANMOTI_GET_TYPE), (*uipb.SCRankTianMoTiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_LINGTONG_LEVEL_GET_TYPE), (*uipb.CSRankLingTongLevelGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_LINGTONG_LEVEL_GET_TYPE), (*uipb.SCRankLingTongLevelGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_LINGTONGDEV_GET_TYPE), (*uipb.CSRankLingTongDevGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_LINGTONGDEV_GET_TYPE), (*uipb.SCRankLingTongDevGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RANK_FEI_SHENG_GET_TYPE), (*uipb.CSRankFeiShengGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RANK_FEI_SHENG_GET_TYPE), (*uipb.SCRankFeiShengGet)(nil))
}
