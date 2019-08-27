package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	gamecodec "fgame/fgame/game/codec"
)

func init() {
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_GET_TYPE), (*uipb.CSMarryGet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_INFO_TYPE), (*uipb.SCMarryInfo)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_PROPOSAL_TYPE), (*uipb.CSMarryProposal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PROPOSAL_TYPE), (*uipb.SCMarryProposal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PUSH_PROPOSAL_TYPE), (*uipb.SCMarryPushProposal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_PROPOSAL_DEAL_TYPE), (*uipb.CSMarryProposalDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PROPOSAL_RESULT_TYPE), (*uipb.SCMarryProposalResult)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_DIVORCE_TYPE), (*uipb.CSMarryDivorce)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DIVORCE_TYPE), (*uipb.SCMarryDivorce)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PUSH_DIVORCE_TYPE), (*uipb.SCMarryPushDivorce)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_DIVORCE_DEAL_TYPE), (*uipb.CSMarryDivorceDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_RING_FEED_TYPE), (*uipb.CSMarryRingFeed)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_RING_FEED_TYPE), (*uipb.SCMarryRingFeed)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_SPOUSE_INFO_CHANGE_TYPE), (*uipb.SCMarrySpouseInfoChange)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_INFO_CHANGE_TYPE), (*uipb.SCMarryInfoChange)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_WED_LIST_TYPE), (*uipb.CSMarryWedList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_LIST_TYPE), (*uipb.SCMarryWedList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GRADE_TYPE), (*uipb.CSMarryWedGrade)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_GRADE_TYPE), (*uipb.SCMarryWedGrade)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_BANQUET_TYPE), (*uipb.SCMarryBanquet)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_HEROISM_TOP_THREE_TYPE), (*uipb.SCMarryHeroismTopThree)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GIFT_TYPE), (*uipb.CSMarryWedGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_GIFT_TYPE), (*uipb.SCMarryWedGift)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GIFT_LIST_TYPE), (*uipb.CSMarryWedGiftList)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_GIFT_LIST_TYPE), (*uipb.SCMarryWedGiftList)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_VIEW_WED_CARD_TYPE), (*uipb.CSMarryViewWedCard)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_VIEW_WED_CARD_TYPE), (*uipb.SCMarryViewWedCard)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PUSH_WED_CARD_TYPE), (*uipb.SCMarryPushWedCard)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_AFTER_LOGIN_TYPE), (*uipb.SCMarryAfterLogin)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DIVORCE_DEAL_PUSH_PEER_TYPE), (*uipb.SCMarryDivorceDealPushPeer)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_RING_REPLACE_TYPE), (*uipb.CSMarryRingReplace)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_RING_REPLACE_TYPE), (*uipb.SCMarryRingReplace)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_TREE_FEED_TYPE), (*uipb.CSMarryTreeFeed)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_TREE_FEED_TYPE), (*uipb.SCMarryTreeFeed)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_PUSH_STATUS_TYPE), (*uipb.SCMarryWedPushStatus)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_WED_TRANSFER_TYPE), (*uipb.CSMarryWedTransfer)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_CLICK_CAR_TYPE), (*uipb.CSMarryClickCar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_CLICK_CAR_TYPE), (*uipb.SCMarryClickCar)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WORSHIP_TYPE), (*uipb.SCMarryWorship)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_CANCEL_TYPE), (*uipb.SCMarryCancel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_END_TYPE), (*uipb.SCMarryWedEnd)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_RECOMMENT_SPOUSES_TYPE), (*uipb.CSMarryRecommended)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_RECOMMENT_SPOUSES_TYPE), (*uipb.SCMarryRecommended)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_CANCLE_TO_OTHER_TYPE), (*uipb.SCMarryCancelToOther)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_GRADE_TO_SPOUSE_TYPE), (*uipb.SCMarryWedGradeToSpouse)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_WED_GRADE_SPOUSE_DEAL_TYPE), (*uipb.CSMarryWedGradeSpouseDeal)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_GRADE_SPOUSE_DEAL_TYPE), (*uipb.SCMarryWedGradeSpouseDeal)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_GRADE_REFUSE_TO_PEER_TYPE), (*uipb.SCMarryWedGradeRefuseToPeer)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_WED_SUCESS_TYPE), (*uipb.SCMarryWedSucess)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_PRE_GIFT), (*uipb.CSMarryPreGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PRE_GIFT), (*uipb.SCMarryPreGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_PRE_GIFT_NOTICE), (*uipb.SCMarryPreGiftMsg)(nil))
	//结婚纪念
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_JINIAN_GET), (*uipb.CSMarryJiNianMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_JINIAN_GET), (*uipb.SCMarryJiNianMsg)(nil))

	// gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_DEVELOP_SEND_GIFT_TYPE), (*uipb.CSMarryDevelopSendGift)(nil))
	// gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DEVELOP_SEND_GIFT_TYPE), (*uipb.SCMarryDevelopSendGift)(nil))
	// gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_DEVELOP_SEND_GIFT_TYPE), (*uipb.CSMarryDevelopSendGift)(nil))
	// gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DEVELOP_SEND_GIFT_TYPE), (*uipb.SCMarryDevelopSendGift)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_DEVELOP_UPLEVEL_TYPE), (*uipb.CSMarryDevelopUplevel)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DEVELOP_UPLEVEL_TYPE), (*uipb.SCMarryDevelopUplevel)(nil))

	//定情信物
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MarryDingQingJiHuo), (*uipb.CSMarryDingQingJiHuoMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MarryDingQingJiHuo), (*uipb.SCMarryDingQingJiHuoMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MarryDingQingSuoYao), (*uipb.CSMarryDingQingSuoYaoMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MarryDingQingSuoYao), (*uipb.SCMarryDingQingSuoYaoMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MarryDingQingSuoYaoRsp), (*uipb.SCMarryDingQingSuoYaoRspMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MarryDingQingSuoYaoDeal), (*uipb.CSMarryDingQingSuoYaoDealMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MarryDingQingSuoYaoDeal), (*uipb.SCMarryDingQingSuoYaoDealMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MarryDingQingZengSong), (*uipb.CSMarryDingQingZengSongMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MarryDingQingZengSong), (*uipb.SCMarryDingQingZengSongMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_BANQUET_SET_CHANGE_TYPE), (*uipb.SCMarryBanquetSetChangeMsg)(nil))

	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_MARRY_DINGQING_YUE_TYPE), (*uipb.CSMarryDingQingYueMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DINGQING_YUE_TYPE), (*uipb.SCMarryDingQingYueMsg)(nil))
	gamecodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_MARRY_DINGQING_YUE_SPOUSE_TYPE), (*uipb.SCMarryDingQingYueSpouseMsg)(nil))
}
