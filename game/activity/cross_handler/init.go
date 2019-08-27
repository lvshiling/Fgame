package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE), (*crosspb.ISPlayerActivityPkDataChanged)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE), (*crosspb.SIPlayerActivityPkDataChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE), (*crosspb.ISPlayerActivityRankDataChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_TICKREW_DATA_CHANGED_TYPE), (*crosspb.ISPlayerActivityTickRewDataChanged)(nil))
}
