package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIENDS_GET_TYPE), (*uipb.SCFriendsGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_RECOMMENT_FRIENDS_GET_TYPE), (*uipb.CSRecommentFriendsGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_RECOMMENT_FRIENDS_GET_TYPE), (*uipb.SCRecommentFriendsGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_ADD_TYPE), (*uipb.CSFriendAdd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_ADD_TYPE), (*uipb.SCFriendAdd)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_BLACK_TYPE), (*uipb.CSFriendBlack)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_BLACK_TYPE), (*uipb.SCFriendBlack)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_DELETE_TYPE), (*uipb.CSFriendDelete)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_DELETE_TYPE), (*uipb.SCFriendDelete)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_SEARCH_TYPE), (*uipb.CSFriendSearch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_SEARCH_TYPE), (*uipb.SCFriendSearch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_GIFT_TYPE), (*uipb.CSFriendGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_GIFT_TYPE), (*uipb.SCFriendGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_GIFT_RECV_TYPE), (*uipb.SCFriendGiftRecv)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_GIFT_FEEDBACK_TYPE), (*uipb.CSFriendGiftFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_GIFT_FEEDBACK_TYPE), (*uipb.SCFriendGiftFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_GIFT_FEEDBACK_RECV_TYPE), (*uipb.SCFriendGiftFeedbackRecv)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_POINT_CHANGE_TYPE), (*uipb.SCFriendPointChange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_REMOVE_BLACK_TYPE), (*uipb.CSFriendRemoveBlack)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_REMOVE_BLACK_TYPE), (*uipb.SCFriendRemoveBlack)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_INVITE_TYPE), (*uipb.CSFriendInvite)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_INVITE_TYPE), (*uipb.SCFriendInvite)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_INVITE_PUSH_PEER_TYPE), (*uipb.SCFriendInvitePushPeer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_INVITE_REFUSE_PUSH_PEER_TYPE), (*uipb.SCFriendInviteRefusePushPeer)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_INVITE_LIST_TYPE), (*uipb.CSFriendInviteList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_INVITE_LIST_TYPE), (*uipb.SCFriendInviteList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_BATCH_TYPE), (*uipb.CSFriendBatch)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_BATCH_TYPE), (*uipb.SCFriendBatch)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_GIVE_FLOWER_LIGHT_TYPE), (*uipb.SCFriendGiveFlowerLight)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_NOTICE_BROADCASE_TYPE), (*uipb.SCFriendNoticeBroadcase)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_NOTICE_FEEDBACK_NOTICE_TYPE), (*uipb.SCFriendNoticeFeedbackNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_NOTICE_FEEDBACK_TYPE), (*uipb.CSFriendNoticeFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_NOTICE_FEEDBACK_TYPE), (*uipb.SCFriendNoticeFeedback)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_NOTICE_FEEDBACK_READ_TYPE), (*uipb.CSFriendNoticeFeedbackRead)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_NOTICE_FEEDBACK_READ_TYPE), (*uipb.SCFriendNoticeFeedbackRead)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_ADD_REW_TYPE), (*uipb.CSFriendAddRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_ADD_REW_TYPE), (*uipb.SCFriendAddRew)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_DUMMY_NUM_CHANGED_TYPE), (*uipb.SCFriendDummyFriendNumChanged)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_ADD_ALL_TYPE), (*uipb.CSFriendAddAll)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_ADD_ALL_TYPE), (*uipb.SCFriendAddAll)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_FEEDBACK_ASK_FOR_TYPE), (*uipb.CSFriendFeedbackAskFor)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_FEEDBACK_ASK_FOR_TYPE), (*uipb.SCFriendFeedbackAskFor)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_FEEDBACK_ASK_FOR_REPLY_TYPE), (*uipb.CSFriendFeedbackAskForReply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_FEEDBACK_ASK_FOR_REPLY_TYPE), (*uipb.SCFriendFeedbackAskForReply)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_FEEDBACK_ASK_FOR_NOTICE_TYPE), (*uipb.SCFriendFeedbackAskForNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_FEEDBACK_ASK_FOR_REPLY_NOTICE_TYPE), (*uipb.SCFriendFeedbackAskForReplyNotice)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_FRIEND_MARRY_DEVELOP_LOG_INCR_TYPE), (*uipb.CSFriendMarryDevelopLogIncr)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_FRIEND_MARRY_DEVELOP_LOG_INCR_TYPE), (*uipb.SCFriendMarryDevelopLogIncr)(nil))
}
