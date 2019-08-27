package lang

const (
	FriendIsNotFriend LangCode = FriendBase + iota
	FriendAlreadyFriend
	FriendAlreadyDelete
	FriendAlreadyFull
	FriednBlackAlreadyFull
	FriendPeerAlreadFull
	FriendIsSpouseNoDelete
	FriendIsBlack
	FriendAddPeerBlack
	FriendNotInBlack
	FriendBatchAgreeFull
	FriendGiveFlowersNotice
	FriendAddRewHadDone
	FriendNumNotEnough
	FriendAskForGiveTitle
	FriendAskForGiveContent
	FriendGiveBiaoBaiGiftNotice
)

var (
	friendLangMap = map[LangCode]string{
		FriendIsNotFriend:           "此用户已经不是好友",
		FriendAlreadyFriend:         "此用户已经是好友",
		FriendAlreadyDelete:         "此用户已经不在联系列表",
		FriendAlreadyFull:           "好友已达上限",
		FriendPeerAlreadFull:        "对方好友已达上限",
		FriednBlackAlreadyFull:      "黑名单已达上限",
		FriendIsSpouseNoDelete:      "对方是您配偶,无法删除",
		FriendIsBlack:               "对方在您的黑名单里,请先将对方从黑名单中移除",
		FriendAddPeerBlack:          "对方把你拉黑了,无法请求",
		FriendNotInBlack:            "对方不在黑名单内",
		FriendBatchAgreeFull:        "您当前好友数量已达上限，无法继续添加",
		FriendGiveFlowersNotice:     "鲜花无声爱有声,%s向%s赠送了%s,双方感情再度加深，亲密度上升了%s",
		FriendAddRewHadDone:         "添加好友奖励已领取",
		FriendNumNotEnough:          "好友数量不足",
		FriendAskForGiveTitle:       "来自好友的馈赠",
		FriendAskForGiveContent:     "您的好友%s慷慨大方的给予了你想要的装备，不要忘了和他说声谢谢哦~",
		FriendGiveBiaoBaiGiftNotice: "%s向%s赠送了%s,双方感情再度加深",
	}
)

func init() {
	mergeLang(friendLangMap)
}
