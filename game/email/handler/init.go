package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_EMAILS_GET_TYPE), (*uipb.CSEmailsGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_EMAILS_GET_TYPE), (*uipb.SCEmailsGet)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_READ_EMAIL_TYPE), (*uipb.CSReadEmail)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_READ_EMAIL_TYPE), (*uipb.SCReadEmail)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DEL_EMAIL_TYPE), (*uipb.CSDelEmail)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DEL_EMAIL_TYPE), (*uipb.SCDelEmail)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_DEL_EMAIL_BATCH_TYPE), (*uipb.CSDelHadReadEmail)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_DEL_EMAIL_BATCH_TYPE), (*uipb.SCDelHadReadEmail)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GET_ATTACHMENT_BATCH_TYPE), (*uipb.CSGetAttachmentBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GET_ATTACHMENT_BATCH_TYPE), (*uipb.SCGetAttachmentBatch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_GET_ATTACHMENT_TYPE), (*uipb.CSGetAttachment)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_GET_ATTACHMENT_TYPE), (*uipb.SCGetAttachment)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_ADD_EMAIL_TYPE), (*uipb.SCAddEmail)(nil))
}
