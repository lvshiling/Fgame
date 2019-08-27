package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MATERIAL_INFO_GET_TYPE), (*uipb.CSMaterialInfoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MATERIAL_INFO_GET_TYPE), (*uipb.SCMaterialInfoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MATERIAL_SAO_DANG_TYPE), (*uipb.CSMaterialSaoDang)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MATERIAL_SAO_DANG_TYPE), (*uipb.SCMaterialSaoDang)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MATERIAL_CHALLENGE_TYPE), (*uipb.CSMaterialChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MATERIAL_CHALLENGE_TYPE), (*uipb.SCMaterialChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MATERIAL_SCENE_INFO_TYPE), (*uipb.SCMaterialSceneInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MATERIAL_REFRESH_BIOLOGY_TYPE), (*uipb.SCMaterialRefreshBiology)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MATERIAL_CHALLENGE_RESULT_TYPE), (*uipb.SCMaterialChallengeResult)(nil))

}
