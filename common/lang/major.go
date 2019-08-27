package lang

const (
	MajorInviteInOtherFuBen LangCode = MajorBase + iota
	MajorInviteNoTimes
	MajorInviteNoSpouse
	MajorInviteSpouseNoOnline
	MajorInviteSpouseInOtherFuBen
	MajorRewTitle
	MajorRewContent
	MajorNoHoldWed
	MajorPeerIn3v3Match
	MajorHadCancle
)

var (
	majorLangMap = map[LangCode]string{
		MajorInviteInOtherFuBen:       "当前在其他副本场景,不允许邀请进入双修",
		MajorInviteNoTimes:            "今日双修次数已用完",
		MajorInviteNoSpouse:           "当前没有配偶,结婚后可邀请配偶双修",
		MajorInviteSpouseNoOnline:     "您的配偶当前不在线,无法邀请双修",
		MajorInviteSpouseInOtherFuBen: "您的配偶当前在其他副本场景,无法邀请",
		MajorRewTitle:                 "夫妻双修奖励",
		MajorRewContent:               "夫妻双修奖励内容",
		MajorNoHoldWed:                "举办婚礼成为正式夫妻才可进入",
		MajorPeerIn3v3Match:           "对方当前正在3v3匹配,邀请取消",
		MajorHadCancle:                "您的配偶已经取消了本次副本邀请",
	}
)

func init() {
	mergeLang(majorLangMap)
}
