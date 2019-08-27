package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_CARD_GET_TYPE), (*uipb.CSQuestSecretCardGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_SECRET_CARD_GET_TYPE), (*uipb.SCQuestSecretCardGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_SPY_TYPE), (*uipb.CSQuestSecretSpy)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_SECRET_SPY_TYPE), (*uipb.SCQuestSecretSpy)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_PICKUP_TYPE), (*uipb.CSQuestSecretPickUp)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_SECRET_PICKUP_TYPE), (*uipb.SCQuestSecretPickUp)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_DISCARD_TYPE), (*uipb.CSQuestSecretDiscard)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_STAR_REW_TYPE), (*uipb.CSQuestSecretStarRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_SECRET_STAR_REW_TYPE), (*uipb.SCQuestSecretStarRew)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_FINISH_TYPE), (*uipb.CSQuestSecretFinish)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_SECRET_FINISH_TYPE), (*uipb.SCQuestSecretFinish)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_SECRET_IMMEDIATE_FINISH_TYPE), (*uipb.CSQuestSecretImmediate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_SECRET_IMMEDIATE_FINISH_TYPE), (*uipb.SCQuestSecretImmediate)(nil))
}
