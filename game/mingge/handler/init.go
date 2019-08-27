package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_PAN_GET_TYPE), (*uipb.CSMingGePanGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_PAN_GET_TYPE), (*uipb.SCMingGePanGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_SYNTHESIS_TYPE), (*uipb.CSMingGeSynthesis)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_SYNTHESIS_TYPE), (*uipb.SCMingGeSynthesis)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_REFINED_GET_TYPE), (*uipb.CSMingGeRefinedGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_REFINED_GET_TYPE), (*uipb.SCMingGeRefinedGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_PAN_UNLOAD_TYPE), (*uipb.CSMingGePanUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_PAN_UNLOAD_TYPE), (*uipb.SCMingGePanUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_REFINED_TYPE), (*uipb.CSMingGeRefined)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_REFINED_TYPE), (*uipb.SCMingGeRefined)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_MINGLI_GET_TYPE), (*uipb.CSMingGeMingLiGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_MINGLI_GET_TYPE), (*uipb.SCMingGeMingLiGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_MINGLI_BAPTIZE_TYPE), (*uipb.CSMingGeMingLiBaptize)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_MINGLI_BAPTIZE_TYPE), (*uipb.SCMingGeMingLiBaptize)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MINGGE_PAN_MOSAIC_TYPE), (*uipb.CSMingGePanMosaic)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_PAN_MOSAIC_TYPE), (*uipb.SCMingGePanMosaic)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MINGGE_MINGGONG_TYPE), (*uipb.SCMingGeMingGongActive)(nil))
}
