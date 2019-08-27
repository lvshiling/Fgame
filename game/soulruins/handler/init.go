package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_GET_TYPE), (*uipb.CSSoulRuinsGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_GET_TYPE), (*uipb.SCSoulRuinsGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_CHALLENGE_TYPE), (*uipb.CSSoulRuinsChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_CHALLENGE_TYPE), (*uipb.SCSoulRuinsChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_EVENT_TYPE), (*uipb.SCSoulRuinsEvent)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_DEALEVENT_TYPE), (*uipb.CSSoulRuinsDealEvent)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_DEALEVENT_TYPE), (*uipb.SCSoulRuinsDealEvent)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_SCENEINFO_TYPE), (*uipb.SCSoulRuinsSceneInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_FORCEGET_TYPE), (*uipb.CSSoulRuinsForceGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_FIRSTPASS_TYPE), (*uipb.SCSoulRuinsFirstPass)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_RESULT_TYPE), (*uipb.SCSoulRuinsResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_SWEEP_TYPE), (*uipb.CSSoulRuinsSweep)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_SWEEP_TYPE), (*uipb.SCSoulRuinsSweep)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_REWRECEIVE_TYPE), (*uipb.CSSoulRuinsRewReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_REWRECEIVE_TYPE), (*uipb.SCSoulRuinsRewReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_BUYNUM_TYPE), (*uipb.CSSoulRuinsBuyNum)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SOULRUINS_BUYNUM_TYPE), (*uipb.SCSoulRuinsBuyNum)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_FINISH_ALL), (*uipb.CSSoulRuinsFinishAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SOULRUINS_NEXT_LEVEL_TYPE), (*uipb.CSSoulRuinsNext)(nil))
}
