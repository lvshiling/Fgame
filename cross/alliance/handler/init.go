package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()

}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_ALLIANCE_MEMBER_POS_TYPE), (*uipb.CSAllianceMemberPos)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ALLIANCE_MEMBER_POS_TYPE), (*uipb.SCAllianceMemberPos)(nil))
}
