package lang

const (
	WingUnrealDanReachedLimit LangCode = WingBase + iota
	WingUnrealDanReachedFull
	WingUnrealCondNotReached
	WingAdanvacedReachedLimit
	WingUnrealNoExist
	WingOtherIsNoActived
	WingTrialCardUseAdvancedIsZero
	WingTrialCardUseIsExist
	WingAdvanceToLow
	WingAdvanceToHigh
	WingAdvanceNotEqual
	FeatherAdvanceNotEqual
	WingSkinUpstarNoActive
	WingSkinReacheFullStar
	WingAdvancedNotice
	WingUnrealActivateNotice
	FeatherAdvancedNotice
	WingUnrealActiveSystem
	WingHadActivate
)

var (
	wingLangMap = map[LangCode]string{
		WingUnrealDanReachedLimit:      "幻化丹食丹等级已达最大,请进阶后再试",
		WingUnrealDanReachedFull:       "幻化丹食丹等级满级",
		WingUnrealCondNotReached:       "还有幻化条件未达成，无法解锁幻化",
		WingAdanvacedReachedLimit:      "已达最高阶",
		WingUnrealNoExist:              "当前没有幻化",
		WingOtherIsNoActived:           "该战翼还未激活",
		WingTrialCardUseAdvancedIsZero: "您的战翼阶数,无法使用战翼试用卡",
		WingTrialCardUseIsExist:        "已获得战翼试用,期间无法重复获取",
		WingAdvanceToLow:               "您战翼系统的阶别不足，无法使用物品",
		WingAdvanceToHigh:              "您战翼系统的阶别过高，无法使用物品",
		WingAdvanceNotEqual:            "您战翼系统的阶别不符，无法使用物品",
		FeatherAdvanceNotEqual:         "免爆仙羽的阶别不符，无法使用物品",
		WingSkinUpstarNoActive:         "未激活的战翼皮肤,无法升星",
		WingSkinReacheFullStar:         "战翼皮肤已满星",
		WingAdvancedNotice:             "恭喜%s成功将战翼提升至%s，战力飙升%s",
		WingUnrealActivateNotice:       "恭喜%s成功将战翼幻化为%s，战力飙升%s",
		FeatherAdvancedNotice:          "%s成功将仙羽提升至%s，战力飙升%s",
		WingUnrealActiveSystem:         "请先激活战翼系统",
		WingHadActivate:                "您当前已激活战翼系统",
	}
)

func init() {
	mergeLang(wingLangMap)
}
