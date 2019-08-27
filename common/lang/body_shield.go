package lang

const (
	BodyShieldEatJJDanReachedLimit = BodyShieldBase + iota
	BodyShieldEatJJDanReachedFull
	BodyShieldAdanvacedReachedLimit
	BodyShieldShieldReachedLimits
	BodyShieldAdvanceToLow
	BodyShieldAdvanceToHigh
	BodyShieldAdvanceNotEqual
	BodyShieldAdvanceNotice
	ShieldAdvanceNotEqual
	ShieldAdvanceNotice
)

var (
	bodyShieldLangMap = map[LangCode]string{
		BodyShieldEatJJDanReachedLimit:  "护体金甲丹食丹等级已达最大,请进阶后再试",
		BodyShieldEatJJDanReachedFull:   "护体金甲丹食丹等级满级",
		BodyShieldAdanvacedReachedLimit: "护体盾已达最高阶",
		BodyShieldShieldReachedLimits:   "神盾尖刺已达最高阶",
		BodyShieldAdvanceToLow:          "您护体盾系统的阶别不足，无法使用物品",
		BodyShieldAdvanceToHigh:         "您护体盾系统的阶别过高，无法使用物品",
		BodyShieldAdvanceNotEqual:       "您护体盾系统的阶别不符，无法使用物品",
		BodyShieldAdvanceNotice:         "铜墙铁壁，%s成功将护盾提升至%s，战力飙升%s，受到伤害大幅度降低",
		ShieldAdvanceNotEqual:           "您盾刺系统的阶别不符，无法使用物品",
		ShieldAdvanceNotice:             "破格无双，%s成功将盾刺提升至%s】，战力飙升%s",
	}
)

func init() {
	mergeLang(bodyShieldLangMap)
}
