package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_GET_TYPE), (*uipb.CSOneArenaGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_GET_TYPE), (*uipb.SCOneArenaGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_ROB_TYPE), (*uipb.CSOneArenaRob)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_ROB_TYPE), (*uipb.SCOneArenaRob)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_ROBBED_PUSH_TYPE), (*uipb.SCOneArenaRobbedPush)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_RECORED_TYPE), (*uipb.CSOneArenaRecord)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_RECORED_TYPE), (*uipb.SCOneArenaRecord)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ONE_ARENA_SELL_TYPE), (*uipb.CSOneArenaSell)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_SELL_TYPE), (*uipb.SCOneArenaSell)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_ROB_RESULT), (*uipb.SCOneArenaRobResult)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ONE_ARENA_ROBOT_TYPE), (*uipb.SCOneArenaRobot)(nil))
}
