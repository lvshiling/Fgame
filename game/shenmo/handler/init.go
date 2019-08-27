package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENMO_CANCLE_LINEUP_TYPE), (*uipb.CSShenMoCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_CANCLE_LINEUP_TYPE), (*uipb.SCShenMoCancleLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_LINEUP_TYPE), (*uipb.SCShenMoLineUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_LINEUP_SUCCESS_TYPE), (*uipb.SCShenMoLineUpSuccess)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_FINISH_TO_LINEUP_TYPE), (*uipb.SCShenMoFinishToLineUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENMO_GET_REWARD_TYPE), (*uipb.CSShenMoGetReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_GET_REWARD_TYPE), (*uipb.SCShenMoGetReward)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENMO_MY_RANK_TYPE), (*uipb.CSShenMoMyRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_MY_RANK_TYPE), (*uipb.SCShenMoMyRank)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_SHENMO_RANK_GET_TYPE), (*uipb.CSShenMoRankGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_RANK_GET_TYPE), (*uipb.SCShenMoRankGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_IFNO_TYPE), (*uipb.SCShenMoSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_DATA_CHANGED_TYPE), (*uipb.SCShenMoSceneDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_END_TYPE), (*uipb.SCShenMoSceneEnd)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PLAYER_GONGXUN_CHANGED_TYPE), (*uipb.SCPlayerGongXunChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_SHENMO_BIO_BROADCAST_TYPE), (*uipb.SCShenMoBioBroadcast)(nil))
}

//TODO
func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_IFNO_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_DATA_CHANGED_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_SHENMO_SCENE_END_TYPE))
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_SC_SHENMO_BIO_BROADCAST_TYPE))
}
