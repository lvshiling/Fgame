package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/processor"
)

func init() {
	initCodec()
	initProxy()
}

func initCodec() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_GET_TYPE), (*uipb.CSMountGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_GET_TYPE), (*uipb.SCMountGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_UNREALDAN_TYPE), (*uipb.CSMountUnrealDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_UNREALDAN_TYPE), (*uipb.SCMountUnrealDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_CULDAN_TYPE), (*uipb.CSMountCulDan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_CULDAN_TYPE), (*uipb.SCMountCulDan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_UNREAL_TYPE), (*uipb.CSMountUnreal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_UNREAL_TYPE), (*uipb.SCMountUnreal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_ADVANCED_TYPE), (*uipb.CSMountAdvanced)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_ADVANCED_TYPE), (*uipb.SCMountAdvanced)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_UNLOAD_TYPE), (*uipb.CSMountUnload)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_UNLOAD_TYPE), (*uipb.SCMountUnload)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_HIDDEN_TYPE), (*uipb.CSMountHidden)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_HIDDEN_TYPE), (*uipb.SCMountHidden)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MOUNT_UPSTAR_TYPE), (*uipb.CSMountUpstar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MOUNT_UPSTAR_TYPE), (*uipb.SCMountUpstar)(nil))
}

func initProxy() {
	processor.RegisterProxy(codec.MessageType(uipb.MessageType_CS_MOUNT_HIDDEN_TYPE))
}
