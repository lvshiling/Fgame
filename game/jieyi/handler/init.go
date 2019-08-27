package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_MEMBER_INFO_TYPE), (*uipb.CSJieYiMemberInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_MEMBER_INFO_TYPE), (*uipb.SCJieYiMemberInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_PLAYER_INFO_TYPE), (*uipb.SCJieYiPlayerInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_INVITE_TYPE), (*uipb.CSJieYiInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_INVITE_TYPE), (*uipb.SCJieYiInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_INVITE_NOTICE_TYPE), (*uipb.SCJieYiInviteNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_HELP_INVITE_TYPE), (*uipb.CSJieYiHandleInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_HELP_INVITE_TYPE), (*uipb.SCJieYiHandleInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_HELP_INVITE_NOTICE_TYPE), (*uipb.SCJieYiHandleInviteNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_NAME_UP_LEV_TYPE), (*uipb.CSJieYiNameUpLev)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_NAME_UP_LEV_TYPE), (*uipb.SCJieYiNameUpLev)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_ACTIVITE_TYPE), (*uipb.CSJieYiTokenActivite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_ACTIVITE_TYPE), (*uipb.SCJieYiTokenActivite)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_INFO_ON_CHANGE_TYPE), (*uipb.SCJieYiInfoOnChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_BROTHER_INFO_ON_CHANGE_TYPE), (*uipb.SCJieYiBrotherInfoOnChange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_UP_LEV_TYPE), (*uipb.CSJieYiTokenUpLev)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_UP_LEV_TYPE), (*uipb.SCJieYiTokenUpLev)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_CHANGE_TYPE), (*uipb.CSJieYiTokenChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_CHANGE_TYPE), (*uipb.SCJieYiTokenChange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_POST_TYPE), (*uipb.CSJieYiPost)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_POST_TYPE), (*uipb.SCJieYiPost)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_LEAVE_WORD_INFO_TYPE), (*uipb.CSJieYiLeaveWordInfo)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_LEAVE_WORD_INFO_TYPE), (*uipb.SCJieYiLeaveWordInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_GIVE_TYPE), (*uipb.CSJieYiTokenGive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_GIVE_TYPE), (*uipb.SCJieYiTokenGive)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_GIVE_NOTICE_TYPE), (*uipb.SCJieYiTokenGiveNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_JIE_CHU_TYPE), (*uipb.CSJieYiJieChu)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_JIE_CHU_TYPE), (*uipb.SCJieYiJieChu)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_JIE_CHU_NOTICE_TYPE), (*uipb.SCJieYiJieChuNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_DAOJU_CHANGE_TYPE), (*uipb.CSJieYiDaoJuChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_DAOJU_CHANGE_TYPE), (*uipb.SCJieYiDaoJuChange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_DAOJU_HELP_CHANGE_TYPE), (*uipb.CSJieYiDaoJuHelpChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_DAOJU_HELP_CHANGE_TYPE), (*uipb.SCJieYiDaoJuHelpChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_DAOJU_HELP_CHANGE_NOTICE_TYPE), (*uipb.SCJieYiDaoJuHelpChangeNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_QIU_YUAN_TYPE), (*uipb.CSJieYiQiuYuan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_QIU_YUAN_TYPE), (*uipb.SCJieYiQiuYuan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_QIU_YUAN_NOTICE_TYPE), (*uipb.SCJieYiQiuYuanNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_JIU_YUAN_TYPE), (*uipb.CSJieYiJiuYuan)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_JIU_YUAN_TYPE), (*uipb.SCJieYiJiuYuan)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_LAO_DA_TI_REN_TYPE), (*uipb.CSJieYiLaoDaTiRen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_LAO_DA_TI_REN_TYPE), (*uipb.SCJieYiLaoDaTiRen)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_LAO_DA_TI_REN_NOTICE_TYPE), (*uipb.SCJieYiLaoDaTiRenNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_LAO_DA_TI_REN_OTHER_NOTICE_TYPE), (*uipb.SCJieYiLaoDaTiRenOtherNotice)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_TOKEN_SUO_YAO_TYPE), (*uipb.CSJieYiTokenSuoYao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_SUO_YAO_TYPE), (*uipb.SCJieYiTokenSuoYao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_TOKEN_SUO_YAO_NOTICE_TYPE), (*uipb.SCJieYiTokenSuoYaoNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_JIEYI_HANDLE_TOKEN_SUO_YAO_TYPE), (*uipb.CSJieYiHandleTokenSuoYao)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_HANDLE_TOKEN_SUO_YAO_TYPE), (*uipb.SCJieYiHandleTokenSuoYao)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_SHENG_WEI_ZHI_DROP_TYPE), (*uipb.SCJieYiShengWeiZhiDrop)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_JIEYI_SHENG_WEI_ZHI_TUI_SONG_TYPE), (*uipb.SCJieYiShengWeiZhiTuiSong)(nil))
}
