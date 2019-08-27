package lang

const (
	RealmReachLimit LangCode = RealmBase + iota
	RealmLevelTooLow
	RealmEnterInOtherFuBen
	RealmPairNoSpouse
	RealmPairSpouseNoOnline
	RealmSpouseInOtherFuBen
	RealmPassRewTitle
	RealmPassRewContent
	RealmNoHoldWed
	RealmIn3v3Match
	RealmPeerIn3v3Match
	RealmReissueTitle
	RealmReissueContent
)

var (
	realmLangMap = map[LangCode]string{
		RealmReachLimit:         "境界已达最高层",
		RealmLevelTooLow:        "境界等级太低",
		RealmEnterInOtherFuBen:  "当前在其他副本场景,不允许进入天劫塔",
		RealmPairNoSpouse:       "当前没有配偶,结婚后可邀请配偶帮忙闯关天劫塔",
		RealmPairSpouseNoOnline: "您的配偶当前不在线,无法邀请助战",
		RealmSpouseInOtherFuBen: "您的配偶当前在其他副本场景,无法邀请",
		RealmPassRewTitle:       "天劫塔%d通关奖励",
		RealmPassRewContent:     "天劫塔通关奖励内容",
		RealmNoHoldWed:          "举办婚礼后可共同闯关",
		RealmPeerIn3v3Match:     "对方当前正在3v3匹配,邀请取消",
		RealmReissueTitle:       "天劫塔通关",
		RealmReissueContent:     "由于修改帝魂相关产出，天劫塔无法重新获得，现将之前可获得的帝魂技能书补发给您，请查收！",
	}
)

func init() {
	mergeLang(realmLangMap)
}
