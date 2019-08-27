package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_LEVEL_TYPE), (*uipb.SCBaGuaLevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BAGUA_TOKILL_TYPE), (*uipb.CSBaGuaToKill)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_TOKILL_TYPE), (*uipb.SCBaGuaToKill)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_TOKILL_RESULT_TYPE), (*uipb.SCBaGuaToKillResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_SCENEINFO_TYPE), (*uipb.SCBaGuaSceneInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BAGUA_PAIR_TYPE), (*uipb.CSBaGuaPair)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_PAIR_TYPE), (*uipb.SCBaGuaPair)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BAGUA_PAIR_CANCLE_TYPE), (*uipb.CSBaGuaPairCancle)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_PAIR_CANCLE_TYPE), (*uipb.SCBaGuaPairCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_PAIR_PUSH_SPOUSE_TYPE), (*uipb.SCBaGuaPairPushSpouse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_PAIR_PUSH_CANCLE_TYPE), (*uipb.SCBaGuaPairPushCancle)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BAGUA_PAIR_DEAL_TYPE), (*uipb.CSBaGuaPairDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_PAIR_DEAL_TYPE), (*uipb.SCBaGuaPairDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_SPOUSE_REFUSED_TYPE), (*uipb.SCBaGuaSpouseRefused)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_PAIR_RESULT_TYPE), (*uipb.SCBaGuaPairResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_BAGUA_NEXT_TYPE), (*uipb.CSBaGuaNext)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_NEXT_TYPE), (*uipb.SCBaGuaNext)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_BAGUA_INVITE_OFFONLINE_TYPE), (*uipb.SCBaGuaInviteOffonline)(nil))
}
