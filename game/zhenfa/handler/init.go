package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENFA_GET_TYPE), (*uipb.CSZhenFaGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENFA_GET_TYPE), (*uipb.SCZhenFaGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENFA_ACTIVATE_TYPE), (*uipb.CSZhenFaActivate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENFA_ACTIVATE_TYPE), (*uipb.SCZhenFaActivate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENQI_GET_TYPE), (*uipb.CSZhenQiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENQI_GET_TYPE), (*uipb.SCZhenQiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENQIXIANHUO_GET_TYPE), (*uipb.CSZhenQiXianHuoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENQIXIANHUO_GET_TYPE), (*uipb.SCZhenQiXianHuoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENFA_SHENGJI_TYPE), (*uipb.CSZhenFaShengJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENFA_SHENGJI_TYPE), (*uipb.SCZhenFaShengJi)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENQI_ADVANCED_TYPE), (*uipb.CSZhenQiAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENQI_ADVANCED_TYPE), (*uipb.SCZhenQiAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ZHENQI_XIANHUO_SHENGJI_TYPE), (*uipb.CSZhenQiXianHuoShengJi)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ZHENQI_XIANHUO_SHENGJI_TYPE), (*uipb.SCZhenQiXianHuoShengJi)(nil))
}
