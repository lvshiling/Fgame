package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()

}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_IFNO_TYPE), (*uipb.SCShenMoSceneInfo)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_DATA_CHANGED_TYPE), (*uipb.SCShenMoSceneDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_END_TYPE), (*uipb.SCShenMoSceneEnd)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_GONGXUN_CHANGED_TYPE), (*uipb.SCPlayerGongXunChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_BIO_BROADCAST_TYPE), (*uipb.SCShenMoBioBroadcast)(nil))
}
