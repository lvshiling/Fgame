package lang

const (
	BaGuaReachLimit LangCode = BaGuaBase + iota
	BaGuaLevelTooLow
	BaGuaEnterInOtherFuBen
	BaGuaPairNoSpouse
	BaGuaPairSpouseNoOnline
	BaGuaSpouseInOtherFuBen
	BaGuaPassRewTitle
	BaGuaPassRewContent
	BaGuaNoHoldWed
	BaGuaIn3v3Match
	BaGuaPeerIn3v3Match
	BaGuaInviteCd
	BaGuaEmailTitleBuChang
	BaGuaEmailContentBuChang
)

var (
	baGuaLangMap = map[LangCode]string{
		BaGuaReachLimit:          "八卦秘境已达最高层",
		BaGuaLevelTooLow:         "八卦秘境等级太低",
		BaGuaEnterInOtherFuBen:   "当前在其他副本场景,不允许进入八卦秘境",
		BaGuaPairNoSpouse:        "当前没有配偶,结婚后可邀请配偶帮忙闯关八卦秘境",
		BaGuaPairSpouseNoOnline:  "您的配偶当前不在线,无法邀请助战",
		BaGuaSpouseInOtherFuBen:  "您的配偶当前在其他副本场景,无法邀请",
		BaGuaPassRewTitle:        "八卦秘境%d通关奖励",
		BaGuaPassRewContent:      "八卦秘境通关奖励内容",
		BaGuaNoHoldWed:           "举办婚礼后可共同闯关",
		BaGuaPeerIn3v3Match:      "对方当前正在3v3匹配,邀请取消",
		BaGuaInviteCd:            "操作太频繁,请%s秒后再试",
		BaGuaEmailTitleBuChang:   "八卦秘境通关",
		BaGuaEmailContentBuChang: "由于八卦秘境新增八卦符石产出，八卦秘境无法重新通关获得，现将之前已通关可获得的八卦符石补发给您，请查收！",
	}
)

func init() {
	mergeLang(baGuaLangMap)
}
