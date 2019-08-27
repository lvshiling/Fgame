package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_IS_PLAYER_GET_DROP_ITEM_TYPE), (*crosspb.ISPlayerGetDropItem)(nil))
	gamecodec.RegisterMsg(codec.MessageType(crosspb.MessageType_SI_PLAYER_GET_DROP_ITEM_TYPE), (*crosspb.SIPlayerGetDropItem)(nil))
}
