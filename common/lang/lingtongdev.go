package lang

const (
	LingTongDevEatCulDanReachedLimit LangCode = LingTongDevBase + iota
	LingTongDevEatCulDanReachedFull
	LingTongDevEatUnDanReachedLimit
	LingTongDevEatUnDanReachedFull
	LingTongDevUnrealCondNotReached
	LingTongDevAdanvacedReachedLimit
	LingTongDevUnrealNoExist
	LingTongDevOtherIsNoActived
	LingTongDevAdvanceToLow
	LingTongDevAdvanceToHigh
	LingTongDevAdvanceNotEqual
	LingTongDevSkinUpstarNoActive
	LingTongDevSkinReacheFullStar
	LingTongDevTongLingReachedFull
	LingTongDevActiveSystem
	LingTongDevNoActivateLingTong
	LingTongDevAdvancedNotice
	LingTongDevUnrealActivateNotice
)

var (
	lingTongDevLangMap = map[LangCode]string{
		LingTongDevEatCulDanReachedLimit: "%s食用培养丹等级已达最大,请进阶后再试",
		LingTongDevEatCulDanReachedFull:  "%s培养丹食丹等级满级",
		LingTongDevEatUnDanReachedLimit:  "%s幻化丹食丹等级已达最大,请进阶后再试",
		LingTongDevEatUnDanReachedFull:   "%s幻化丹食丹等级满级",
		LingTongDevUnrealCondNotReached:  "还有条件未达成，无法解锁幻化",
		LingTongDevAdanvacedReachedLimit: "%s已达最高阶",
		LingTongDevUnrealNoExist:         "当前没有幻化",
		LingTongDevOtherIsNoActived:      "该%s还未激活",
		LingTongDevAdvanceToLow:          "您%s系统的阶别不足,无法使用物品",
		LingTongDevAdvanceToHigh:         "您%s系统的阶别过高,无法使用物品",
		LingTongDevAdvanceNotEqual:       "您%s系统的阶别不符,无法使用物品",
		LingTongDevSkinUpstarNoActive:    "未激活的%s皮肤,无法升星",
		LingTongDevSkinReacheFullStar:    "%s皮肤已满星",
		LingTongDevTongLingReachedFull:   "%s通灵等级满级",
		LingTongDevActiveSystem:          "请先激活灵童%s系统",
		LingTongDevNoActivateLingTong:    "您还未激活过灵童,请先激活一只灵童",
		LingTongDevAdvancedNotice:        "恭喜%s成功将%s提升至%s，战力飙升%s",
		LingTongDevUnrealActivateNotice:  "恭喜%s成功将%s幻化为%s，战力飙升%s",
	}
)

func init() {
	mergeLang(lingTongDevLangMap)
}
