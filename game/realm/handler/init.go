package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_LEVEL_TYPE), (*uipb.SCRealmLevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REALM_TOKILL_TYPE), (*uipb.CSRealmToKill)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_TOKILL_TYPE), (*uipb.SCRealmToKill)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_TOKILL_RESULT_TYPE), (*uipb.SCRealmToKillResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_SCENEINFO_TYPE), (*uipb.SCRealmSceneInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REALM_RANK_GET_TYPE), (*uipb.CSRealmRankGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_RANK_GET_TYPE), (*uipb.SCRealmRankGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REALM_PAIR_TYPE), (*uipb.CSRealmPair)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_PAIR_TYPE), (*uipb.SCRealmPair)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REALM_PAIR_CANCLE_TYPE), (*uipb.CSRealmPairCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_PAIR_CANCLE_TYPE), (*uipb.SCRealmPairCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_PAIR_PUSH_SPOUSE_TYPE), (*uipb.SCRealmPairPushSpouse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_PAIR_PUSH_CANCLE_TYPE), (*uipb.SCRealmPairPushCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REALM_PAIR_DEAL_TYPE), (*uipb.CSRealmPairDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_PAIR_DEAL_TYPE), (*uipb.SCRealmPairDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_SPOUSE_REFUSED_TYPE), (*uipb.SCRealmSpouseRefused)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_PAIR_RESULT_TYPE), (*uipb.SCRealmPairResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_REALM_NEXT_TYPE), (*uipb.CSRealmNext)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_NEXT_TYPE), (*uipb.SCRealmNext)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_REALM_INVITE_OFFONLINE_TYPE), (*uipb.SCRealmInviteOffonline)(nil))
}
