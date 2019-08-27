package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_LIST_TYPE), (*uipb.SCQuestList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_ACCEPT_TYPE), (*uipb.CSQuestAccept)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_ACCEPT_TYPE), (*uipb.SCQuestAccept)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_COMMIT_TYPE), (*uipb.CSQuestCommit)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_COMMIT_TYPE), (*uipb.SCQuestCommit)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_UPDATE_TYPE), (*uipb.SCQuestUpdate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_NPC_DIALOG_TYPE), (*uipb.CSQuestNPCDialog)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_NPC_DIALOG_TYPE), (*uipb.SCQuestNPCDialog)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_GATHER_TYPE), (*uipb.CSQuestGather)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_GATHER_TYPE), (*uipb.SCQuestGather)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_TUMONUM_GET_TYPE), (*uipb.CSQuestTuMoNumGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_TUMONUM_GET_TYPE), (*uipb.SCQuestTuMoNumGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_USETOKEN_TYPE), (*uipb.CSQuestTuMoUseToken)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_TUMO_USETOKEN_TYPE), (*uipb.SCQuestTuMoUseToken)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_DISCARD_TYPE), (*uipb.CSQuestTuMoDiscard)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_TUMO_DISCARD_TYPE), (*uipb.SCQuestTuMoDiscard)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_BUYNUM_TYPE), (*uipb.CSQuestTuMoBuyNum)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_TUMO_BUYNUM_TYPE), (*uipb.SCQuestTuMoBuyNum)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_FINISH_ALL_TYPE), (*uipb.CSQuestTuMoFinishAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_TUMO_FINISH_ALL_TYPE), (*uipb.SCQuestTuMoFinishAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_TUMO_IMMEDIATE_FINISH_TYPE), (*uipb.CSQuestTuMoImmediate)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_TUMO_IMMEDIATE_FINISH_TYPE), (*uipb.SCQuestTuMoImmediate)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_FEIXIE_TYPE), (*uipb.CSQuestFeiXie)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_FEIXIE_TYPE), (*uipb.SCQuestFeiXie)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_DAILY_FINISH_ALL_TYPE), (*uipb.CSQuestDailyFinishAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_DAILY_FINISH_ALL_TYPE), (*uipb.SCQuestDailyFinishAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_DAILY_SEQ_TYPE), (*uipb.SCQuestDailySeq)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_KAIFUMUBIAO_GET_TYPE), (*uipb.CSQuestKaiFuMuBiaoGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_KAIFUMUBIAO_GET_TYPE), (*uipb.SCQuestKaiFuMuBiaoGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_KAIFUMUBIAO_RECEIVE_TYPE), (*uipb.CSQuestkaiFuMuBiaoReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_KAIFUMUBIAO_RECEIVE_TYPE), (*uipb.SCQuestkaiFuMuBiaoReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_QIYU_NOTICE_TYPE), (*uipb.SCQuestQiYuNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_QIYU_RECEIVE_TYPE), (*uipb.CSQuestQiYuReceive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_QIYU_RECEIVE_TYPE), (*uipb.SCQuestQiYuReceive)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_DAILY_COMMIT_REW_TYPE), (*uipb.SCQuestDailyCommitRew)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_QUEST_DAILY_FINISH_ONECE_TYPE), (*uipb.CSQuestDailyFinishOnce)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_QUEST_DAILY_FINISH_ONECE_TYPE), (*uipb.SCQuestDailyFinishOnce)(nil))

}
