package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()
}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE), (*crosspb.ISPlayerActivityPkDataChanged)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_ACTIVITY_PK_DATA_CHANGED_TYPE), (*crosspb.SIPlayerActivityPkDataChanged)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_RANK_DATA_CHANGED_TYPE), (*crosspb.ISPlayerActivityRankDataChanged)(nil))

	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_ACTIVITY_TICKREW_DATA_CHANGED_TYPE), (*crosspb.ISPlayerActivityTickRewDataChanged)(nil))
}
