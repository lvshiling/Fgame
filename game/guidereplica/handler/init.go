package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GUIDE_REPLICA_CHALLENGE_TYPE), (*uipb.CSGuideReplicaChallenge)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUIDE_REPLICA_CHALLENGE_TYPE), (*uipb.SCGuideReplicaChallenge)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUIDE_REPLICA_CHALLENGE_RESULT_TYPE), (*uipb.SCGuideReplicaChallengeResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUIDE_REPLICA_SCENE_INFO_TYPE), (*uipb.SCGuideReplicaSceneInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUIDE_REPLICA_SCENE_DATA_CHANGED_NOTICE_TYPE), (*uipb.SCGuideReplicaSceneDataChangedNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GUIDE_REPLICA_PLAYER_COMMON_OPERATE_TYPE), (*uipb.CSGuideReplicaPlayerCommonOperate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GUIDE_REPLICA_PLAYER_COMMON_OPERATE_TYPE), (*uipb.SCGuideReplicaPlayerCommonOperate)(nil))
}
