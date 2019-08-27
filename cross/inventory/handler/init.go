package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GET_DROP_ITEM_TYPE), (*crosspb.ISPlayerGetDropItem)(nil))
	crosscodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_GET_DROP_ITEM_TYPE), (*crosspb.SIPlayerGetDropItem)(nil))
}
